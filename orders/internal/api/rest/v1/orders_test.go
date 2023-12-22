package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
}

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
