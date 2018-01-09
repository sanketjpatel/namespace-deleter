TARGET = namespace-deleter
VERSION = v0.0.2
IMAGE = gcr.io/heptio-images/$(TARGET)
GOTARGET = github.com/heptio/$(TARGET)
DOCKER ?= docker

image: cbuild
	$(DOCKER) build -t $(IMAGE):$(VERSION) .
