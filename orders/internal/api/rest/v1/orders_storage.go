package v1

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/rs/zerolog"
	"google.golang.org/api/iterator"
)

// Store provides all functions to execute database queries and transactions.
type Store interface {
	GetCustomers(ctx context.Context) ([]Customer, error)
	PostCustomer(ctx context.Context, customer Customer) (Customer, error)
	GetCustomerByID(ctx context.Context, id string) (Customer, error)
	DeleteCustomerByID(ctx context.Context, id string) error
	UpdateCustomerByID(ctx context.Context, customer Customer) (Customer, error)
	GetOrders(ctx context.Context) ([]Order, error)
	PostOrder(ctx context.Context, order Order) (Order, error)
	GetOrderByID(ctx context.Context, id string) (Order, error)
	DeleteOrderByID(ctx context.Context, id string) error
	UpdateOrderByID(ctx context.Context, order Order) (Order, error)
	GetProducts(ctx context.Context) ([]Product, error)
	PostProduct(ctx context.Context, product Product) (Product, error)
	GetProductByID(ctx context.Context, id string) (Product, error)
	DeleteProductByID(ctx context.Context, id string) error
	UpdateProductByID(ctx context.Context, product Product) (Product, error)
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
func (r *Repository) GetCustomers(ctx context.Context) ([]Customer, error) {
	var customers []Customer

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

		var customer Customer
		if err := doc.DataTo(&customer); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil

}

// PostCustomer creates a customer
func (r *Repository) PostCustomer(ctx context.Context, customer Customer) (Customer, error) {
	_, err := r.client.Collection("customers").Doc(customer.Id.String()).Set(ctx, customer)
	if err != nil {
		return Customer{}, err
	}

	customer, err = r.GetCustomerByID(ctx, customer.Id.String())
	if err != nil {
		return Customer{}, err
	}

	return customer, err
}

// GetCustomerByID returns a single customer
func (r *Repository) GetCustomerByID(ctx context.Context, id string) (Customer, error) {
	var customer Customer
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
func (r *Repository) UpdateCustomerByID(ctx context.Context, customer Customer) (Customer, error) {
	_, err := r.client.Collection("customers").Doc(customer.Id.String()).Set(ctx, customer)
	if err != nil {
		return Customer{}, err
	}

	customer, err = r.GetCustomerByID(ctx, customer.Id.String())
	if err != nil {
		return Customer{}, err
	}

	return customer, err
}

// GetOrders returns all orders
func (r *Repository) GetOrders(ctx context.Context) ([]Order, error) {
	var orders []Order

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

		var order Order
		if err := doc.DataTo(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// PostOrder creates an order
func (r *Repository) PostOrder(ctx context.Context, order Order) (Order, error) {
	_, err := r.client.Collection("orders").Doc(order.Id.String()).Set(ctx, order)
	if err != nil {
		return Order{}, err
	}

	order, err = r.GetOrderByID(ctx, order.Id.String())
	if err != nil {
		return Order{}, err
	}

	return order, err
}

// GetOrderByID returns a single order
func (r *Repository) GetOrderByID(ctx context.Context, id string) (Order, error) {
	var order Order
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
func (r *Repository) UpdateOrderByID(ctx context.Context, order Order) (Order, error) {
	_, err := r.client.Collection("orders").Doc(order.Id.String()).Set(ctx, order)
	if err != nil {
		return Order{}, err
	}

	order, err = r.GetOrderByID(ctx, order.Id.String())
	if err != nil {
		return Order{}, err
	}

	return order, err
}

// GetProducts returns all products
func (r *Repository) GetProducts(ctx context.Context) ([]Product, error) {
	var products []Product

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

		var product Product
		if err := doc.DataTo(&product); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// PostProduct creates a product
func (r *Repository) PostProduct(ctx context.Context, product Product) (Product, error) {
	_, err := r.client.Collection("products").Doc(product.Id.String()).Set(ctx, product)
	if err != nil {
		return Product{}, err
	}

	product, err = r.GetProductByID(ctx, product.Id.String())
	if err != nil {
		return Product{}, err
	}

	return product, err
}

// GetProductByID returns a single product
func (r *Repository) GetProductByID(ctx context.Context, id string) (Product, error) {
	var product Product
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
func (r *Repository) UpdateProductByID(ctx context.Context, product Product) (Product, error) {
	_, err := r.client.Collection("products").Doc(product.Id.String()).Set(ctx, product)
	if err != nil {
		return Product{}, err
	}

	product, err = r.GetProductByID(ctx, product.Id.String())
	if err != nil {
		return Product{}, err
	}

	return product, err
}
