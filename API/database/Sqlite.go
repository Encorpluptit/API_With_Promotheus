package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// initSqlite: Initialise sqlite DBHandle
func (db *Database) initSqlite() (err error) {
	if db.DB, err = gorm.Open(sqlite.Open(db.Config["DB_NAME"]), &gorm.Config{}); err != nil {
		return err
	}
	return nil
}
