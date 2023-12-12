package handlers

import (
	"net"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"

	api "github.com/hungaikev/kibandaa/orders/internal/api/rest/v1"
)

func API(ordersServer *api.OrdersServer, port string) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		ordersServer.Log.Fatal().Err(err).Msg("failed to load swagger spec")
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(middleware.OapiRequestValidator(swagger))
	r.Use(cors.Default())
	r.Use(requestid.New())

	api.RegisterHandlers(r, ordersServer)
	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}

	return s
}
