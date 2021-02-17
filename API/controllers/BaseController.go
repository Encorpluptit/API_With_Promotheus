package controllers

import "BackendGo/database"

type BaseController struct {
	Db database.Database
}

func NewBaseController() (*BaseController, error) {
	handler := &BaseController{}
	if err := handler.Db.Init(); err != nil {
		return nil, err
	}
	return handler, nil
}
