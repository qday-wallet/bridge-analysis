BUILDDIR ?= $(CURDIR)/build
NAMESPACE := ghcr.io/qday-wallet
PROJECT := bridge-analysis
DOCKER_IMAGE := $(NAMESPACE)/$(PROJECT)
COMMIT_HASH := $(shell git rev-parse --short=7 HEAD)
DATE=$(shell date +%Y%m%d-%H%M%S)
## DOCKER_TAG := ${DATE}-$(COMMIT_HASH)
DOCKER_TAG := 0.1.0
MODULES := $(wildcard api/*)
SYSTEM := $(shell uname -s)

###############################################################################
###                                  Build                                  ###
###############################################################################


image-build:
	docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .

image-build-linux:
	docker build --platform=linux/amd64 -t ${DOCKER_IMAGE}:${DOCKER_TAG} .
	
image-push:
	docker push --all-tags ${DOCKER_IMAGE}

image-list:
	docker images | grep ${DOCKER_IMAGE}

$(MOCKS_DIR):
	mkdir -p $(MOCKS_DIR)

distclean: clean tools-clean

clean:
	rm -rf \
    $(BUILDDIR)/ \
    artifacts/ \
    tmp-swagger-gen/