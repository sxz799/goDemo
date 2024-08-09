# Makefile
.PHONY: build-web move-dist build-server run clean

APP_NAME = gsCheckApp

build:
	make build-web
	make move-dist
	make build-server
	make move-resources

build-linux:
	make build-web
	make move-dist
	make build-server-linux
	make move-resources

build-web:
ifeq ($(shell uname -s), Darwin)
	cd ./web && sed -i '' 's#//sed_tag/publicPath#publicPath#g' vue.config.js && npm install && npm run build && sed -i '' 's#publicPath#//sed_tag/publicPath#g' vue.config.js
else
	cd ./web && sed -i 's#//sed_tag/publicPath#publicPath#g' vue.config.js && npm install && npm run build && sed -i 's#publicPath#//sed_tag/publicPath#g' vue.config.js
endif
move-dist:
	rm -rf ./server/dist/ && mv ./web/dist/ ./server/

build-server:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 go build -ldflags "-s -w" -o $(APP_NAME)

build-server-linux:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(APP_NAME)

move-resources:
	rm -rf ./bin/$(APP_NAME) &&  mv ./server/$(APP_NAME) ./bin/