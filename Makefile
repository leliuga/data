GO_CMD      ?= CGO_ENABLED=$(CGO_ENABLED) GOOS=$(OS) GOARCH=$(ARCH) go

.PHONY: default
default: update

.PHONY: update
update:
	$(GO_CMD) get -t -v -d -u ./...
	$(GO_CMD) mod tidy
	@echo "Dependencies updated"