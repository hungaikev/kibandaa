test:
	@go test -v ./...

tidy:  ## Get the dependencies
	@go mod tidy

compile: tidy  ## compiles orders service code
	@go build ./...

test-race-cond:
	@go test -v -race ./...

run: tidy ## Run the api server
	@go run cmd/orders/*.go

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic
	@cat cover.out >> coverage.txt

check: ## runs code linting and formatting
	@golangci-lint run --disable=typecheck ./...

# inner target to support generating endpoint code.
_install-codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
.PHONY: _install-codegen

# Updates the generated API Endpoint code based on the API spec for the version provided. Version is specified using an
# environment variable. eg. `version=v1 make generate-api`
generate-api: _install-codegen
	APIFILE=$(PWD)/api/rest/$(version)/orders.yaml go generate ./internal/api/rest/$(version)/...
.PHONY: generate

.PHONY: fix
fix:
	for file in `golangci-lint --max-same-issues=1000 --max-issues-per-linter=0 run ./...|grep 'goimports'|cut -f 1 -d:`; do	\
		goimports -local "github.com/hungaikev/kibandaa" -w $$file;	\
	done


help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: compile test test-race-cond check tidy run test-coverage help
