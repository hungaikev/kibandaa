package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/*func newFirestoreTestClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, "test")
	if err != nil {
		panic(err)
	}

	return client
}

func TestMain(m *testing.M) {
	// command to start firestore emulator
	cmd := exec.Command("gcloud", "beta", "emulators", "firestore", "start", "--host-port=localhost")

	// this makes it killable
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// we need to capture it's output to know when it's started
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stderr.Close()

	// start her up!
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// ensure the process is killed when we're finished, even if an error occurs
	// (thanks to Brian Moran for suggestion)
	var result int
	defer func() {
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		os.Exit(result)
	}()

	// we're going to wait until it's running to start
	var wg sync.WaitGroup
	wg.Add(1)

	// by starting a separate go routine
	go func() {
		// reading it's output
		buf := make([]byte, 256, 256)
		for {
			n, err := stderr.Read(buf[:])
			if err != nil {
				// until it ends
				if err == io.EOF {
					break
				}
				log.Fatalf("reading stderr %v", err)
			}

			if n > 0 {
				d := string(buf[:n])

				// only required if we want to see the emulator output
				log.Printf("%s", d)

				// checking for the message that it's started
				if strings.Contains(d, "Dev App Server is now running") {
					wg.Done()
				}

				// and capturing the FIRESTORE_EMULATOR_HOST value to set
				pos := strings.Index(d, FirestoreEmulatorHost+"=")
				if pos > 0 {
					host := d[pos+len(FirestoreEmulatorHost)+1 : len(d)-1]
					os.Setenv(FirestoreEmulatorHost, host)
				}
			}
		}
	}()

	// wait until the running message has been received
	wg.Wait()

	// now it's running, we can run our unit tests
	result = m.Run()
}

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"

func TestGetLiveness(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	orderServer := NewOrdersServer(nil, "1.0.0", nil)
	router.GET("/liveness", func(c *gin.Context) {
		orderServer.GetLiveness(c, GetLivenessParams{})
	})

	req, err := http.NewRequest("GET", "/liveness", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body content
	expectedBody := `{"message":"UP"}`
	assert.Equal(t, expectedBody, w.Body.String())
}*/

func TestGetReadiness(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	orderServer := NewOrdersServer(nil, "1.0.0", nil)
	router.GET("/readiness", func(c *gin.Context) {
		orderServer.GetReadiness(c, GetReadinessParams{})
	})

	req, err := http.NewRequest("GET", "/readiness", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body content
	expectedBody := `{"message":"OK"}`
	assert.Equal(t, expectedBody, w.Body.String())
}

/*func TestCreateCustomer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	logger := zerolog.Nop()
	repository, _ := NewRepository(&logger, newFirestoreTestClient(context.Background()))

	orderServer := NewOrdersServer(nil, "1.0.0", repository)

	router.POST("/customers", func(c *gin.Context) {
		orderServer.CreateCustomer(c, CreateCustomerParams{})
	})

	requestBody := `{
		"firstName": "John",
		"lastName": "Doe",
		"email": "hungaikevin@gmail.com"
  }`

	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		t.Fatal(err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.NotEmpty(t, w.Body.String())
}*/
