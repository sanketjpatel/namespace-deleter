TARGET = namespace-deleter
VERSION = v0.0.1
IMAGE = gcr.io/heptio-images/$(TARGET)
GOTARGET = github.com/heptio/$(TARGET)
DOCKER ?= docker
DIR := ${CURDIR}

GO_BUILDMNT = /go/src/$(GOTARGET)
GO_BUILD_IMAGE ?= golang:1.9

BUILDCMD = CGO_ENABLED=0 go build -installsuffix cgo -o $(TARGET) -v
BUILD = $(BUILDCMD) $(GOTARGET)

.PHONY: all image cbuild

all: image

image: cbuild
	$(DOCKER) build -t $(IMAGE):$(VERSION) .

cbuild:
	$(DOCKER) run --rm --volume $(DIR):$(GO_BUILDMNT) --workdir $(GO_BUILDMNT) $(GO_BUILD_IMAGE) /bin/sh -c '$(BUILD)'
