SHELL=PATH='$(PATH)' /bin/sh

GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'

PLATFORM := $(shell uname -o)

COMMIT := $(shell git rev-parse HEAD)
VERSION ?= $(shell git describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
ROOT := github.com/ninjahome/web-bridge/util
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME)

NAME := dessage-twitter.exe
OS := windows

ifeq ($(PLATFORM), Msys)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else ifeq ($(PLATFORM), Cygwin)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else
	INCLUDE := $(GOPATH)
	NAME=dessage-twitter
	OS=linux
endif


JS_DIR := assets/js
JS_FILES := $(shell find $(JS_DIR) -name '*.js' ! -name '*.min.js')

.SECONDEXPANSION:

.PHONY: all
.PHONY: pbs
.PHONY: test
.PHONY: contract
.PHONY: js

BINDIR=./bin

all: pbs sol build

build:
	GOOS=$(OS) GOARCH=amd64 $(GOBUILD) -o $(BINDIR)/$(NAME)

pbs:
	cd pbs/ && $(MAKE)

sol:
	cd blockchain/ && $(MAKE)

target:=mac

tar: js
	tar -czvf  $(BINDIR)/assets.tar.gz assets

mac:
	GOOS=darwin go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).mac  -ldflags="$(LD_FLAGS)"
arm:
	CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARM=7 GOARCH=arm64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).arm  -ldflags="$(LD_FLAGS)"
linux:
	GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).lnx  -ldflags="$(LD_FLAGS)"
win:
	GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).exe  -ldflags="$(LD_FLAGS)"

js: $(JS_FILES)
	@for file in $(JS_FILES); do \
		minified=$${file%.js}.min.js; \
		echo "Compressing $$file to $$minified"; \
		terser $$file -o $$minified --compress --mangle; \
	done

clean:
	rm $(BINDIR)/$(NAME).*  $(BINDIR)/assets.tar.gz
