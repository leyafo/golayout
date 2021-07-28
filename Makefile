.PHONY: build

SHELL:=/bin/sh
PROJECT_NAME := golayout
DATETIME = $(shell date '+%Y%m%d_%H%M%S')
PROTOS = `ls proto`

# Path Related
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}bin
DOCKER_TAG := zeroone/opsnft

# Version
RELEASE?=1.0.1
ifndef GIT_COMMIT
  GIT_COMMIT := git-$(shell git rev-parse --short HEAD)
endif

GIT_REPO_INFO=$(shell git config --get remote.origin.url)
ifndef GIT_COMMIT
  GIT_COMMIT := git-$(shell git rev-parse --short HEAD)
endif

# Build Flags
GO_LD_FLAGS= "-s -w -X ${PROJECT_NAME}/pkg/version.RELEASE=${RELEASE} -X ${PROJECT_NAME}/pkg/version.COMMIT=${GIT_COMMIT} -X ${PROJECT_NAME}/pkg/version.REPO=${GIT_REPO_INFO} -X ${PROJECT_NAME}/pkg/version.BUILDTIME=${DATETIME} -X ${PROJECT_NAME}/pkg/version.SERVICENAME=$@"
CGO_SWITCH := 0

build: restful

restful:
	cd ${MKFILE_DIR} && \
	CGO_ENABLED=${CGO_SWITCH} go build -v -trimpath -ldflags ${GO_LD_FLAGS} \
	-o ${RELEASE_DIR}/$@ ${MKFILE_DIR}cmd/$@/

docker_build:
	docker build -t ${DOCKER_TAG}:${RELEASE} -f deployment/Dockerfile . --network=host

docker_clean:
	 docker images | grep \<none\> | awk '{print  $$3}' |xargs docker rmi -f #clean <none>

clean:
	@rm -f ${MKFILE_DIR}bin/*

up:
	docker-compose -f deployment/docker-compose-local.yaml up -d

down:
	docker-compose -f deployment/docker-compose-local.yaml down
