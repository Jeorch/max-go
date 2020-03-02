# builder 源镜像
FROM golang:1.12.4-alpine as builder

# 安装git
RUN apk add --no-cache git

ENV GOPROXY="https://goproxy.cn"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

# prod 源镜像
FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=0 /app/max-go .
COPY resource/* ./resource/
COPY resource-public/* ./resource-public/

EXPOSE 9001
ENTRYPOINT ["/app/max-go"]
