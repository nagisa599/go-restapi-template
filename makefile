ROOT_FILE = app/schema/root.yaml
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
	go install github.com/golang/mock/mockgen@v1.6.0; \


mock-gen:
	$(GOBIN)/mockgen -source=./app/internal/domain/repository/user_repository.go -destination=./app/mock/repository_mock/mock.go -package=repository_mock
$(GOBIN)/mockgen -source=./app/internal/domain/repository/user_repository.go -destination=./app/mock/repository_mock/mock.go -package=repository_mock
	$(GOBIN)/mockgen -source=./app/internal/usecase/user_usecase.go -destination=./app/mock/usecase_mock/user_mock.go -package=usecase_mock

