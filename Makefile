COMMIT=`git rev-parse --short HEAD`
APP=kappa
REPO?=opsfactory/$(APP)
TAG?=latest
BUILD?=-dev

export GO15VENDOREXPERIMENT=1
export GOPATH:=$(PWD)/vendor:$(GOPATH)

all: build-container build-image

deps:
	@rm -rf Godeps vendor
	@godep save ./...

build: format build-app

build-app:
	@cd cmd/$(APP) && go build -o ../../bin/$(APP) -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT) -X github.com/$(REPO)/version.Build=$(BUILD)" .

build-container:
	@docker build -t kappa-build -f build/Dockerfile.build .
	@docker run -it -e BUILD -e TAG --name kappa-build -ti kappa-build make build
	@docker cp kappa-build:/go/src/github.com/$(REPO)/bin/$(APP) ./bin/$(APP)
	@docker rm -fv kappa-build

build-image:
	@docker build -t $(REPO):$(TAG) -f build/Dockerfile .

format:
	@gofmt -s -w .

test:
	@go test -v -cover -race `go list ./... | grep -v /vendor/`

clean:
	@docker rmi kappa-build
	@rm bin/$(APP)

.PHONY: all deps build build-app build-container build-image format test clean
