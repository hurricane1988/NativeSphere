# Setup the output name of binary file.
BIN_NAME=native-sphere

# App Version
APP_VERSION = v0.1

# Output directory.
OUTPUT_DIR=bin

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: all
all: check native-sphere; $(info $(M)...Begin to check and build all of binary) @ ## check and build all of binary.

help:
	@grep -hE '^[ a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: binary
# Build the native-sphere binary
native-sphere: ; $(info $(M)...Begin to build native-sphere binary.)  @ ## Build native-sphere.
	hack/gobuild.sh cmd/ks-apiserver;
