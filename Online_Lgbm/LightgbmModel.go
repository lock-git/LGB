package main

import (
	"fmt"
	"github.com/dmitryikh/leaves"
)

type FeatureData struct {
	Values []float64 `json:"values"`
	Rows   int       `json:"rows"`
	Cols   int       `json:"cols"`
}

var model *leaves.Ensemble

func forecast(data FeatureData) []float64 {

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
