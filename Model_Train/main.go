package main

import (
	"fmt"
	"github.com/dmitryikh/leaves"
	jsoniter "github.com/json-iterator/go"
	//"io/ioutil"
	//"os"
)

type IT struct {
	Values []float64 `json:"values"`
	Rows   int       `json:"rows"`
	Cols   int       `json:"cols"`
}

var model *leaves.Ensemble

func forecast(data IT) []float64 {

	predictions := make([]float64, data.Rows*model.NOutputGroups())
	// specify num of threads and do predictions
	model.PredictDense(data.Values, data.Rows, data.Cols, predictions, 0, 4)
	return predictions
}

func initModel() {
	// lgb_ranker.model 放入项目下
	model2, err := leaves.LGEnsembleFromFile("lgb_ranker.model", true)
	if err != nil {
		print(err)
		fmt.Println("err init model = ", err)
	}
	model = model2
}

func main1() {
	//json data template
	jsonBuf := `
    {
    "values": [
    60,3241234,29,266122,72754,11394,0.16,0.869999999999999,58840,20,0,44,395511,34.7122169562927,0.878262153482706,170736,27091,947298,34.9672584991325,0,33,2,0,1,2,1,30,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0,
    60,3241234,29,266122,72754,11394,0.16,0.869999999999999,58840,20,0,44,395511,34.7122169562927,0.878262153482706,170736,27091,947298,34.9672584991325,0,33,2,0,1,2,1,30,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0
    ],
    "rows": 2,
    "cols": 147
    }`
	//f,errs := os.Open("a.json")
	//if errs != nil {
	//	fmt.Println("errs = ", errs)
	//	return
	//}
	//jsonBuf,_ := ioutil.ReadAll(f)

	var tmp IT
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp feature ============ : %s\n", tmp)

	// loading model
	initModel()
	fmt.Printf("Name ============ : %s\n", model.Name())
	fmt.Printf("NFeatures: %d\n", model.NFeatures())
	fmt.Printf("NOutputGroups: %d\n", model.NOutputGroups())
	fmt.Printf("NEstimators: %d\n", model.NEstimators())
	fmt.Printf("Transformation: %s\n", model.Transformation().Name())

	// preallocate slice to store model predictions
	predictions := make([]float64, tmp.Rows*model.NOutputGroups())
	// specify num of threads and do predictions
	model.PredictDense(tmp.Values, tmp.Rows, tmp.Cols, predictions, 0, 1)
	fmt.Println(predictions)

}
