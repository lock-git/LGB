package main

import (
	"github.com/dmitryikh/leaves"
)

type FeatureDataWithEssay struct {
	Values   []float64 `json:"values"`
	EssayIds []string  `json:"essayIds"`
	Rows     int       `json:"rows"`
	Cols     int       `json:"cols"`
}

type FeatureData struct {
	Values []float64 `json:"values"`
	Rows   int       `json:"rows"`
	Cols   int       `json:"cols"`
}

type EssayInfo struct {
	EssayId   string  `json:"essayId"`
	SortScore float64 `json:"sortScore"`
}

var model *leaves.Ensemble

func ForecastLgb(data FeatureData) []float64 {

	predictions := make([]float64, data.Rows*model.NOutputGroups())
	// specify num of threads and do predictions
	model.PredictDense(data.Values, data.Rows, data.Cols, predictions, 0, 8)
	return predictions
}

func ForecastLgbV2(data FeatureDataWithEssay) []EssayInfo {

	EssayInfoArr := make([]EssayInfo, data.Rows*model.NOutputGroups())
	predictions := make([]float64, data.Rows*model.NOutputGroups())

	_ = model.PredictDense(data.Values, data.Rows, data.Cols, predictions, 0, 8)

	for k, v := range data.EssayIds {
		EssayInfoArr[k] = EssayInfo{EssayId: v, SortScore: predictions[k]}
	}

	return EssayInfoArr
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
