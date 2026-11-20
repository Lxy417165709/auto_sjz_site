# auto_sjz_site

auto_sjz 的独立介绍官网：**Go + Vue 3 静态站点**，负责产品介绍、按号收费说明、部署接入指引与试用申请。

- 不接 Manager 登录、数据库与订单；登录与购买统一跳转到 Manager。
- 独立仓库、独立部署，升级官网不影响设备接入与业务执行。
- 站点文案集中在 `server/content/*.json`，改文案不用改前端代码。

## 技术栈

| 层 | 选型 | 说明 |
| --- | --- | --- |
| 前端 | Vue 3 + Vite + TypeScript + Tailwind CSS + Vue Router | 单页应用，history 路由 |
| 后端 | Go 1.23（标准库 `net/http`，零第三方依赖） | 提供页面静态资源、站点内容 API、试用表单与 sitemap |
| 部署 | Docker 多阶段构建 | 单镜像包含二进制 + 前端产物 + 文案 |

## 目录结构

```
.
├── server/                 # Go 服务
│   ├── main.go             # 启动、优雅停机
│   ├── internal/config/    # 环境变量配置
│   ├── internal/content/   # 加载站点文案 JSON
│   ├── internal/httpapi/   # 路由、API、表单、robots/sitemap
│   └── content/            # 站点文案（site/pricing/faq/docs）
├── web/                    # Vue 前端
│   ├── src/pages/          # 首页 / 业务能力 / 按号收费 / 技术架构 / 部署接入 / 文档 / FAQ / 联系
│   ├── src/components/     # 导航、页脚、CTA、架构图等
│   ├── src/lib/            # API 客户端、SEO、类型
│   └── src/router/         # 路由与页面 meta
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

## 页面

| 路由 | 页面 | 内容来源 |
| --- | --- | --- |
| `/` | 首页（业务视角） | 固定文案 + 管理台示意 |
| `/features` | 业务能力 | 固定文案 + 按设备/按号对比 |
| `/pricing` | 按号收费 | `/api/site` 的 pricing |
| `/start` | 接入流程 | 固定文案：设备接入 / 添加账号 / 添加模块 |
| `/tech` | 技术架构 | Manager 树形拓扑图 + 模块职责 + 授权链路 + 技术栈 |
| `/console` | 管理台演示 | 可交互的业务配置 / 业务状态示意（纯前端演示数据） |
| `/deploy` | 部署接入 | 固定文案 + 端口表 |
| `/docs` | 文档（业务 / 开发分组） | `/api/site` 的 docs（`audience` 字段分组） |
| `/faq` | FAQ | `/api/site` 的 faq |
| `/contact` | 联系 / 咨询 | 表单提交 `/api/lead` |

## 本地开发

```bash
make bootstrap   # 安装前端依赖
make dev-server  # 终端 1：启动 Go，默认 8080
make dev-web     # 终端 2：启动 Vite，默认 5173，/api 已代理到 8080
```

访问 `http://localhost:5173`。Vite 开发时前端来自 5173，后端 API 自动代理，无需 CORS 配置。

## 构建与运行

```bash
make build                                                  # 构建 web/dist + bin/site
WEB_DIR=web/dist CONTENT_DIR=server/content bin/site         # 在仓库根目录运行，默认 8080
```

或：

```bash
cd web && npm run build && cd ..
cd server && WEB_DIR=../web/dist go run .
```

## 环境变量

| 变量 | 默认值 | 说明 |
| --- | --- | --- |
| `PORT` | `8080` | 官网监听端口（与 Manager 的 5001、discovery 的 5003 不冲突） |
| `WEB_DIR` | `../web/dist` | 前端产物目录，容器内为 `/app/web` |
| `CONTENT_DIR` | `content` | 站点文案目录，容器内为 `/app/content` |
| `SITE_BASE_URL` | `http://localhost:8080` | 生成 sitemap 与 canonical 的站点地址，上线必须改成真实域名 |
| `MANAGER_URL` | 空 | 「登录 / 购买」跳转的 Manager 地址；为空时按钮退回 `/contact` |
| `LEAD_STORE_PATH` | `data/leads.jsonl` | 试用申请落盘文件 |
| `LEAD_WEBHOOK_URL` | 空 | 可选，表单提交后异步转发到该地址（飞书机器人、内部工单等） |
| `LOG_LEVEL` | `info` | `debug` / `info` / `warn` / `error` |

