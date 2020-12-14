# Venus
Gin CRUD framework

## Basic Knowledge

### Xorm Doc
    https://www.kancloud.cn/xormplus/xorm/167077
### Gin Doc
    https://www.kancloud.cn/shuangdeyu/gin_book/949411
### RBAC By Casbin
    https://casbin.org/zh-CN/
### gRPC Doc
    https://grpc.io/docs/languages/go/quickstart/

## Requirement

### Install Protobuf
    apt install -y protobuf-compiler
    go get github.com/golang/protobuf/protoc-gen-go@v1.3.2
### Package
    go mod download
### Setup Redis & MySQL In Docker Or (By Manual)
    docker-compose up -d
### API doc
    npm install apidoc -g

    // Do not forget to delete the doc dir
    npm install apidoc-markdown -g

## Test
### Run whole project test
    go test ./...

## Docker Run Worker
    docker run -it --network="host" venus-tag "/app/worker -conf.ini /app/setting/conf.dev.ini"
    docker run -it --network="host" venus-tag "/app/srv -conf.ini /app/setting/conf.dev.ini"
