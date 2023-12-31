// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/oapi-codegen/runtime"
)

// Defines values for LivenessStatus.
const (
	DOWN LivenessStatus = "DOWN"
	UP   LivenessStatus = "UP"
)

// N400 Bad Request
type N400 = interface{}

// N401 Unauthorized
type N401 = interface{}

// N403 Forbidden
type N403 = interface{}

// N404 Not Found
type N404 = interface{}

// N500 Internal Server Error
type N500 = interface{}

// Error General API Error Response
type Error struct {
	Code      int        `json:"code"`
	Domain    *string    `json:"domain,omitempty"`
	Message   string     `json:"message"`
	Reason    *string    `json:"reason,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// Liveness defines model for Liveness.
type Liveness struct {
	Message *string         `json:"message,omitempty" validate:"required"`
	Status  *LivenessStatus `json:"status,omitempty" validate:"required,oneof=UP DOWN"`
}

// LivenessStatus defines model for Liveness.Status.
type LivenessStatus string

// MpesaCallbackMessage defines model for MpesaCallbackMessage.
type MpesaCallbackMessage struct {
	BillRefNumber     *string `json:"billRefNumber,omitempty"`
	BusinessShortCode *string `json:"businessShortCode,omitempty"`
	FirstName         *string `json:"firstName,omitempty"`
	InvoiceNumber     *string `json:"invoiceNumber,omitempty"`
	LastName          *string `json:"lastName,omitempty"`
	MiddleName        *string `json:"middleName,omitempty"`
	Msisdn            *string `json:"msisdn,omitempty"`
	OrgAccountBalance *string `json:"orgAccountBalance,omitempty"`
	ResultCode        *string `json:"resultCode,omitempty"`
	ResultDesc        *string `json:"resultDesc,omitempty"`
	ThirdPartyTransID *string `json:"thirdPartyTransID,omitempty"`
	TransAmount       *string `json:"transAmount,omitempty"`
	TransID           *string `json:"transID,omitempty"`
	TransactionStatus *string `json:"transactionStatus,omitempty"`
	TransactionType   *string `json:"transactionType,omitempty"`
}

// XRequestIdType X-RequestID
type XRequestIdType = uuid.UUID

// XRequestID X-RequestID
type XRequestID = XRequestIdType

// GetLivenessParams defines parameters for GetLiveness.
type GetLivenessParams struct {
	// XRequestID X-Request-ID
	XRequestID *XRequestID `json:"X-Request-ID,omitempty"`
}

// PostPaymentCallbackJSONRequestBody defines body for PostPaymentCallback for application/json ContentType.
type PostPaymentCallbackJSONRequestBody = MpesaCallbackMessage

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get liveness status
	// (GET /liveness)
	GetLiveness(c *gin.Context, params GetLivenessParams)
	// Callback endpoint for M-Pesa payment updates
	// (POST /payments/callback)
	PostPaymentCallback(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetLiveness operation middleware
func (siw *ServerInterfaceWrapper) GetLiveness(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLivenessParams

	headers := c.Request.Header

	// ------------- Optional header parameter "X-Request-ID" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Request-ID")]; found {
		var XRequestID XRequestID
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for X-Request-ID, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Request-ID", runtime.ParamLocationHeader, valueList[0], &XRequestID)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter X-Request-ID: %w", err), http.StatusBadRequest)
			return
		}

		params.XRequestID = &XRequestID

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetLiveness(c, params)
}

// PostPaymentCallback operation middleware
func (siw *ServerInterfaceWrapper) PostPaymentCallback(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostPaymentCallback(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/liveness", wrapper.GetLiveness)
	router.POST(options.BaseURL+"/payments/callback", wrapper.PostPaymentCallback)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xX224bNxD9FYLt465WbhIgENAHX5rWQO0IdoQWMPwwIke7TLkkw4tg1dC/FyR3tZa1",
	"cgPU7pMlzeXMkGfO0I+U6dZohco7OnukBiy06NGmb3+WN/gtoPOXF/ErR8esMF5oRWeDsby8oAUV8bcG",
	"gaOlBVXQ4qGPYw22EFP9aHFFZ/SHaoCvstVVT1G3220flip6P53GP0wrj8rHj2CMFAxiUdVXFyt7/E6c",
	"X6zVNiPsd3YGnHQl0G1B309P3h5zoSD4RlvxN/IM+u7tQT9puxSco8qI798e8Vp78kkHlXr88H9c5qXy",
	"aBVIcot2jZZ0jgXNHw5o/SsqtCDJ6fwy+5IbdEYrh7SgxmqD1gt0uXCO8a/fmEh2oTzWmHJz3YJQT2zO",
	"W6HqaGrROahx1GYRuqYPTF606Dy0JlpX2rbg6Yxy8FhGEy2eh6R034KwyOnsbgdb5KrvY0YvY0A+h128",
	"Xn5Flnj/u1ijQpd1Ya/xoz0U9KHUYEQZQWpUJT54C6WHOsWtQYpYMp0NtaUB9+BD8kAV2ljuYk4LevH5",
	"j+tU6X/EKLRCvfp5MScpYyLJQbdXBh2cg5RLYH9dDQ3ud74UUt7g6jq0S7SjF7UMTsRTu2209ef7DBm8",
	"VsI6f51EcsQq1FoLhi+gSHghvBWcSzxudsLxcZZpW58ypoPyZyBBsWM0dUEe7y2bL9CxcSY3wvI5WL/5",
	"YkG5vFoOvaLttI2lHLe/FAssDvTtjlkveX1JtseRETqgyfetxLTtHspal/0i7A08YXXGLnsIgk8WiyEo",
	"/l6K1mibmu9yRLeoQeAbOqO18E1YTphuq1rrWmKV7NvEbqFW+rC830BxiY4Y2LRRPIlvrA51Q67KOTog",
	"oDgxVq8FR0dYNwoEFTdaxPfBoBnznCFpqmAYxZIWdI3WZaSTyXQyTYQyqMAIOqPvJieTaVd+upBKPtGX",
	"GlOncdSS+l/yJMV+p0HF3tPkbnwZDC77j4j7RMok4gntp1dcPLsKR3bPbWAMnVsFSXad5VX7evgx17+/",
	"YV5z1X4YRzyyaKO8h7YFu8k3SvprJ53qFzRL992wcO5jVNXTtOqpmORYuxGmzLXzHSd7Cad5/6HzZ5pv",
	"Xq350TWx3d+23gbcjlNu/8j6PMQiQ7FGTtyOMXLz7OzOnw8kWWnbj253ViSYuPueHmp3LOlQY750OXmE",
	"9ou5wDVKbVKa7EULGqyMz3rvzayqpGYgG+387OP040kaqw7lea7P/e04YlGCR0683unO8P/Brrjt/faf",
	"AAAA///L4pKYiwwAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
