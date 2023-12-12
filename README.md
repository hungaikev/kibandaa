# Kibandaa

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


## Product service

The product service is a REST API that allows you to create, read, update and delete products.

### API Documentation

API Documentation is written in OpenAPI3, and is located in the `products/api` directory of this project. You can copy / import this file into a live editor, such as [Swagger's Online Editor](https://editor.swagger.io/), and see more information about the endpoints, parameters and response types.

### Precursor

Before beginning, the following needs to be installed:

- [Golang](https://go.dev/doc/install)
- [GolangCI-Lint](https://golangci-lint.run/usage/install/)
- `make`

`GO111MODULE=off /usr/local/bin/go get -u -v golang.org/x/tools/cmd/goimports`

### Build and Run

Once the Go and golangci-lint are installed, you should be able to run:

```bash

cd products

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

