
USERNAME := traliaa
PROJECT :=myFirstApp
APP_NAME := myapp
DOCKER_VERSION := latest

GOPATH?=$(HOME)/go
GIT_COMMIT := $(shell git rev-parse HEAD)
GO_VERSION:=$(shell go version)
GO_VERSION_SHORT:=$(shell echo $(GO_VERSION) | sed -E 's/.* go(.*) .*/\1/g')
BUILD_TS:=$(shell date +%FT%T%z)
GIT_LOG:=$(shell git log --decorate --oneline -n1 2> /dev/null | base64 | tr -d '\n')
GIT_HASH:=$(shell git log --format="%h" -n 1 2> /dev/null)
GIT_BRANCH:=$(shell git branch 2> /dev/null | grep '*' | cut -f2 -d' ')
APP_VERSION:=$(if $(CI_COMMIT_REF_SLUG),$(CI_COMMIT_REF_SLUG),$(if $(GIT_BRANCH),$(GIT_BRANCH),$(GIT_HASH)))

LDFLAGS:=-X 'github.com/Traliaa/go_back_2/version.Name=$(APP_NAME)'\
         -X 'github.com/Traliaa/go_back_2/version.Version=$(APP_VERSION)'\
         -X 'github.com/Traliaa/go_back_2/version.GoVersion=$(GO_VERSION_SHORT)'\
         -X 'github.com/Traliaa/go_back_2/version.BuildDate=$(BUILD_TS)'\
         -X 'github.com/Traliaa/go_back_2/version.GitLog=$(GIT_LOG)'\
         -X 'github.com/Traliaa/go_back_2/version.GitHash=$(GIT_HASH)'\
         -X 'github.com/Traliaa/go_back_2/version.GitBranch=$(GIT_BRANCH)'\


BUILD_ENVPARMS:= CGO_ENABLED=0 GOOS=linux GOARCH=amd64

.PHONY: run
run:
	$(info #Running...)
	$(BUILD_ENVPARMS) go run -ldflags "$(LDFLAGS)" ./cmd/$(APP_NAME)

.PHONY: .build
.build: BIN?=./bin/$(APP_NAME)
.build:
	$(info #Building...)
	$(BUILD_ENVPARMS) go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/$(APP_NAME)


build_container:
		$(info #build_container...)
		 DOCKER_SCAN_SUGGEST=false docker build -t docker.io/$(USERNAME)/$(APP_NAME):$(DOCKER_VERSION) -f deployments/Dockerfile  .

run_container:
	  docker run  docker.io/$(USERNAME)/$(APP_NAME):$(DOCKER_VERSION)


push_container:
	 docker push docker.io/$(USERNAME)/$(APP_NAME):$(DOCKER_VERSION)

deployment_minikube:
	 kubectl apply -f ./deployments/

.PHONY: lint
lint:
	golangci-lint --timeout 5m0s run ./...

.PHONY: install-lint
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1


.PHONY: test
test:
	go test -race  -cover -v ./...
