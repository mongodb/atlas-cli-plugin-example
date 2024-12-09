
CLI_SOURCE_FILES?=./cmd/plugin
CLI_BINARY_NAME?=binary
CLI_DESTINATION=./bin/$(CLI_BINARY_NAME)
MANIFEST_FILE?=manifest.yml
WIN_MANIFEST_FILE?=manifest.windows.yml

.PHONY: build
build: ## Generate the binary in ./bin
	@echo "==> Building $(CLI_BINARY_NAME) binary"
	go build -o $(CLI_DESTINATION) $(CLI_SOURCE_FILES)

.PHONY: generate-all-manifests
generate-all-manifests: generate-manifest generate-manifest-windows

.PHONY: generate-manifest
generate-manifest: ## Generate the manifest file for non-windows OSes
	@echo "==> Generating non-windows manifest file"
	printenv
	BINARY=$(CLI_BINARY_NAME) envsubst < manifest.template.yml > $(MANIFEST_FILE)

.PHONY: generate-manifest-windows
generate-manifest-windows: ## Generate the manifest file for windows OSes
	@echo "==> Generating windows manifest file"
	printenv
	CLI_BINARY_NAME="${CLI_BINARY_NAME}.exe" MANIFEST_FILE="$(WIN_MANIFEST_FILE)" $(MAKE) generate-manifest

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'