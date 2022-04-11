GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=simple_service
LINTER=golangci-lint

all: 
	test build

test:
	go test ./...
	  

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

lint:
	$(LINTER) run

# test:
# 	go test ./...

# test_coverage:
# 	go test ./... -coverprofile=coverage.out

# dep:
# 	go mod download

# vet:
# 	go vet

# lint:
# 	golangci-lint run --enable-all