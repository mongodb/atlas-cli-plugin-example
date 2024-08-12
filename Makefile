
CLI_SOURCE_FILES?=./cmd/plugin
CLI_BINARY_NAME=binary
CLI_DESTINATION=./bin/$(CLI_BINARY_NAME)

.PHONY: build
build: ## Generate the binary in ./bin
	@echo "==> Building $(CLI_BINARY_NAME) binary"
	go build -o $(CLI_DESTINATION) $(CLI_SOURCE_FILES)
