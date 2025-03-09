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

