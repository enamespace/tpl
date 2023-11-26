SHELL := /bin/bash


GIT_COMMIT:=$(shell git rev-parse HEAD)

COPY_GITHOOKS:=$(shell cp -f githooks/* .git/hooks/)

# BLOCKER_TOOLS: ci will fail if miss
# CRITICAL_TOOLS: nessecery operation fail if miss
# TRIVIAL_TOOLS: optional tools, missing these tool have no affect.
BLOCKER_TOOLS ?= golangci-lint
CRITICAL_TOOLS ?= protoc-gen-go mockgen
TRIVIAL_TOOLS ?= go-callvis

