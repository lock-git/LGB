package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type result struct {
	Stamp int         `json:"stamp"`
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func ParseGZipString(str string) (string, error) {

	dataByte, byteErr := base64.StdEncoding.DecodeString(str)
	if byteErr != nil {
		return "", byteErr
	}
	r, gErr := gzip.NewReader(bytes.NewReader(dataByte))
	if gErr != nil {

		return "", gErr
	}
	s, rErr := ioutil.ReadAll(r)
	if rErr != nil {
		return "", rErr
	}
	return string(s), nil
}

func LgbPredict(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	var beginTime = time.Now().UnixNano() / 1e6
	fmt.Println(time.Now(), "====================== initModel_start...")

	if r.Form["feature"] == nil || len(r.Form["feature"]) == 0 || r.Form["feature"][0] == "" {
		result := result{Stamp: 0, Code: -1, Msg: "失败", Data: "ERROR:feature参数不能为nil"}
		nilJsonErr, _ := json.Marshal(result)
		_, _ = w.Write(nilJsonErr)
		return
	}

	if r.Form["modelName"] == nil || len(r.Form["modelName"]) == 0 || r.Form["modelName"][0] == "" {
		result := result{Stamp: 0, Code: -1, Msg: "失败", Data: "ERROR:modelName参数不能为nil"}
		nilJsonErr, _ := json.Marshal(result)
		_, _ = w.Write(nilJsonErr)
		return
	} else {
		fmt.Println(time.Now(), "====================== model_name:", r.Form["modelName"][0])
	}

	initModelErr := InitModel(r.Form["modelName"][0])
	if initModelErr != nil {
		data := "ERROR：初始化模型失败 " + initModelErr.Error()
		result := result{Stamp: 0, Code: -1, Msg: "失败", Data: data}
		errJson, _ := json.Marshal(result)
		_, _ = w.Write(errJson)
		return
	}

	var initModelEndTime = int32(time.Now().UnixNano() / 1e6)
	fmt.Println(time.Now(), "====================== initModel_duration:", initModelEndTime-int32(beginTime), "ms")

	parseFeature, parseGZipErr := ParseGZipString(r.Form["feature"][0])
	if parseGZipErr != nil {
		data := "ERROR：feature参数GZip解压失败 " + parseGZipErr.Error()
		result := result{Stamp: 0, Code: -1, Msg: "失败", Data: data}
		errJson, _ := json.Marshal(result)
		_, _ = w.Write(errJson)
		return
	}

	var f FeatureDataWithEssay
	var json2 = jsonIter.ConfigCompatibleWithStandardLibrary
	err := json2.UnmarshalFromString(parseFeature, &f)
	if err != nil {
		data := "ERROR：json格式有误 " + err.Error()
		result := result{Stamp: 0, Code: -1, Msg: "失败", Data: data}
		errJson, _ := json.Marshal(result)
		_, _ = w.Write(errJson)
		return
	}

	predictData, predictErr := ForecastLgbV2(f)
	if predictErr != nil {
		data := "ERROR：model prediction fail " + predictErr.Error()
		result := result{Stamp: 0, Code: -1, Msg: "失败", Data: data}
		errJson, _ := json.Marshal(result)
		_, _ = w.Write(errJson)
		return
	} else {
		result := result{Stamp: 0, Code: 0, Msg: "成功", Data: predictData}
		resultJsonStr, _ := json.Marshal(result)
		_, _ = w.Write(resultJsonStr)
	}

	var predictEndTime = int32(time.Now().UnixNano() / 1e6)
	fmt.Println(time.Now(), "====================== predict_duration:", predictEndTime-initModelEndTime, "ms", "============================================= get serve success")
}

/*
关闭http
*/
func SayBye(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	_, _ = w.Write([]byte("bye bye ,shutdown the server"))
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
			log.Fatal("Server closed under request \n", time.Now())
		} else {
			log.Fatal("Server closed unexpected \n", time.Now(), err)
		}
	}
	log.Fatal("Server exited")

}
