package main

import (
	"encoding/json"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type result struct {
	Status bool
	Code   int
	Data   interface{}
}

/*
模型预测
*/
func LgbPredict(w http.ResponseWriter, r *http.Request) {

	var beginTime = time.Now().UnixNano() / 1e6
	fmt.Println("initModel_start == ", beginTime)

	r.ParseForm()
	if r.Form["feature"] == nil || len(r.Form["feature"]) == 0 || r.Form["feature"][0] == "" {
		result := result{Status: true, Code: -1, Data: "ERROR:feature参数不能为nil"}
		nilJsonErr, _ := json.Marshal(result)
		w.Write(nilJsonErr)
		return
	}

	if r.Form["modelName"] == nil || len(r.Form["modelName"]) == 0 || r.Form["modelName"][0] == "" {
		result := result{Status: true, Code: -1, Data: "ERROR:modelName参数不能为nil"}
		nilJsonErr, _ := json.Marshal(result)
		w.Write(nilJsonErr)
		return
	}

	initModel(r.Form["modelName"][0])

	var initModelEndTime = int32(time.Now().UnixNano() / 1e6)
	fmt.Println("initModel_duration == ", initModelEndTime-int32(beginTime))

	var f FeatureData
	var json2 = jsonIter.ConfigCompatibleWithStandardLibrary
	err := json2.UnmarshalFromString(r.Form["feature"][0], &f)
	if err != nil {
		data := "ERROR：json格式有误" + err.Error()
		result := result{Status: true, Code: -1, Data: data}
		errJson, _ := json.Marshal(result)
		w.Write(errJson)
		return
	}
	//fmt.Println("feature_into == ", f)
	fmt.Println("feature_into_time == ", time.Now())
	predictData := forecast(f)
	result := result{Status: true, Code: 0, Data: predictData}
	resultJsonStr, _ := json.Marshal(result)
	w.Write(resultJsonStr)

	var predictEndTime = int32(time.Now().UnixNano() / 1e6)
	fmt.Println("predict_duration == ", predictEndTime-initModelEndTime)
}

/*
关闭http
*/
func SayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the server"))
	err := server.Shutdown(nil)
	if err != nil {
		log.Fatal([]byte("shutdown the server err"))
	}
}

// 主动关闭服务器
var server *http.Server

func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/lgbmPredict", LgbPredict)
	mux.HandleFunc("/bye", SayBye)

	server = &http.Server{
		Addr:         ":8034",
		WriteTimeout: time.Second * 2,
		Handler:      mux,
	}

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting  HttpServer")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request", time.Now())
		} else {
			log.Fatal("Server closed unexpected", time.Now(), err)
		}
	}
	log.Fatal("Server exited")

}
