package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-hello-world-rest-api/internal/config"

	"github.com/gin-gonic/gin"
)

func TestRootHandler(t *testing.T) {
	// 1. 设置 Gin 为测试模式（减少输出）
	gin.SetMode(gin.TestMode)

	// 2. 加载配置
	cfg := config.Load()

	// 3. 创建路由并注册测试用的 handler
	r := gin.Default()
	r.GET("/", RootHandler(cfg))

	// 4. 创建一个模拟的 HTTP 请求
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// 5. 执行请求
	r.ServeHTTP(w, req)

	// 6. 检查状态码
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 %d，实际 %d", http.StatusOK, w.Code)
	}

	// 7. 检查返回的 JSON 内容（可选）
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("解析JSON失败: %v", err)
	}

	if msg, ok := resp["message"]; !ok || msg != "Hello from Gin!" {
		t.Errorf("期望message为：'Hello from Gin!',实际为：%v", msg)
	}
}

func TestEnvHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.Load()

	r := gin.Default()
	r.GET("/env", EnvHandler(cfg))

	req := httptest.NewRequest("GET", "/env", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，实际 %d", w.Code)
	}

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("解析 JSON 失败: %v", err)
	}

	//期望的返回结果不同
	if _, ok := resp["PORT"]; !ok {
		t.Error("返回的 JSON 中缺少 PORT 字段")
	}
}
