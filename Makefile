.PHONY: help
.DEFAULT_GOAL := help

OS_LIST=windows linux darwin
ARCH_LIST=amd64 arm64

test: ## Run all test
	@go test ./... -cover -v -race

build: ## Build binary file
	@go build -o gfrt main.go

cross-compile: ## Build binaries for Windows, Linux, and macOS of AMD64 and ARM64
	@mkdir -p bin
	@for GOOS in ${OS_LIST}; do \
		for GOARCH in ${ARCH_LIST}; do \
			if [ $$GOOS = windows ]; then \
				GOOS=$$GOOS GOARCH=$$GOARCH go build -o bin/gfrt-$$GOOS-$$GOARCH.exe main.go; \
			else \
				GOOS=$$GOOS GOARCH=$$GOARCH go build -o bin/gfrt-$$GOOS-$$GOARCH main.go; \
			fi \
		done \
	done

clean: ## Delete all atrifacts (no warning!)
	@-rm -rf bin/

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
