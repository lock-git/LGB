package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// 初始化路由器并添加被测试的处理器方法
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthCheckHandler)

	// 新建一个请求实例模拟客户段请求，其中包含了请求方法、URL、参数等
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 新建一个 ResponseRecorder 来捕获响应
	rr := httptest.NewRecorder()

	// 传入测试请求和响应类实例并执行请求
	mux.ServeHTTP(rr, req)

	// 检查响应状态码（通过 ResponseRecorder 获取）是否是200，如果不是，则测试不通过
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 检查响应实体是否符合预期结果，如果不是，则测试不通过
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