## 部署

### Docker Compose（推荐）

```bash
SITE_BASE_URL=https://auto-sjz.example.com MANAGER_URL=https://manager.example.com docker compose up -d --build
```

- 容器内已有前端产物，**不需要**外部 Nginx 转发静态文件；如需 TLS，在前面加一层反向代理即可。
- `./data` 挂载到 `/data`，试用申请以 JSONL 追加写入，便于导出或清洗。
- 健康检查：`GET /api/health`。

### Windows（Docker Desktop）

在仓库根目录用 PowerShell 执行，二选一：

```powershell
# 方式一：compose（推荐，自动建 data 目录、自动重启）
Copy-Item .env.example .env
docker compose up -d --build
docker compose logs -f site
```

```powershell
# 方式二：单容器，改端口或复用已有镜像时更方便
docker build -t auto_sjz_site:local .
docker run -d --name auto_sjz_site -p 8080:8080 -e "SITE_BASE_URL=http://localhost:8080" -v "${PWD}\data:/data" auto_sjz_site:local
```

也可以直接跑脚本（默认 8080，读 `MANAGER_URL` 环境变量）：

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\run-local.ps1
```

打开 <http://localhost:8080> 即可。常用排查：

- `docker ps` 看容器是否在跑，`docker logs auto_sjz_site` 看端口占用或启动报错。
- 宿主机 8080 被占用时改映射：`docker run -p 8090:8080 ...`，同时把 `SITE_BASE_URL` 改成 `http://localhost:8090`。
- 容器内固定监听 8080，`PORT` 环境变量只影响容器内部，映射端口以 `-p 宿主:8080` 为准。
- 国内拉 npm 包慢时：`docker build --build-arg NPM_REGISTRY=https://registry.npmmirror.com -t auto_sjz_site:local .`（compose 可先设 `$env:NPM_REGISTRY`）。
- 若 Dockerfile 报 `$'\r'` 一类错误，说明检出时被转成了 CRLF，执行 `git config core.autocrlf false` 后重新检出。
- 不挂载 `data` 也能跑，但试用申请写在容器内，容器删了就没了。

#### 基础镜像拉取慢

`node:22-alpine`、`golang:1.23-alpine`、`alpine:3.20` 都要从 Docker Hub 拉，国内经常卡在几十 MB。三种办法，任选其一：

1. 给 Docker Desktop 配镜像加速：Settings → Docker Engine，加上（可用性会变，逐个试）：

```json
{
  "registry-mirrors": ["https://docker.m.daocloud.io", "https://docker.1ms.run"]
}
```

Apply & Restart 后先单独验证：`docker pull node:22-alpine`。

2. 不改 daemon，直接用镜像站地址构建（Dockerfile 已支持这三个 build-arg）：

```
docker build --build-arg NODE_IMAGE=docker.m.daocloud.io/library/node:22-alpine --build-arg GO_IMAGE=docker.m.daocloud.io/library/golang:1.23-alpine --build-arg RUNTIME_IMAGE=docker.m.daocloud.io/library/alpine:3.20 -t auto_sjz_site:local .
```

3. 先拉镜像站版本再打回官方 tag，后续 build 直接命中本地缓存：

```
docker pull docker.m.daocloud.io/library/node:22-alpine; docker tag docker.m.daocloud.io/library/node:22-alpine node:22-alpine
docker pull docker.m.daocloud.io/library/golang:1.23-alpine; docker tag docker.m.daocloud.io/library/golang:1.23-alpine golang:1.23-alpine
docker pull docker.m.daocloud.io/library/alpine:3.20; docker tag docker.m.daocloud.io/library/alpine:3.20 alpine:3.20
```

有本地代理时也可以走 Docker Desktop 代理：Settings → Resources → Proxies → Manual，HTTP/HTTPS 填 `http://127.0.0.1:7890`（换成你自己的端口），Apply & Restart。

