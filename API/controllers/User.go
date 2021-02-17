package controllers

import (
	"BackendGo/models"
	"github.com/google/uuid"
	"log"
)

// CreateUser: Create an user in Database given a login, password and a UUID
func (h *BaseController) CreateUser(login, password, UUID string) (*models.User, error) {
	userUuid, err := uuid.Parse(UUID)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Login:    login,
		Password: password,
		UUID:     userUuid,
	}
	if err = user.Create(h.Db.DB); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (h *BaseController) GetUser(login, password string) (*models.User, error) {
	user := &models.User{
		Login:    login,
		Password: password,
	}
	if err := user.FindByLoginAndPassword(h.Db.DB); err != nil {
		return nil, err
	}
	return user, nil
}

func (h *BaseController) getExportersFromUserID(user *models.User) ([]models.Exporter, error) {
	exporters, err := user.FindExporters(h.Db.DB)
	if err != nil {
		return nil, err
	}
	return exporters, nil
}

func (h *BaseController) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	user.ID = id
	if err := user.FindByID(h.Db.DB); err != nil {
		return nil, err
	}
	return user, nil
}

func (h *BaseController) GetUserByUUIDID(UUID string) (*models.User, error) {
	userUuid, err := uuid.Parse(UUID)
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	user.UUID = userUuid
	if err = user.FindByID(h.Db.DB); err != nil {
		return nil, err
	}
	return user, nil
}

func (h *BaseController) UpdateUser(id uint, fields map[string]interface{}) (*models.User, error) {
	user := &models.User{}
	user.ID = id
	if err := user.FindByID(h.Db.DB); err != nil {
		return nil, err
	}
	if err := user.UpdateFields(h.Db.DB, fields); err != nil {
		return nil, err
	}
	return user, nil
}

func (h *BaseController) DeleteUser(id uint) (*models.User, error) {
	user := &models.User{}
	user.ID = id
	if err := user.Delete(h.Db.DB); err != nil {
		return nil, err
	}
	return user, nil
}
