PACKAGE  = webquic
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

GO      = go
GODOC   = godoc
GOFMT   = gofmt
GOLINT   = gometalinter

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[0;35m▶\033[0m")


all: check

# Dependencies
.PHONY: deps
deps:
	$(info $(M) building vendor…) @
	$Q dep ensure -vendor-only
	$Q $(MAKE) goquic

# Goquic library
.PHONY: goquic
goquic:
	$(info $(M) building goquic…) @
	$Q cd vendor/github.com/devsisters/goquic &&\
		git clone git@github.com:devsisters/libquic.git libquic &&\
		./build_libs.sh -r

# Check
.PHONY: check
check: lint test

# Tests
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Tools
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…) @
	$Q GO_VENDOR=1 $(GOLINT) "--vendor" \
					"--exclude=.pb.go" \
					"--disable=gotype" \
					"--disable=vetshadow" \
					"--disable=gocyclo" \
					"--fast" \
					"--json" \
					"./..." \

.PHONY: fmt
fmt:
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc:
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: version
version:
	@echo $(VERSION)
