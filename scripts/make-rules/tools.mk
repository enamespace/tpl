TOOLS ?=$(BLOCKER_TOOLS) $(CRITICAL_TOOLS) $(TRIVIAL_TOOLS)

.PHONY: tools.install
tools.install: $(addprefix tool.install., $(TOOLS))


.PHONY: tools.install.%
tool.install.%:
	@echo "========> Installing $*"
	@$(MAKE) install.$*


.PHONY: install.protoc-gen-go
install.protoc-gen-go:
	@$(GO) install github.com/golang/protobuf/protoc-gen-go@latest

install.mockgen:
	@$(GO) install github.com/golang/mock/mockgen@latest

.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
	@golangci-lint completion bash > $(HOME)/.golangci-lint.bash
	@if ! grep -q .golangci-lint.bash $(HOME)/.bashrc; then echo "source \$$HOME/.golangci-lint.bash" >> $(HOME)/.bashrc; fi

.PHONY: install.go-callvis
install.go-callvis:
	@$(GO) install github.com/ofabry/go-callvis@latest