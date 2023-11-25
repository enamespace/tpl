SHELL := /bin/bash


GIT_COMMIT:=$(shell git rev-parse HEAD)

COPY_GITHOOKS:=$(shell cp -f githooks/* .git/hooks/)


