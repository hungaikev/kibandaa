package storage

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/rs/zerolog"
	"google.golang.org/api/iterator"

	api "github.com/hungaikev/kibandaa/orders/internal/api/rest/v1"
)

// Store provides all functions to execute database queries and transactions.
type Store interface {
	GetCustomers(ctx context.Context) ([]api.Customer, error)
	PostCustomer(ctx context.Context, customer api.Customer) (api.Customer, error)
	GetCustomerByID(ctx context.Context, id string) (api.Customer, error)
	DeleteCustomerByID(ctx context.Context, id string) error
	UpdateCustomerByID(ctx context.Context, customer api.Customer) (api.Customer, error)
	GetOrders(ctx context.Context) ([]api.Order, error)
	PostOrder(ctx context.Context, order api.Order) (api.Order, error)
	GetOrderByID(ctx context.Context, id string) (api.Order, error)
	DeleteOrderByID(ctx context.Context, id string) error
	UpdateOrderByID(ctx context.Context, order api.Order) (api.Order, error)
	GetProducts(ctx context.Context) ([]api.Product, error)
	PostProduct(ctx context.Context, product api.Product) (api.Product, error)
	GetProductByID(ctx context.Context, id string) (api.Product, error)
	DeleteProductByID(ctx context.Context, id string) error
	UpdateProductByID(ctx context.Context, product api.Product) (api.Product, error)
}

// Repository provides all functions to execute database queries and transactions.
type Repository struct {
	client *firestore.Client
	log    *zerolog.Logger
}

// NewRepository creates a new store.
func NewRepository(log *zerolog.Logger, firestoreClient *firestore.Client) (*Repository, error) {
	return &Repository{
		client: firestoreClient,
		log:    log,
	}, nil
}

// GetCustomers returns all customers
func (r *Repository) GetCustomers(ctx context.Context) ([]api.Customer, error) {
	var customers []api.Customer

	iter := r.client.Collection("customers").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}

		var customer api.Customer
		if err := doc.DataTo(&customer); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil

}

// PostCustomer creates a customer
func (r *Repository) PostCustomer(ctx context.Context, customer api.Customer) (api.Customer, error) {
	_, err := r.client.Collection("customers").Doc(customer.Id.String()).Set(ctx, customer)
	if err != nil {
		return api.Customer{}, err
	}

	customer, err = r.GetCustomerByID(ctx, customer.Id.String())
	if err != nil {
		return api.Customer{}, err
	}

	return customer, err
}

// GetCustomerByID returns a single customer
func (r *Repository) GetCustomerByID(ctx context.Context, id string) (api.Customer, error) {
	var customer api.Customer
	doc, err := r.client.Collection("customers").Doc(id).Get(ctx)
	if err != nil {
		return customer, err
	}

	err = doc.DataTo(&customer)
	return customer, err
}

// DeleteCustomerByID deletes a single customer
func (r *Repository) DeleteCustomerByID(ctx context.Context, id string) error {
	_, err := r.client.Collection("customers").Doc(id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCustomerByID updates a single customer
func (r *Repository) UpdateCustomerByID(ctx context.Context, customer api.Customer) (api.Customer, error) {
	_, err := r.client.Collection("customers").Doc(customer.Id.String()).Set(ctx, customer)
	if err != nil {
		return api.Customer{}, err
	}

	customer, err = r.GetCustomerByID(ctx, customer.Id.String())
	if err != nil {
		return api.Customer{}, err
	}

	return customer, err
}

// GetOrders returns all orders
func (r *Repository) GetOrders(ctx context.Context) ([]api.Order, error) {
	var orders []api.Order

	iter := r.client.Collection("orders").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}

		var order api.Order
		if err := doc.DataTo(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// PostOrder creates an order
func (r *Repository) PostOrder(ctx context.Context, order api.Order) (api.Order, error) {
	_, err := r.client.Collection("orders").Doc(order.Id.String()).Set(ctx, order)
	if err != nil {
		return api.Order{}, err
	}

	order, err = r.GetOrderByID(ctx, order.Id.String())
	if err != nil {
		return api.Order{}, err
	}

	return order, err
}

// GetOrderByID returns a single order
func (r *Repository) GetOrderByID(ctx context.Context, id string) (api.Order, error) {
	var order api.Order
	doc, err := r.client.Collection("orders").Doc(id).Get(ctx)
	if err != nil {
		return order, err
	}

	err = doc.DataTo(&order)

	return order, err
}

// DeleteOrderByID deletes a single order
func (r *Repository) DeleteOrderByID(ctx context.Context, id string) error {
	_, err := r.client.Collection("orders").Doc(id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOrderByID updates a single order
func (r *Repository) UpdateOrderByID(ctx context.Context, order api.Order) (api.Order, error) {
	_, err := r.client.Collection("orders").Doc(order.Id.String()).Set(ctx, order)
	if err != nil {
		return api.Order{}, err
	}

	order, err = r.GetOrderByID(ctx, order.Id.String())
	if err != nil {
		return api.Order{}, err
	}

	return order, err
}

// GetProducts returns all products
func (r *Repository) GetProducts(ctx context.Context) ([]api.Product, error) {
	var products []api.Product

	iter := r.client.Collection("products").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}

		var product api.Product
		if err := doc.DataTo(&product); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// PostProduct creates a product
func (r *Repository) PostProduct(ctx context.Context, product api.Product) (api.Product, error) {
	_, err := r.client.Collection("products").Doc(product.Id.String()).Set(ctx, product)
	if err != nil {
		return api.Product{}, err
	}

	product, err = r.GetProductByID(ctx, product.Id.String())
	if err != nil {
		return api.Product{}, err
	}

	return product, err
}

// GetProductByID returns a single product
func (r *Repository) GetProductByID(ctx context.Context, id string) (api.Product, error) {
	var product api.Product
	doc, err := r.client.Collection("products").Doc(id).Get(ctx)
	if err != nil {
		return product, err
	}

	err = doc.DataTo(&product)

	return product, err
}

// DeleteProductByID deletes a single product
func (r *Repository) DeleteProductByID(ctx context.Context, id string) error {
	_, err := r.client.Collection("products").Doc(id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProductByID updates a single product
func (r *Repository) UpdateProductByID(ctx context.Context, product api.Product) (api.Product, error) {
	_, err := r.client.Collection("products").Doc(product.Id.String()).Set(ctx, product)
	if err != nil {
		return api.Product{}, err
	}

	product, err = r.GetProductByID(ctx, product.Id.String())
	if err != nil {
		return api.Product{}, err
	}

	return product, err
}
