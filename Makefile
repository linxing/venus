GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go" -type f -not -path "./pkg/*")

# These are the values we want to pass for Version and BuildTime
GITTAG=`git describe --tags --always`
BUILD_TIME=`date +%FT%T`
LDFLAGS=-ldflags "-X main.GitTag=${GITTAG} -X main.BuildTime=${BUILD_TIME}"

SPEC_TEST := $(run)

MARK="\
	 __ | / /  _ \_  __ \  / / /_  ___/\n\
     __ |/ //  __/  / / / /_/ /_(__  )\n\
     _____/ \___//_/ /_/\__,_/ /____/ "

all: compile fmt-check test

.PHONY: compile
compile: ## Compile the proto file.
	@echo "Complie proto file"
	@protoc -I protos/venus/ protos/venus/venus.proto --go_out=plugins=grpc:pkg/
	@echo "Complie proto file success"

.PHONY: fmt-check
fmt-check:
	@echo get all go files and run go fmt on them
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: test
test: ## default run all test
    ifeq (${SPEC_TEST},)
		@echo "Go test with coverage percentage"
		@go test ./... -cover | { grep -v 'no test files'; true; }
    else
		@go test -run ${SPEC_TEST} ./... -cover | { grep -v -e 'no test files' -e 'no tests to run'; true; }
    endif

	@find . -name "*.log" |xargs rm -f

# Example: make docker tag=0.0.3
.PHONY: docker
docker:
	@echo "build docker image"
	@docker build -t venus-$(tag) .

.PHONY: build
build:
	@echo "build srv tag:" ${GITTAG} "build at" ${BUILD_TIME}
	@go build ${LDFLAGS} -o srv cmd/srv/main.go
	@echo "build worker tag:" ${GITTAG} "build at" ${BUILD_TIME}
	@go build ${LDFLAGS} -o worker cmd/worker/main.go
	@echo "build grpc tag:" ${GITTAG} "build at" ${BUILD_TIME}
	@go build ${LDFLAGS} -o grpc_srv cmd/grpc/main.go

.PHONY: doc
doc: ## doc
	@echo "Gen api doc"
	@apidoc -i handler/
	@apidoc-markdown -p doc -o apidoc.md
	@rm -rf doc

.PHONY: run
run: ## run
	@echo "Run Server"
	@apidoc -i handler/
	@apidoc-markdown -p doc -o apidoc.md
	@rm -rf doc
	@go run cmd/srv/main.go -conf.ini setting/conf.dev.ini

.PHONY: worker
worker: ## run
	@echo "Run Worker"
	@go run cmd/worker/main.go -conf.ini setting/conf.dev.ini

.PHONY: grpc
grpc: ## run
	@echo "Run gRpc Server"
	@go run cmd/grpc/main.go -conf.ini setting/conf.dev.ini

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -error .

# Download migrate from https://github.com/golang-migrate/migrate/releases
# export MYSQLDSN="mysql://root:root@tcp(127.0.0.1:3306)/venus"
.PHONY: migrate
migrate:
	@migrate -source file://./migrations/ -database "$(MYSQLDSN)" up

# https://golangci-lint.run/usage/install/#local-installation
.PHONY: lint
lint:
	@golangci-lint run
