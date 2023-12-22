package v1

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetLiveness(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	paymentServer := NewPaymentsServer(nil, "1.0.0")
	router.GET("/liveness", func(c *gin.Context) {
		paymentServer.GetLiveness(c, GetLivenessParams{})
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
}

func TestPostPaymentCallback(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	logger := zerolog.Nop()
	paymentServer := NewPaymentsServer(&logger, "1.0.0")

	router.POST("/payments/callback", func(c *gin.Context) {
		paymentServer.PostPaymentCallback(c)
	})

	requestBody := `{
  "transactionType": "PayBill",
  "transID": "LGR12345678",
  "transAmount": "1000.00",
  "tusinessShortCode": "174379",
  "tillRefNumber": "order123",
  "tnvoiceNumber": "",
  "orgAccountBalance": "50000.00",
  "thirdPartyTransID": "0",
  "mSISDN": "254708374149",
  "firstName": "John",
  "middleName": "Doe",
  "lastName": "Smith",
  "transactionStatus": "Completed",
  "resultCode": "0",
  "resultDesc": "The service request is processed successfully."
	}`

	req, err := http.NewRequest("POST", "/payments/callback", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		t.Fatal(err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body content (can be further validated based on your implementation)
	// For now, let's check if the response body is not empty
	assert.NotEmpty(t, w.Body.String())
}
