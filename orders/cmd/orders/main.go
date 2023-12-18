package main

import (
	"context"
	"expvar"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/ardanlabs/conf/v3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/hungaikev/kibandaa/orders/internal/api/handlers"
	api "github.com/hungaikev/kibandaa/orders/internal/api/rest/v1"
)

const (
	// LogStrKeyModule is for use with the logger as a key to specify the module name.
	LogStrKeyModule = "module"
	// LogStrKeyService is for use with the logger as a key to specify the service name.
	LogStrKeyService = "service"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {

	z := zerolog.New(os.Stderr).With().Str(LogStrKeyService, "order").Timestamp().Logger().With().Caller().Logger()
	mainLog := z.With().Str(LogStrKeyModule, "main").Logger()
	mainLog.Info().Msg("starting server...")

	if err := run(&mainLog); err != nil {
		mainLog.Info().Msgf("main: error %s:", err.Error())
		os.Exit(1)
	}
}

func run(log *zerolog.Logger) error {
	log.Info().Msg("Welcome to the Order Service :)")

	// =========================================================================
	// Configuration

	// Call env.Load func to ensure a .env file is loaded when available.
	// This command only affects dev environments.

	LoadDevEnv(log)

	var cfg struct {
		conf.Version
		Web struct {
			APIPort         string        `conf:"default:8000"`
			ShutdownTimeout time.Duration `conf:"default:10s"`
		}
	}

	cfg.Version.Build = build
	cfg.Version.Desc = "Order Service"

	// =========================================================================
	// App Starting

	expvar.NewString("build").Set(build)
	log.Info().Msgf("Started: Application initializing: version %q", build)
	defer log.Info().Msg("Completed")

	help, err := conf.Parse("", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			log.Info().Msg(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	log.Info().Msgf("Config:\n%v\n", out)

	// =========================================================================
	// Start API Service

	log.Info().Msg("main: Initializing API support")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Create a connection to the database.
	firestoreClient, err := firestore.NewClient(context.Background(), "kibandaa")
	if err != nil {
		return errors.Wrap(err, "could not create firestore client")
	}
	// Create a new storage type and connection.
	store, err := api.NewRepository(log, firestoreClient)
	if err != nil {
		return errors.Wrap(err, "could not create storage")
	}

	ordersServer := api.NewOrdersServer(log, build, store)

	server := handlers.API(ordersServer, cfg.Web.APIPort)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Info().Msgf("main: API listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Info().Msgf("main: %v: Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := server.Shutdown(ctx); err != nil {
			if err := server.Close(); err != nil {
				return errors.Wrap(err, "could not stop server gracefully")
			}
			return errors.Wrap(err, "could not stop server")
		}
	}

	return nil
}
