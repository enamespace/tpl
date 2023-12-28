.DEFAULT_GOAL := all

.PHONY: all
all: tidy gen format lint cover build

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/tools.mk

.PHONY: tidy
tidy:
	@echo "make tidy"

.PHONY: gen
gen:
	@echo "make gen"

## format: Gofmt (reformat) package sources (exclude vendor dir if existed).
.PHONY: format
format: tools.verify.golines tools.verify.goimports
	@echo "===========> Formating codes"
	@find -type f -name '*.go' | xargs gofmt -s -w
	@find -type f -name '*.go' | xargs goimports -w -local $(ROOT_PACKAGE)
	@find -type f -name '*.go' | xargs golines -w --max-len=120 --reformat-tags --shorten-comments --ignore-generated .
	@$(GO) mod edit -fmt

.PHONY: lint
lint:
	@$(MAKE) go.lint

.PHONY: cover
cover:
	@echo "make cover"

.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: clean
clean:
	@$(MAKE) go.clean

## tools: install dependent tools.
.PHONY: tools
tools:
	@$(MAKE) tools.install

## swagger: Generate swagger document.
.PHONY: swagger
swagger:
	@$(MAKE) swagger.run