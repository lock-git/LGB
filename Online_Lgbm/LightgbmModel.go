package main

import (
	"github.com/dmitryikh/leaves"
)

type FeatureData struct {
	Values []float64 `json:"values"`
	Rows   int       `json:"rows"`
	Cols   int       `json:"cols"`
}

var model *leaves.Ensemble

func ForecastLgb(data FeatureData) []float64 {

	predictions := make([]float64, data.Rows*model.NOutputGroups())
	// specify num of threads and do predictions
	model.PredictDense(data.Values, data.Rows, data.Cols, predictions, 0, 8)
	return predictions
}

func InitModel(fileName string) error {
	// lgb_ranker.model 放入项目下
	//model2, err := leaves.LGEnsembleFromFile("lightgbm_model_20210224.txt", true)
	model2, err := leaves.LGEnsembleFromFile(fileName, true)
	if err != nil {
		return err
	}
	model = model2
	return err
}
