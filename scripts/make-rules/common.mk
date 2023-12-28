SHELL := /bin/bash
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../.. && pwd -P))
ROOT_PACKAGE=github.com/enamespace/tpl
BUILD_TIME=$(shell date -u +'%Y-%m-%d_%H:%M:%S_GMT')
LEFT_BRACKET=(
BRANCH=$(shell git branch | grep "*" | awk '{print $$2}' | cut -d "$(LEFT_BRACKET)" -f 2)
GIT_COMMIT=$(shell git log --max-count=1 --pretty="format:%h" |cut -c1-7)

COPY_GITHOOKS:=$(shell cp -f githooks/* .git/hooks/)

# BLOCKER_TOOLS: ci will fail if miss
# CRITICAL_TOOLS: nessecery operation fail if miss
# TRIVIAL_TOOLS: optional tools, missing these tool have no affect.
BLOCKER_TOOLS ?= golangci-lint
CRITICAL_TOOLS ?= protoc-gen-go mockgen
TRIVIAL_TOOLS ?= go-callvis

