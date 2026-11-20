# 基础镜像可通过 build-arg 换成镜像站地址，便于国内网络拉取
ARG NODE_IMAGE=node:22-alpine
ARG GO_IMAGE=golang:1.23-alpine
ARG RUNTIME_IMAGE=alpine:3.20

# ---- 前端构建 ----
FROM ${NODE_IMAGE} AS web
WORKDIR /app/web
COPY web/package.json web/package-lock.json ./
# 国内网络可传 --build-arg NPM_REGISTRY=https://registry.npmmirror.com
ARG NPM_REGISTRY=
RUN if [ -n "$NPM_REGISTRY" ]; then npm config set registry "$NPM_REGISTRY"; fi \
 && npm ci --no-audit --no-fund
COPY web/ ./
RUN npm run build

# ---- Manager 演示子应用（官网 /console 直接复用 Manager 前端） ----
FROM ${NODE_IMAGE} AS console_demo
ARG MANAGER_REPO=https://github.com/Lxy417165709/auto_sjz_manager.git
ARG MANAGER_REF=main_dev
ARG NPM_REGISTRY=
RUN apk add --no-cache git \
 && git clone --depth 1 --branch ${MANAGER_REF} ${MANAGER_REPO} /manager
WORKDIR /manager/frontend
RUN if [ -n "$NPM_REGISTRY" ]; then npm config set registry "$NPM_REGISTRY"; fi \
 && npm ci --no-audit --no-fund \
 && npm run build:demo

# ---- Go 构建 ----
FROM ${GO_IMAGE} AS server
WORKDIR /app/server
COPY server/go.mod ./
RUN go mod download
COPY server/ ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /out/site .

# ---- 运行 ----
FROM ${RUNTIME_IMAGE}
RUN apk add --no-cache ca-certificates tzdata \
 && adduser -D -u 10001 site \
 && mkdir -p /app /data \
 && chown -R site:site /app /data

WORKDIR /app
COPY --from=server /out/site /app/site
COPY --from=web /app/web/dist /app/web
# 用 Manager 仓库现场构建的 demo 覆盖静态兜底产物
COPY --from=console_demo /manager/frontend/dist-demo /app/web/console-demo
COPY server/content /app/content

USER site
ENV PORT=8080 \
    WEB_DIR=/app/web \
    CONTENT_DIR=/app/content \
    LEAD_STORE_PATH=/data/leads.jsonl \
    LOG_LEVEL=info

EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://127.0.0.1:8080/api/health || exit 1

ENTRYPOINT ["/app/site"]
