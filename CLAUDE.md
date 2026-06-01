# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

GoRedisAdmin 是一个使用 Go (Gin) + Vue 2 (Element UI) 开发的 Redis 后台管理平台，提供在线数据库管理和操作界面。默认端口 9527。

## 常用命令

### Go 后端

```bash
# 安装依赖
go mod tidy

# 运行服务（需要先创建 var 目录）
mkdir -p var
go run main.go

# 构建可执行文件
go build -o gradmin main.go

# 运行测试
go test ./utils/...

# 使用 shell 脚本构建并启动
cd shell && sh run.sh gradmin build
# 其他操作: start / stop / restart / build
```

### Vue 前端

```bash
cd web
npm install
npm run serve    # 开发服务器
npm run build    # 构建到 html/ 目录
npm run lint     # ESLint 检查
```

前端构建产物直接输出到 `html/` 目录，Go 后端会静态托管这些文件。

## 架构

### 分层结构

```
main.go                  → 入口：init() 初始化配置，main() 启动日志和路由
config.ini               → 全部配置（Redis连接、IP白名单、管理员账号、端口、日志）

routers/
  index.go               → Gin 引擎初始化，挂载静态资源、视图路由、API路由
  view_router/            → 前端静态页面路由（/ → html/index.html）
  api_router/
    user_router/          → /api/v1/user/login, /logout（无需JWT）
    db_router/            → /api/v1/db/*（需JWT认证）
    info_router/          → /api/v1/info/get_info（需JWT认证）
  middleware/
    index.go              → IP白名单中间件
    authorize_jwt.go      → JWT认证中间件（Bearer Token，300秒自动续期）

controller/
  base_controller.go      → 基础控制器，封装响应、参数转换、JSON解析
  user_controller/        → 登录/注销（凭据来自 config.ini）
  db_data_controller/     → Redis数据CRUD核心逻辑
    db_data_cont.go       → 主控制器：DbList, GetKeys, GetVal, DelKey, AddVal, Flush, ExportKey, ExpireKey
    db_cont_help.go       → 辅助控制器+模型：按类型(string/list/set/zset/hash)操作Redis
  info_controller/        → Redis INFO 命令代理

global/
  global_redis/           → Redis连接工厂（go-redis v6），支持 db 0-15
  global_write_ip/        → IP白名单内存存储
  global_response/        → 统一JSON响应结构 {code, message, data}
  initData/               → 配置加载（ini.v1），初始化Redis配置和IP白名单

utils/
  bcrypt_utils.go         → JWT Token 生成/验证（HS256，密钥硬编码）、bcrypt密码工具
  log_utils/              → logrus 异步日志（channel + goroutine），按天切割
  exoprt_utils/           → Redis key批量导出为JSON文件
  handle_error_utils.go   → 全局错误处理工具
```

### 关键设计决策

- **单 Redis 实例**：配置只支持一个 Redis 服务器（config.ini `[redis]`），通过 `global_redis.GetRedisClient(db)` 选择 db 0-15。每次操作创建新连接，用完即关。
- **无数据库层**：管理员账号直接从 config.ini 读取，无持久化用户表。
- **认证流程**：用户登录 → 前端存储 JWT Token → 后续请求带 `Authorization: Bearer <token>` → 中间件校验并在即将过期时通过 `new-token` header 自动续期。
- **前端是预编译产物**：`html/` 目录是 Vue 项目构建后的静态文件，Go 直接托管。前端源码在 `web/` 目录。
- **响应码约定**：`0` = 成功，`7` = 错误，`6` = 未登录。

### 前端结构 (web/)

Vue 2 + Element UI + Vue Router + Axios。API 调用集中在 `web/src/api/`，视图在 `web/src/views/`（login、homeItem、header、right、infoItem）。

## 配置说明

`config.ini` 是唯一配置文件，包含：
- `[redis]` — Redis 连接参数（host, port, pwd, timeout）
- `[whitelist_ip]` — IP 白名单，逗号分隔，为空则不限制
- `[admin]` — 登录凭据和后台端口
- `[log]` — 日志路径和保留天数

## 当前改动摘要（2026-06-01）

基于 `git status` 与 `git diff --stat`，当前工作区核心改动如下：

- **后端 Redis 取值修复与补全**
  - `controller/db_data_controller/db_cont_help.go`
    - 修复 `GetString()` 读取参数（由 `d.Val` 改为 `d.Key`）。
    - 新增 `GetList()` / `GetSet()` / `GetZSet()` / `GetHash()` / `GetStream()`，统一返回 `{data, count}` 结构，便于前端渲染。
  - `controller/db_data_controller/db_data_cont.go`
    - `handleGetVal` 增加 `list`/`set`/`zset`/`hash`/`stream` 分支处理，替换原先未实现逻辑。

- **前端数据查看体验增强（homeItem）**
  - `web/src/views/homeItem/data.vue`
    - 重构 key 详情弹窗：增加元信息面板（类型、TTL、大小）、按类型分表格展示、复制单项/全部、字符串 JSON 美化、stream 字段解析。
  - `web/src/views/homeItem/form_page.vue`
    - 增加按数据类型动态 placeholder 与格式提示（list/set/zset/stream）。
  - `web/src/views/homeItem/index.vue`
    - 查看数据弹窗标题带 key 名称，弹窗宽度调整为 720px。

- **前端通用交互与文案补充**
  - `web/src/api/request.js`：登录失效（code=6）时主动清理 token 并跳转登录页。
  - `web/src/App.vue`：修复下拉层级（z-index）与对话框内选择器宽度。
  - `web/src/i18n/messages/en-US.js`、`web/src/i18n/messages/zh-CN.js`：新增 data/form 相关文案键，支持新 UI。

- **构建产物与项目文件变更**
  - `html/` 下静态资源哈希文件更新（前端重新构建导致）。
  - `.idea/` 目录下若干文件删除。
  - 新增 `CLAUDE.md`（仓库协作说明文档）。

## 最近 Git 提交摘要（最近 5 次）

- `8bc54fa` feat(redis): support stream key type in web admin
  - 支持 stream 类型在管理台查看/处理，联动后端控制器与前端页面。
- `eb3c664` docs: refresh bilingual screenshots for login and panels
  - 更新 README 双语截图资源。
- `c636e96` feat(i18n): add Chinese/English UI and language switch
  - 新增中英文界面切换、文案与前端 i18n 基础设施。
- `caa9168` docs(go): add detailed English comments across backend code
  - 后端主要模块补充英文注释，提升可维护性。
- `7e4556f` chore: upgrade deps and fix failing go tests
  - 升级依赖并修复部分测试/构建问题，更新前端构建产物。
