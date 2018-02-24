VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
VERSION_BUILD ?= 0
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)

REPO_PATH := github.com/yolo3301/dumb-flow
APISERVER_PATH := ./cmd/df-apiserver
WORKER_PATH := ./cmd/df-worker
DECIDER_PATH := ./cmd/df-decider
OUTPUT_PATH := ./bin

LDFLAGS := "-X $(REPOPATH)/version.version=$(VERSION)"

# Build the project
all: test vet build 

test:
	@echo "Test coming soon..."

vet:
	@echo "Vet coming soon..."

build:
	@echo "Building " $(APISERVER_PATH)
	@echo "Building " $(WORKER_PATH)
	@echo "Building " $(DECIDER_PATH)

.PHONY: test vet build 