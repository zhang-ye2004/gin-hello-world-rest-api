# 项目名
APP_NAME = my-api

# 编译输出目录
BUILD_DIR = ./build

# 默认目标：直接输入 make 时执行的任务
.PHONY: all
all: build

# 编译：生成可执行文件
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/server/main.go
	@echo "✅ 编译完成: $(BUILD_DIR)/$(APP_NAME)"

# 运行：直接 go run（开发时用）
.PHONY: run
run:
	go run cmd/server/main.go

# 清理：删除编译产物和日志文件
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -f app.log
	@echo "🧹 清理完成"

# 测试：运行所有测试
.PHONY: test
test:
	go test -v ./...

# 跨平台编译：同时生成 Linux、Windows、macOS 版本
.PHONY: build-all
build-all:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux cmd/server/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-windows.exe cmd/server/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-darwin cmd/server/main.go
	@echo "✅ 所有平台编译完成: $(BUILD_DIR)/"

