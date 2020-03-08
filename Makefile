MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
BIN      = $(CURDIR)/bin
DEPLOY   = $(CURDIR)/scripts/deploy.sh

GO      = go
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

export GO111MODULE=on

.PHONY: all
all: fmt lint | $(BIN) ; $(info $(M) building executable…) @ ## Build program binary
	$Q $(GO) build \
		-tags release \
		-o $(BIN)/$(basename $(MODULE)) server/main.go

# Tools

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(PACKAGE)…)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOLINT = $(BIN)/golint
$(BIN)/golint: PACKAGE=golang.org/x/lint/golint

# Start

.PHONY: start
start: ; $(info $(M) start…) @ ## Start the client
	@URL="$$($(DEPLOY) print_endpoint)"; (cd client; REACT_APP_ENDPOINT=$$URL yarn start)

# Deploy

.PHONY: deploy
deploy: ; $(info $(M) deploy…) @ ## Deploy the application
	$Q $(DEPLOY)

.PHONY: print_endpoint
print_endpoint: ; $(info $(M) print endpoint…) @ ## Print the endpoint
	$Q $(DEPLOY) print_endpoint

.PHONY: delete
delete: ; $(info $(M) delete…) @ ## Delete the application
	$Q $(DEPLOY) delete

# Code

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q $(GOLINT) -set_exit_status $(PKGS)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt $(PKGS)

# Misc

.PHONY: proto
proto: ; $(info $(M) generating service ...)	@ ## Generating gRPC service
	@protoc --go_out=plugins=grpc:./proto --js_out=import_style=commonjs:./client/src/proto --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:./client/src/proto --proto_path=proto proto/*.proto

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf zk-single-kafka-single

.PHONY: help
help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
