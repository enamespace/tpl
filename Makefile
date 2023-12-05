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

.PHONY: format
format:
	@echo "make format"

.PHONY: lint
lint:
	@echo "make lint"

.PHONY: cover
cover:
	@echo "make cover"

.PHONY: build
build:
	@echo "make build"


## tools: install dependent tools.
.PHONY: tools
tools:
	@$(MAKE) tools.install

## swagger: Generate swagger document.
.PHONY: swagger
swagger:
	@$(MAKE) swagger.run