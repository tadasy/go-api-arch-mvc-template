IMAGE_TAG=web

generate-code-from-openapi: ## Generate code from OpenAPI spec
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go tool oapi-codegen -config ./api/config.yaml ./api/openapi.yaml

external-up: ## Start external containers
	docker compose up -d mysql swagger-ui
external-down: ## Stop all external containers
	docker compose down

mysql-cli: ## Run mysql cli
	docker compose run mysql-cli --rm

run: ## Run the app
	export APP_ENV=development
	go run main.go

docker-build: ## Build the docker image
	docker build -t $(IMAGE_TAG) .
docker-run: ## Run the app in a docker container
	docker run -p 8080:8080 -i -t $(IMAGE_TAG)

docker-compose-up: docker-build ## Build and run the app with docker compose
	docker compose up -d --wait mysql web swagger-ui
docker-compose-down: ## Stop all containers
	docker compose down

unittest: ## Run unit tests
	go clean -testcache
	go test -v `go list ./... | grep -v /integration | grep -v /testutils | grep -v /app`
	go test -v -p 1 ./app/...

test-cover: ## Run unit tests with coverage
	go test -coverprofile=coverage.out `go list ./... | grep -v /integration` && go tool cover -html=coverage.out

integration_test: generate-code-from-openapi docker-build ## Run integration tests
	export APP_ENV=integration
	-docker compose up -d --wait
	-go clean-testcache
	-go test -v `go list ./... | grep -v /integration`
	+docker compose down

lint: ## Run lint
	golangcli-lint run

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
