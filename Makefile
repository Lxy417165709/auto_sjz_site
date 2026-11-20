SHELL := /bin/sh

BIN := bin/site
WEB_DIR := web/dist

.PHONY: help bootstrap build build-web build-server dev-server dev-web run fmt vet clean docker

help:
	@echo "make bootstrap    安装前端依赖"
	@echo "make build        构建前端产物与 Go 二进制"
	@echo "make dev-server   启动 Go（8080，WEB_DIR=../web/dist）"
	@echo "make dev-web      启动 Vite（5173，/api 代理到 8080）"
	@echo "make run          构建后在仓库根目录启动"
	@echo "make docker       构建镜像 auto_sjz_site:local"

bootstrap:
	cd web && npm ci

build-web:
	cd web && npm run build

build-server:
	cd server && CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o ../$(BIN) .

build: build-web build-server

dev-server:
	cd server && WEB_DIR=../web/dist go run .

dev-web:
	cd web && npm run dev

run: build
	WEB_DIR=$(WEB_DIR) CONTENT_DIR=server/content PORT=$${PORT:-8080} $(BIN)

fmt:
	cd server && gofmt -l -w .

vet:
	cd server && go vet ./...
	cd web && npm run typecheck

docker:
	docker build -t auto_sjz_site:local .

clean:
	rm -rf $(BIN) web/dist
