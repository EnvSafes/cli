# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GODEV = air

# Build target
BINARY_NAME = envsafes

ifeq ($(firstword $(MAKECMDGOALS)),$(filter dev,$(MAKECMDGOALS)))
  # use the rest as arguments for "run" or "dev"
  TARGET := $(firstword $(MAKECMDGOALS))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

UNAME_S := $(shell uname -s)
ifeq ($(findstring MINGW,$(UNAME_S)),MINGW)
    OS := windows
    BINARY_NAME := $(BINARY_NAME).exe
else ifeq ($(findstring CYGWIN,$(UNAME_S)),CYGWIN)
    OS := windows
    BINARY_NAME := $(BINARY_NAME).exe
else ifeq ($(findstring MSYS,$(UNAME_S)),MSYS)
    OS := windows
    BINARY_NAME := $(BINARY_NAME).exe
else ifeq ($(UNAME_S),Linux)
    OS := linux
else ifeq ($(UNAME_S),Darwin)
    OS := darwin
else
    $(error Unsupported OS: $(UNAME_S))
endif

build:
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=amd64 $(GOBUILD) -o ./tmp/$(BINARY_NAME) -v

test:
	$(GOTEST) -v 

clean:
	$(GOCLEAN)
	rm -f ./tmp/$(BINARY_NAME)

dev:
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=amd64 $(GOBUILD) -o ./tmp/$(BINARY_NAME) -v
	./tmp/$(BINARY_NAME) $(RUN_ARGS)

.PHONY: build test clean dev