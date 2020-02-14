package controllers

import (
	"errors"
	"github.com/bitmon-world/bitmon-api/models"
	"github.com/bitmon-world/bitmon-api/types"
)

type TestRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BitmonController struct {
	dbModel *models.BitmonDBModel
}

func (ctrl *BitmonController) GetMonList(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetGenMon(params.ID)
	if err != nil {
		return nil, errors.New("general monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) GetMonInfo(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetGenMon(params.ID)
	if err != nil {
		return nil, errors.New("general monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) AddMon(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetGenMon(params.ID)
	if err != nil {
		return nil, errors.New("general monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) CalcAdventure(params types.ReqParams) (interface{}, error) {
	return nil, nil
}

func NewBitmonController(dbUri string, dbName string) *BitmonController {
	model := models.NewDBModel(dbUri, dbName)
	return &BitmonController{dbModel: model}
}
