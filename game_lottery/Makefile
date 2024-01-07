SHELL=PATH='$(PATH)' /bin/sh

GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'

PLATFORM := $(shell uname -o)

COMMIT := $(shell git rev-parse HEAD)
VERSION ?= $(shell git describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
ROOT := github.com/ninjahome/web-bridge/blockchain/game_lottery
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME)

NAME := game_lottery.exe
OS := windows

ifeq ($(PLATFORM), Msys)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else ifeq ($(PLATFORM), Cygwin)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else
	INCLUDE := $(GOPATH)
	NAME=game_lottery
	OS=linux
endif

# enable second expansion
.SECONDEXPANSION:

.PHONY: all
.PHONY: pbs
.PHONY: test
.PHONY: contract

BINDIR=$(INCLUDE)/bin

all: pbs ../blockchain/sol build

build:
	GOOS=$(OS) GOARCH=amd64 $(GOBUILD) -o $(BINDIR)/$(NAME)

pbs:
	cd pbs/ && $(MAKE)

sol:
	cd blockchain/sol/ && $(MAKE)

target:=mac

mac:
	GOOS=darwin go build -ldflags '-w -s' -o $(NAME).mac  -ldflags="$(LD_FLAGS)"
arm:
	CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARM=7 GOARCH=arm64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).arm  -ldflags="$(LD_FLAGS)"
linux:
	GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $(NAME).lnx  -ldflags="$(LD_FLAGS)"
win:
	GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).exe  -ldflags="$(LD_FLAGS)"

clean:
	rm $(BINDIR)/$(NAME)
