# Gin Hello World REST API

一个基于 Go Gin 框架的 REST API 示例项目，用于学习和实践完整的 CI/CD 流程。

**项目地址**：[https://github.com/zhang-ye2004/gin-hello-world-rest-api](https://github.com/zhang-ye2004/gin-hello-world-rest-api)

---

## 📌 功能

- 提供 `/` 根路径，返回欢迎信息
- 提供 `/env` 路径，显示当前运行环境变量
- 提供 `/error` 路径，模拟错误日志记录
- 支持通过环境变量配置（端口、运行模式等）
- 完整的日志记录（同时输出到文件和控制台）

---

## 🚀 本地运行

### 1. 克隆项目
```bash
git clone https://github.com/zhang-ye2004/gin-hello-world-rest-api.git
cd gin-hello-world-rest-api
2. 下载依赖
bash
go mod download
3. 运行
bash
go run cmd/server/main.go
服务默认运行在 :8080，访问 http://localhost:8080 即可看到响应。

4. 通过环境变量配置
bash
PORT=9090 GIN_MODE=release go run cmd/server/main.go
环境变量	默认值	说明
PORT	8080	服务监听端口
GIN_MODE	debug	Gin 运行模式（debug/release/test）
LOG_FILE	app.log	日志文件路径
🐳 使用 Docker 运行
直接拉取镜像运行
bash
docker run -p 8080:8080 zhangye2004/my-api:latest
或自行构建镜像
bash
docker build -t my-api .
docker run -p 8080:8080 my-api
🛠️ 开发工具
bash
# 编译
make build          # 生成 ./build/my-api

# 运行（开发模式）
make run

# 运行测试
make test

# 清理编译产物和日志
make clean

# 跨平台编译（Linux/Windows/macOS）
make build-all
🔧 项目结构
text
.
├── cmd/server/           # 程序入口
├── internal/
│   ├── config/           # 配置加载
│   ├── handler/          # HTTP 处理器
│   ├── logger/           # 日志初始化
│   ├── middleware/       # 中间件（日志等）
│   └── model/            # 数据模型
├── pkg/utils/            # 可复用的工具函数
├── .github/workflows/    # CI/CD 工作流
├── Dockerfile            # Docker 镜像构建文件
├── Makefile              # 自动化构建脚本
├── go.mod
├── go.sum
└── README.md
⚙️ CI/CD 流程
本项目通过 GitHub Actions 实现了自动化 CI/CD。

工作流	触发条件	执行内容
ci.yml	推送代码到 main 分支	拉取代码 → 安装 Go → 下载依赖 → 运行测试 → 编译项目
docker-publish.yml	推送代码到 main 分支	构建 Docker 镜像 → 推送到 Docker Hub
镜像地址：zhangye2004/my-api:latest

每次推送代码到 main 分支，都会自动触发上述流程，确保代码始终处于可部署状态。

📝 环境变量与配置
变量	默认值	说明
PORT	8080	服务端口
GIN_MODE	debug	Gin 运行模式
LOG_FILE	app.log	日志文件路径
DB_HOST	localhost	数据库主机（示例）
DB_PORT	5432	数据库端口（示例）
