# Makefile with Golang and Docker/Docker Compose commands

BINDIR      := $(CURDIR)/bin
BINNAME     ?= glabby

COMMIT   := $(shell git rev-parse --short HEAD)
BUILT_AT := $(shell date +%FT%T%z)
BUILT_ON := $(shell hostname)

LDFLAGS := -w -s
LDFLAGS += -X main.commit=${COMMIT}
LDFLAGS += -X main.builtAt=${BUILT_AT}
LDFLAGS += -X main.builtBy=${USER}
LDFLAGS += -X main.builtOn=${BUILT_ON}

# Docker/Docker Compose Shortcuts

ifeq ($(OS), Windows_NT)
	DOCKER_CONTAINER_LIST = $(shell docker ps -aq)
else
	DOCKER_CONTAINER_LIST = $(shell docker ps -aq)
endif

.PHONY: dsp
dsp:
	@-docker system prune -af

.PHONY: dvp
dvp:
	@-docker volume prune -f

.PHONY: dnp
dnp:
	@-docker network prune -f

.PHONY: ds
ds:
	$(if $(strip $(DOCKER_CONTAINER_LIST)), docker stop $(DOCKER_CONTAINER_LIST))

.PHONY: dv
dv:
	$(if $(strip $(DOCKER_CONTAINER_LIST)), docker rm $(DOCKER_CONTAINER_LIST))

.PHONY: clean
clean: ds dv dvp dnp

.PHONY: remove
remove: ds dv dvp dnp dsp

.PHONY: dcbn
dcbn:
	docker-compose build --no-cache

.PHONY: dcub
dcub:
	docker-compose up --build

.PHONY: dcubd
dcubd:
	docker-compose up --build -d

.PHONY: dcs
dcs:
	docker-compose down

.PHONY: dcps
dcps:
	docker-compose ps

.PHONY: run
run: dcps dcs dcubd
ifneq ("$(wildcard $(./.env-exemple))","")
    dcps dcs dcubd
endif

# GoLang Shortcuts

.PHONY: all
all: gbuild

.PHONY: gbuild
gbuild: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BINNAME) main.go

.PHONY: gclean
gclean:
	go clean
	rm -f ./bin/*

.PHONY: glint
glint: gbuild
	golint -set_exit_status ./...

all: gbuild
