package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type OrdersServer struct {
	Log     *zerolog.Logger
	build   string
	storage *Repository
}

// GetCustomers returns a list of customers
func (o *OrdersServer) GetCustomers(c *gin.Context, params GetCustomersParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	customers, err := o.storage.GetCustomers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customers)
}

// CreateCustomer creates a new customer
func (o *OrdersServer) CreateCustomer(c *gin.Context, params CreateCustomerParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	var customerRequest NewCustomerRequest
	if err := c.BindJSON(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cus := Customer{
		Id:    uuid.New(),
		Name:  customerRequest.Name,
		Email: customerRequest.Email,
		Phone: customerRequest.Phone,
	}

	customer, err := o.storage.PostCustomer(ctx, cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customer)

}

// DeleteCustomerByID deletes a customer by ID
func (o *OrdersServer) DeleteCustomerByID(c *gin.Context, customerID CustomerID, params DeleteCustomerByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	custID, err := uuid.Parse(customerID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = o.storage.DeleteCustomerByID(ctx, custID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted successfully",
	})
}

// GetCustomerByID returns a customer by ID
func (o *OrdersServer) GetCustomerByID(c *gin.Context, customerID CustomerID, params GetCustomerByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	custID, err := uuid.Parse(customerID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer, err := o.storage.GetCustomerByID(ctx, custID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UpdateCustomerByID updates a customer by ID
func (o *OrdersServer) UpdateCustomerByID(c *gin.Context, customerID CustomerID, params UpdateCustomerByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	custID, err := uuid.Parse(customerID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var customerUpdateRequest Customer
	if err := c.BindJSON(&customerUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cus := Customer{
		Id:    custID,
		Name:  customerUpdateRequest.Name,
		Email: customerUpdateRequest.Email,
		Phone: customerUpdateRequest.Phone,
	}

	customer, err := o.storage.UpdateCustomerByID(ctx, cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// GetLiveness returns a 200 OK response
func (o *OrdersServer) GetLiveness(c *gin.Context, params GetLivenessParams) {
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	c.JSON(http.StatusOK, gin.H{
		"message": "UP",
	})

}

// GetOrders returns a list of orders
func (o *OrdersServer) GetOrders(c *gin.Context, params GetOrdersParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	orders, err := o.storage.GetOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)

}

// CreateOrder creates a new order
func (o *OrdersServer) CreateOrder(c *gin.Context, params CreateOrderParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	var orderRequest NewOrderRequest
	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	products := make([]Product, len(orderRequest.Products))
	for i, p := range orderRequest.Products {
		products[i] = Product{
			Name:        p.Name,
			UnitPrice:   p.UnitPrice,
			Description: p.Description,
		}
	}

	order := Order{
		Customer: Customer{
			Email: orderRequest.Customer.Email,
			Name:  orderRequest.Customer.Name,
			Phone: orderRequest.Customer.Phone,
		},
		Id:       uuid.New(),
		Products: products,
		Status:   "PENDING",
		Taxes:    orderRequest.Taxes,
		Total:    orderRequest.Total,
	}

	ord, err := o.storage.PostOrder(ctx, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ord)
}

// DeleteOrderByID deletes an order by ID
func (o *OrdersServer) DeleteOrderByID(c *gin.Context, orderID OrderID, params DeleteOrderByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	ordID, err := uuid.Parse(orderID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = o.storage.DeleteOrderByID(ctx, ordID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}

// GetOrderByID returns an order by ID
func (o *OrdersServer) GetOrderByID(c *gin.Context, orderID OrderID, params GetOrderByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	ordID, err := uuid.Parse(orderID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := o.storage.GetOrderByID(ctx, ordID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrderByID updates an order by ID
func (o *OrdersServer) UpdateOrderByID(c *gin.Context, orderID OrderID, params UpdateOrderByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	ordID, err := uuid.Parse(orderID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var orderUpdateRequest Order
	if err := c.BindJSON(&orderUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	orderUpdateRequest.Id = ordID

	order, err := o.storage.UpdateOrderByID(ctx, orderUpdateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)

}

// GetProducts returns a list of products
func (o *OrdersServer) GetProducts(c *gin.Context, params GetProductsParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	products, err := o.storage.GetProducts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

// CreateProduct creates a new product
func (o *OrdersServer) CreateProduct(c *gin.Context, params CreateProductParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	var productRequest NewProduct
	if err := c.BindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product := Product{
		Name:        productRequest.Name,
		UnitPrice:   productRequest.UnitPrice,
		Description: productRequest.Description,
	}

	prod, err := o.storage.PostProduct(ctx, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, prod)
}

// DeleteProductByID deletes a product by ID
func (o *OrdersServer) DeleteProductByID(c *gin.Context, productID ProductID, params DeleteProductByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	prodID, err := uuid.Parse(productID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = o.storage.DeleteProductByID(ctx, prodID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

// GetProductByID returns a product by ID
func (o *OrdersServer) GetProductByID(c *gin.Context, productID ProductID, params GetProductByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	prodID, err := uuid.Parse(productID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := o.storage.GetProductByID(ctx, prodID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProductByID updates a product by ID
func (o *OrdersServer) UpdateProductByID(c *gin.Context, productID ProductID, params UpdateProductByIDParams) {
	ctx := c.Request.Context()
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	prodID, err := uuid.Parse(productID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var productUpdateRequest Product
	if err := c.BindJSON(&productUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productUpdateRequest.Id = prodID

	product, err := o.storage.UpdateProductByID(ctx, productUpdateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetReadiness returns a 200 OK response
func (o *OrdersServer) GetReadiness(c *gin.Context, params GetReadinessParams) {
	o.Log.Info().Msgf("Request ID: %s", params.XRequestID.String())

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// NewOrdersServer constructs a new OrdersServer.
func NewOrdersServer(log *zerolog.Logger, build string, storage *Repository) *OrdersServer {
	return &OrdersServer{
		Log:     log,
		build:   build,
		storage: storage,
	}
}
