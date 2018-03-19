VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
VERSION_BUILD ?= 0
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)

REPO_PATH := github.com/yolo3301/dumb-flow
APISERVER_PATH := ./cmd/df-apiserver
WORKER_PATH := ./cmd/df-worker
DECIDER_PATH := ./cmd/df-decider
DFCTL_PATH := ./pkg/dfctl
OUTPUT_PATH := ./bin

LDFLAGS := "-X $(REPOPATH)/version.version=$(VERSION)"

# Build the project
all: dep test vet build 

dep:
	dep ensure

test:
	@echo "Test coming soon..."

vet:
	@echo "Vet coming soon..."

build:
	@echo "Building..."
	go build -v -i -o $(OUTPUT_PATH)/df-server $(APISERVER_PATH)
	go build -v -i -o $(OUTPUT_PATH)/df-worker $(WORKER_PATH)
	go build -v -i -o $(OUTPUT_PATH)/df-decider $(DECIDER_PATH)
	go build -v -i -o $(OUTPUT_PATH)/dfctl $(DFCTL_PATH)

.PHONY: test vet build 