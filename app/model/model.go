package model

import (
	"errors"
)

var onmemory_db = make(map[string]string)

func GetById(id string) (string, error) {

	data, ok := onmemory_db[id]

	if !ok {
		return "", errors.New("not found")
	}

	return data, nil
}

type Model struct {
	ID   string `json:"id" binding:"required"`
	Data string `json:"data" binding:"required"`
}

func PostNewData(model *Model) (*Model, error) {
	if _, ok := onmemory_db[model.ID]; ok {
		return nil, errors.New("already exists")
	}

	onmemory_db[model.ID] = model.Data

	return model, nil
}
