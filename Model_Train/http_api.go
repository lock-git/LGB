package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type indexHandler struct{}

type result struct {
	Status bool
	Code   int
	Data   interface{}
}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var tmp IT
	var json2 = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json2.Unmarshal(body, &tmp)
	if err != nil {
		fmt.Println("err_ServeHTTP = ", err)
		return
	}
	ret := forecast(tmp)
	jsonStr, _ := json.Marshal(ret)
	w.Write(jsonStr)
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello lock!")
}

func sayHelloName1(w http.ResponseWriter, r *http.Request) {

	// 查看完整传递参数
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(r).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	fmt.Println("fields 完整参数的map结构 ====", fields)

	r.ParseForm()                                      // 解析参数
	fmt.Println(r.Form)                                // 将所有传入的参数以map集合的方式打印输出
	fmt.Println("path == ", r.URL.Path)                // 路由
	fmt.Println("feature1 == ", r.Form["feature1"][0]) // 第一个参数
	fmt.Println("feature2 == ", r.Form["feature2"][0]) // 第二个参数
	for k, v := range r.Form {                         // 遍历参数
		fmt.Println("k == ", k)
		fmt.Println("v == ", strings.Join(v, ""))
		fmt.Println("============================================================================")
	}
	fmt.Fprintf(w, "Hello lock! \n\n\n ")

	var tmp IT
	var json2 = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json2.UnmarshalFromString(r.Form["feature2"][0], &tmp)
	if err != nil {
		fmt.Println("err_sayHelloName1 = ", err)
		return
	}
	ret := forecast(tmp)
	if err != nil {
		data := fmt.Sprintf("ERROR：数据格式有误! values:%T,cols:%T,rows:%T,err:%s", tmp.Values, tmp.Cols, tmp.Rows, err.Error())
		result := result{Status: true, Code: -1, Data: data}
		nilJsonErr, _ := json.Marshal(result)
		w.Write(nilJsonErr)
		return
	}
	jsonStr, _ := json.Marshal(ret)
	w.Write(jsonStr)                                        // 预测结果
	fmt.Println("result_data ====== ", byteString(jsonStr)) // 打印输出结果
}

/**
[]byte 转化为 string
*/
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func main() {
	initModel()
	//http.Handle("/model_predict", &indexHandler{})
	http.HandleFunc("/sayHelloName1", sayHelloName1) // 设置访问的路由
	err := http.ListenAndServe(":8034", nil)         // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServer Failed:", err)
	}

}
