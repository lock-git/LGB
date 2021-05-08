package Heath_Check

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// 方式一： curl
	/*	url := "http://localhost:8034/lgbmPredict"
		fmt.Println(url)

		curl := exec.Command("curl", "-v", url)
		out, err := curl.Output()
		if err != nil {
			fmt.Println("error" , err)
			return
		}
		fmt.Println(string(out))*/

	// 方式二：http
	response, err := http.Get("http://localhost:8034/lgbmPredict")
	if err != nil {
		fmt.Println("error ==", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println("get back success :", string(body))

}
