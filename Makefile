DOCKER_BUF := docker run -t --rm -v $(shell pwd):/workspace -w /workspace bufbuild/buf:1.11.0
DOCKER_GOLANGCI_LINTER := docker run -t --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.50.1

.PHONY: proto
proto:
	@$(DOCKER_BUF) generate

.PHONY: lint
lint:
	@$(DOCKER_GOLANGCI_LINTER) golangci-lint run -v
