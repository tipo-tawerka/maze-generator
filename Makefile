COVERAGE_FILE ?= coverage.out

TARGET ?= app # CHANGE THIS TO YOUR BINARY NAME

.PHONY: build
build:
	@echo "Выполняется go build для таргета ${TARGET}"
	@mkdir -p .bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

## test: run all tests
.PHONY: test
test:
	@go test --race -count=1 -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

## lint: run golangci-lint
.PHONY: lint
lint:
	@echo "Выполняется линтинг с golangci-lint"
	@golangci-lint run