# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.20-alpine as builder

# 设置工作目录
WORKDIR /go/src/github.com/sxz799/gsCheck-server

RUN apk --no-cache add gcc musl-dev

# 将应用的代码复制到容器中
COPY ./server/ .

# 编译应用程序
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env \
    && go mod tidy \
    && go build -o app .

FROM node:16


WORKDIR /gsCheck-web/
COPY ./web/ .


RUN npm config set registry https://registry.npm.taobao.org/ \
    && npm install \
    && npm run build



FROM alpine:latest

WORKDIR /home

COPY --from=0 /go/src/github.com/sxz799/gsCheck-server/app ./
COPY --from=1 /gsCheck-web/dist/ ./dist

EXPOSE 7990

# 运行应用程序
CMD ["./app"]