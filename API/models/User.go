package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID            uuid.UUID `gorm: unique`
	Login, Password string
	Exporters       []Exporter
}

type Exporter struct {
	gorm.Model
	UserID uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:DELETE;"`
	UUID   uuid.UUID `gorm: unique`
}

func (u *User) FindExporters(db *gorm.DB) (exporters []Exporter, err error) {
	if err = db.Debug().Model(u).Association("Exporters").Find(&exporters); err != nil {
		return nil, err
	}
	return exporters, nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Debug().Create(&u).Error
}

func (u *User) FindByLoginAndPassword(db *gorm.DB) error {
	return db.Debug().Where("login = ? AND password = ?", u.Login, u.Password).First(&u).Error
}

func (u *User) FindByID(db *gorm.DB) error {
	return db.Debug().First(&u, u.ID).Error
}

func (u *User) FindByUUID(db *gorm.DB) error {
	return db.Debug().Where("UUID = ?", u.UUID).First(&u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Debug().Save(&u).Error
}

func (u *User) UpdateFields(db *gorm.DB, fields map[string]interface{}) error {
	return db.Debug().Model(&u).Updates(fields).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Debug().Delete(&u).Error
}

func (u *User) DatasEqual(lhs *User) bool {
	return u.Login == lhs.Login && u.Password == lhs.Password
}

func (u *User) Equal(lhs *User) bool {
	return u.ID == lhs.ID && u.DatasEqual(lhs)
}
