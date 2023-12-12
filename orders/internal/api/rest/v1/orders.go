package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type OrdersServer struct {
	Log   *zerolog.Logger
	build string
}

func (o OrdersServer) GetCustomers(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) PostCustomers(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) DeleteCustomersCustomerId(c *gin.Context, customerId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) GetCustomersCustomerId(c *gin.Context, customerId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) PutCustomersCustomerId(c *gin.Context, customerId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) GetLiveness(c *gin.Context, params GetLivenessParams) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) GetOrders(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) PostOrders(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) DeleteOrdersOrderId(c *gin.Context, orderId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) GetOrdersOrderId(c *gin.Context, orderId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) PutOrdersOrderId(c *gin.Context, orderId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) GetProducts(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) PostProducts(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) DeleteProductsProductId(c *gin.Context, productId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) GetProductsProductId(c *gin.Context, productId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) PutProductsProductId(c *gin.Context, productId int64) {
	//TODO implement me
	panic("implement me")
}

func (o OrdersServer) GetReadiness(c *gin.Context, params GetReadinessParams) {
	//TODO implement me
	panic("implement me")
}

// NewOrdersServer constructs a new OrdersServer.
func NewOrdersServer(log *zerolog.Logger, build string) *OrdersServer {
	return &OrdersServer{
		Log:   log,
		build: build,
	}
}
