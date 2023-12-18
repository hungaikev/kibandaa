package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type PaymentsServer struct {
	Log   *zerolog.Logger
	build string
}

// GetLiveness returns a simple UP message to indicate the service is running.
func (p *PaymentsServer) GetLiveness(c *gin.Context, params GetLivenessParams) {
	p.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	c.JSON(200, gin.H{
		"message": "UP",
	})
}

// PostPaymentCallback handles the callback from the payment provider.
func (p *PaymentsServer) PostPaymentCallback(c *gin.Context) {
	ctx := c.Request.Context()
	p.Log.Info().Msgf("Request ID: %s", ctx.Value("request_id"))

	var callbackMessage MpesaCallbackMessage
	if err := c.ShouldBindJSON(&callbackMessage); err != nil {
		p.Log.Error().Err(err).Msg("failed to parse callback message")
		c.JSON(400, gin.H{
			"message": "failed to parse callback message",
		})
		return
	}

	p.Log.Info().Msgf("Received callback message: %+v", callbackMessage)

	c.JSON(200, callbackMessage)

}

// NewPaymentsServer constructs a new PaymentsServer.
func NewPaymentsServer(log *zerolog.Logger, build string) *PaymentsServer {
	return &PaymentsServer{
		Log:   log,
		build: build,
	}
}
