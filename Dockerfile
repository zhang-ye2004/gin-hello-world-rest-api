# ===== 构建阶段 =====
# 使用官方的 Go 镜像来编译程序
FROM golang:1.25-alpine AS builder

# 设置 Go 代理（国内加速）
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum，先下载依赖（利用 Docker 缓存）
COPY go.mod go.sum ./
RUN go mod download

# 复制所有源代码
COPY . .

# 编译成可执行文件
RUN go build -o my-api cmd/server/main.go


# ===== 运行阶段 =====
# 使用极简的 Alpine Linux 镜像来运行程序（只有几 MB）
FROM alpine:latest

# 安装 ca-certificates，用于 HTTPS 请求
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译好的可执行文件
COPY --from=builder /app/my-api .

# 暴露端口
EXPOSE 8080

# 运行程序
CMD ["./my-api"]

