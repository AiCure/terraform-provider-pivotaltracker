# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=pivotal-tracker
BINARY_DIR=build
BINARY_WIN=$(BINARY_NAME).exe
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_DARWIN=$(BINARY_NAME)_osx

all: test build
build: build-darwin build-win build-linux 
test: gen unit integration
unit: 
	$(GOTEST) ./pkg/... -v
integration: 
	GOCACHE=off $(GOTEST) ./test/integration/... -v
e2e: 
	GOCACHE=off $(GOTEST) ./test/e2e/... -v
clean: 
	$(GOCLEAN)
	find . -name "*.test" | xargs rm 
	rm -fr $(BINARY_DIR)
gen:
	go generate ./...
dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
build-darwin: 
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags "-X main.Buildtime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=${CIRCLE_SHA1} -X main.Platform=OSX/amd64" -v -o $(BINARY_DIR)/$(BINARY_DARWIN) ./cmd/pivotal-tracker
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags "-X main.Buildtime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=${CIRCLE_SHA1} -X main.Platform=Windows/amd64"-v -o $(BINARY_DIR)/$(BINARY_WIN) ./cmd/pivotal-tracker
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-X main.Buildtime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=${CIRCLE_SHA1} -X main.Platform=Linux/amd64"-v -o $(BINARY_DIR)/$(BINARY_UNIX) ./cmd/pivotal-tracker
release:
	./bin/create_new_release.sh

.PHONY: all test clean build