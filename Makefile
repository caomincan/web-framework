ROOT_DIR    = $(shell pwd)
# change here for app name
APP_NAME = "web-framework"
MODULE_NAME := $(shell head -1 go.mod | awk '{print $$2}')
PACKAGES = $(shell find . -type d -not -path './.git*' -not -path './vendor*')
VERSION := $(shell git describe --tags --dirty 2>/dev/null)
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
TIME := $(shell date +"%Y-%m-%d.%H:%M:%S")

#  compile target can be linux x86_64
GOOS ?= $(shell uname -s | tr [:upper:] [:lower:])
GOARCH ?= $(shell go env GOARCH)

SETTINGS := $(MODULE_NAME)/internal/controller/version
FLAGS = -ldflags '-s -w -X $(SETTINGS).commit=$(COMMIT) -X $(SETTINGS).version=$(VERSION) -X $(SETTINGS).date=$(TIME)' -trimpath
app:
	@echo [Build] version=$(VERSION) commit=$(COMMIT) date=$(TIME)
	@echo [Build] GOOS=$(GOOS) GOARCH=$(GOARCH)
	@echo [Build] settings=$(SETTINGS)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(FLAGS) -o $(APP_NAME) main.go

unittest:
	@bash unittest.sh

clean:
	@rm -rf $(ROOT_DIR)/$(APP_NAME)

.PHONY: app