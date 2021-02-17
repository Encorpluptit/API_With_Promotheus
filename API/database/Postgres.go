package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initPostGreSql: Initialise Postgresql DBHandle
func (db *Database) initPostGreSql() (err error) {
	url := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		db.Config["DB_HOST"], db.Config["DB_PORT"], db.Config["DB_USER"], db.Config["DB_NAME"], db.Config["DB_PASSWORD"],
	)
	db.DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
