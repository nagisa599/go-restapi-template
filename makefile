ROOT_FILE = schema/root.yaml
OPENAPI_FILE = OpenAPI.yaml
export GOBIN := $(PWD)/bin

gen-openapi:
	@echo "Generating OpenAPI spec..."
	swagger-cli bundle -o $(OPENAPI_FILE) -t yaml $(ROOT_FILE)

bl-local:
	@echo "Building local docker image..."
	cd ./build && docker compose build --no-cache

up-local:
	@echo "Starting local docker container..."
	cd ./build && docker compose up

goBackend:
	@echo "Starting go backend..."
	docker exec -it  rest-api-backend bash

gen-go:
	docker exec -it rest-api-backend sh -c "cd app/gen && oapi-codegen --config config.yml ../../OpenAPI.yaml"

b:
	docker exec -it rest-api-backend bash

install:
	mkdir -p ./bin; \
	go install go.uber.org/mock/mockgen@latest
	go install github.com/sanposhiho/gomockhandler@latest; \


mock-gen-example:
	gomockhandler -source=./app/internal/domain/repository/user_repository.go -destination=./app/mock/repository_mock/user_mock.go -package=repository_mock


mock-first:
	gomockhandler -config=gomockhandler.json string





mock-check:
	gomockhandler check