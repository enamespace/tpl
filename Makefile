.DEFAULT_GOAL := all

.PHONY: all
all: tidy gen format lint cover build

include scripts/make-rules/common.mk

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

