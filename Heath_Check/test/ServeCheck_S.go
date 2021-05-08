package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// 一个非常简单的健康检查实现：如果此 HTTP 接口调用成功，则表示应用健康
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// 后续我们还可以通过执行 PING 指令反馈 DB、缓存状态，并将它们的健康检查结果放到响应中
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
