# Kibandaa

Kibandaa is a simple e-commerce platform that allows users to buy and sell products.

## Payment service

The Payment service is a REST API that allows you to receive a callback from Mpesa.

### API Documentation

API Documentation is written in OpenAPI3, and is located in the `payments/api` directory of this project. You can copy / import this file into a live editor, such as [Swagger's Online Editor](https://editor.swagger.io/), and see more information about the endpoints, parameters and response types.

### Precursor

Before beginning, the following needs to be installed:

- [Golang](https://go.dev/doc/install)
- [GolangCI-Lint](https://golangci-lint.run/usage/install/)
- `make`

`GO111MODULE=off /usr/local/bin/go get -u -v golang.org/x/tools/cmd/goimports`

### Build and Run

Once the Go and golangci-lint are installed, you should be able to run:

```bash

cd payments

make run

```

### Additional Targets

For your convenience, all necessary operations are centralized into a `Makefile` located in the root of the project. See that file's comments for specific information.

### API Development

This project uses [OpenAPI v3](https://spec.openapis.org/oas/v3.1.0), [Deepmap's OpenAPI Code generator](https://github.com/deepmap/oapi-codegen), and [Gin Framework](https://gin-gonic.com/) to manage API and REST Endpoint development.

Workflow:

1. All changes to any part of the REST API must start from the OpenAPI specs in the projects `orders/api` directory.
2. Run `version=v1 make generate-api` which will call the code generator and overwrite the `api.gen.go` file(s).
3. Make required changes to the code that is implementing the server endpoints.


### Testing the callback endpoint

```bash 

curl --location 'localhost:8000/payments/callback' \
--header 'Content-Type: application/json' \
--data '{
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
}
'

```


## Order service

The order service is a REST API that allows you to create, read, update and delete orders.

### API Documentation

API Documentation is written in OpenAPI3, and is located in the `orders/api` directory of this project. You can copy / import this file into a live editor, such as [Swagger's Online Editor](https://editor.swagger.io/), and see more information about the endpoints, parameters and response types.

### Precursor

Before beginning, the following needs to be installed:
- [Golang](https://go.dev/doc/install)
- [GolangCI-Lint](https://golangci-lint.run/usage/install/)
- `make`

`GO111MODULE=off /usr/local/bin/go get -u -v golang.org/x/tools/cmd/goimports`

### Build and Run

Once the Go and golangci-lint are installed, you should be able to run:

```bash 

cd orders

make run 

```
### Additional Targets

For your convenience, all necessary operations are centralized into a `Makefile` located in the root of the project. See that file's comments for specific information.

### API Development

This project uses [OpenAPI v3](https://spec.openapis.org/oas/v3.1.0), [Deepmap's OpenAPI Code generator](https://github.com/deepmap/oapi-codegen), and [Gin Framework](https://gin-gonic.com/) to manage API and REST Endpoint development.

Workflow:

1. All changes to any part of the REST API must start from the OpenAPI specs in the projects `orders/api` directory.
2. Run `version=v1 make generate-api` which will call the code generator and overwrite the `api.gen.go` file(s).
3. Make required changes to the code that is implementing the server endpoints.

### Testing the endpoints

#### Customers endpoints

```bash 

curl --location 'localhost:8000/customers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "hungaikevin@gmail.com",
    "name": "Hungai Amuhinda",
    "phone": "254724490814"
}'

```

