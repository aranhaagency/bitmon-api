package controllers

import (
	"errors"
	"github.com/bitmon-world/bitmon-api/models"
)

type TestRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BitmonController struct {
	dbModel *models.BitmonDBModel
}

func (ctrl *BitmonController) GetGeneralMon(id string) (interface{}, error) {
	data, err := ctrl.dbModel.GetGenMon(id)
	if err != nil {
		return nil, errors.New("general monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) GetParticularMon(id string) (interface{}, error) {
	data, err := ctrl.dbModel.GetPartMon(id)
	if err != nil {
		return nil, errors.New("particular monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) GetUserMons(id string) (interface{}, error) {
	return nil, nil
}

func NewBitmonController(dbUri string, dbName string) *BitmonController {
	model := models.NewDBModel(dbUri, dbName)
	return &BitmonController{dbModel: model}
}
