################################################################################
## Dependencies
################################################################################

dependencies-install:
	@./scripts/dependencies/install.sh

dependencies-check:
	@./scripts/dependencies/check.sh

################################################################################
## Code generation
################################################################################

generate: dependencies-check
	@find . -type f -name '*_string.go' -exec rm {} +
	@find . -type d -name 'generated' -exec rm -rf {} +
	@PATH="$(shell pwd)/local/bin:$$PATH" && go generate ./...
	@cd api && go tool buf generate && cd .. && mv api/generated/ internal/adapters/handler/connect/

################################################################################
## Local Development
################################################################################

format:
	@sqlfluff fix .
	@go fmt ./...
	@gofmt -w .
	@goimports -w .
	@./local/bin/golines -w --ignore-generated --no-reformat-tags --chain-split-dots -m $(shell awk '/line-length/ {print $$2}' .golangci.yaml) .

lint:
	@CGO_ENABLED=1 golangci-lint run
	@sqlfluff lint --config .sqlfluff

################################################################################
## Run Server
################################################################################

run-http:
	@go run cmd/http/main.go

################################################################################
## Test Endpoints
################################################################################

test-health:
	@curl \
		--header "Content-Type: application/json" \
		--data '{"service":"petstore.v1.PetStoreService"}' \
		http://localhost:8080/grpc.health.v1.Health/Check

test-list-pets:
	@echo "Listing all pets:"
	@grpcurl \
		-proto api/petstore/v1/rpc.proto \
		-import-path api \
		-plaintext \
		localhost:8080 \
		petstore.v1.PetStoreService/ListPets

################################################################################
## Data Management
################################################################################

seed-pets:
	@echo "Seeding sample pets into the database..."
	@./scripts/seed_pets.sh

