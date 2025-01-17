# Makefile
.PHONY: build-web move-dist build-server run clean

APP_NAME = gsCheckApp

build:
	make build-web
	make build-server

build-linux:
	make build-web
	make build-server-linux

build-web:
	cd ./web && npm install && npm run build
	rm -rf ./server/dist/ && mv ./web/dist/ ./server/

build-server:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 go build -ldflags "-s -w" -o ../bin/$(APP_NAME)

build-server-linux:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ../bin/$(APP_NAME)

