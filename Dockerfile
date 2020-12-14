# Step 1 to build binary
FROM golang:alpine AS builder

RUN mkdir /app
COPY . /app/
WORKDIR /app

# Set up go proxy env
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"

# Set up requirement
RUN go mod download
RUN apk add make git

# Build exec file
RUN make build

# Step 2 to build a small image
FROM alpine
LABEL maintainer "Linxing <linxing301@gmail.com>"

COPY --from=builder /app/srv /app/worker /app/grpc_srv /app/setting/ /home/

# Service at port 8888 Metric at 23333
EXPOSE 8888 23333

ENTRYPOINT ["/bin/sh", "-c"]
CMD ["exec /home/srv -conf.ini /home/conf.dev.ini"]