镜像只需拉一次，之后的构建会复用本地层；npm 依赖慢另见上面的 `NPM_REGISTRY`。

### 独立部署要点

- 官网与 Manager、discovery 完全解耦，可以部署在任意公网机器。
- 官网端口仅需对公网开放；不要复用 Manager 的 5001 或 discovery 的 5003。
- 上线前必须配置 `SITE_BASE_URL`，否则 sitemap 里的链接会指向 localhost。

## 站点内容维护

文案全部在 `server/content/`：

| 文件 | 内容 |
| --- | --- |
| `site.json` | 品牌名、导航、CTA、联系方式、页脚、备案号 |
| `pricing.json` | 按号收费的规则、包含项、说明；当前单价 3 元/号/月，改 `unitPriceCny` 即可调价 |
| `faq.json` | FAQ 列表（按 `group` 分组） |
| `docs.json` | 文档正文（分节 + 段落 + 列表） |

改完 JSON 重启服务即可生效（Go 启动时加载，解析失败会直接启动失败，避免线上空白站）。

### 待替换的占位项

- `site.json` 的 `contact.email`（当前为 `contact@auto-sjz.example`）
- `site.json` 的 `icp`（备案号）
- 环境变量 `MANAGER_URL`、`SITE_BASE_URL`

## API

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/health` | 健康检查 |
| GET | `/api/site` | 一次性返回 site + pricing + faq + docs（前端只请求这一个） |
| GET | `/api/site/pricing` | 仅计费内容 |
| GET | `/api/site/faq` | 仅 FAQ |
| GET | `/api/site/docs` | 仅文档 |
| POST | `/api/lead` | 试用申请，字段 `name`、`email`、`phone`、`company`、`devices`、`message`；`website` 为 honeypot |

`/api/lead` 单 IP 限流 5 次 / 10 分钟，进程内实现；多副本部署时建议在网关统一限流。

## SEO

- 每个路由有独立 title / description，`router.afterEach` 同步 canonical 与 og 标签。
- Go 侧提供 `/robots.txt` 与 `/sitemap.xml`，sitemap 由 `site.json` 的导航自动生成。
- 当前为 SPA，首屏内容依赖 JS；若后续需要更强的收录效果，可在此基础上接预渲染（如 `vite-plugin-prerender`）或把文档页改为服务端直出。

## 与业务系统的边界

- Manager 是配置唯一权威源，Auth 是授权时效唯一权威：官网只做说明，不参与判定。
- 官网不保存账号、订单、支付信息；授权数量按账号实例计算。
- 本仓库不修改 `auto_sjz`（设备端）、`auto_sjz_manager`、`auto_sjz_deploy`。

## 分支约定

- `main`：当前线上版本
- `main_dev`：日常开发与联调分支，改完合并回 `main` 后再发布

## 管理台演示子应用（/console）

`/console` 不再手写复刻 Manager 界面，而是嵌入 Manager 前端仓库的 demo 构建。两边只共享同一套 Vue 组件，官网不复制业务页面代码。

- 子应用目录：`web/public/console-demo/`，构建时原样进入 `web/dist/console-demo/`，访问 `/console-demo/`
- 来源：`auto_sjz_manager` 的 `main_dev` 分支，`frontend` 下执行 `npm run build:demo`
- 演示环境不登录、不连真实后端，接口由 `frontend/src/demo/mock.js` 接管，只保留「业务配置 / 业务状态」
- 静态服务已支持目录 index，`/console-demo/` 不会回退到官网首页

Manager 界面更新后同步演示产物：

```bash
cd /tmp/manager_ref/frontend && git pull && npm ci && npm run build:demo && rm -rf /home/user/Doubao/chats/38454246989782786/auto_sjz_site/web/public/console-demo && cp -r dist-demo /home/user/Doubao/chats/38454246989782786/auto_sjz_site/web/public/console-demo
```

> 生产环境若把 demo 单独部署，可在构建官网时设置 `VITE_CONSOLE_DEMO_URL` 指向外部子应用地址，官网不用再带静态产物。
