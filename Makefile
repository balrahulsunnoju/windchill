
all: test

# Tool versions
GOLANGCI_LINT_VERSION?=v1.61.0

.PHONY: test
test:

.PHONY: test-cli
test: test-cli
test-cli:
	go run ./cmd/windchill -table
	tests/table.sh
	tests/bad-prompt.sh
	tests/prompt.sh

.PHONY: test-go
test: test-go
test-go:
	go test -v -timeout=2s ./...

.PHONY: lint
test: lint
lint: tool/golangci-lint
	tool/golangci-lint run

.PHONY: fix
fix: tool/golangci-lint
	tool/golangci-lint run --fix

tool/golangci-lint: tool/.golangci-lint.$(GOLANGCI_LINT_VERSION)
	@mkdir -p tool
	GOBIN="$(PWD)/tool" go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

tool/.golangci-lint.$(GOLANGCI_LINT_VERSION):
	@rm -f tool/.golangci-lint.*
	@mkdir -p tool
	touch $@

.PHONY: tool
tool: tool/golangci-lint
