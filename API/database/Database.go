package database

import (
	"BackendGo/models"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"os"
)

const (
	DbDriverEnvVar = "DB_DRIVER"
	DriverPostgres = "postgres"
	DriverSqlite   = "sqlite"
)

type Database struct {
	*gorm.DB
	Config map[string]string
}

const (
	User1UUID = "4d365568-d3ac-4880-8403-8d4e2638e008"
	User2UUID = "1d149631-4141-4be4-9244-cfd78afbfc57"
)

var (
	User1Exporters = []string{
		"c9bb3e99-df72-4603-b188-b6f856f0ef9f",
		"8d4c8325-a1c6-4bb1-b50d-dc7f7d92a5b5",
	}
	User2Exporters = []string{
		"ba720cfd-33dc-45c8-9879-7038a965ca18",
	}
)

// driver: Type to quickly change database with correct environment variables and init database function.
type driver struct {
	EnvVars map[string]string
	InitFct func(d *Database) error
}

// DriverMap: for each DB_DRIVER, contains necessary environment variables to load and associated database initialisation function
// database.initPostGreSql
var DriverMap = map[string]driver{
	DriverPostgres: {
		EnvVars: map[string]string{
			"DB_NAME": "", "DB_HOST": "", "DB_PORT": "", "DB_USER": "", "DB_PASSWORD": "",
		},
		InitFct: func(d *Database) error { return d.initPostGreSql() },
	},
	DriverSqlite: {
		EnvVars: map[string]string{
			"DB_NAME": "",
		},
		InitFct: func(d *Database) error { return d.initSqlite() },
	},
}

// Initialise GormDatabase
func (db *Database) Init() error {
	dbInitFct, err := db.getConfig()
	if err != nil {
		return err
	}
	if err = dbInitFct(db); err != nil {
		return fmt.Errorf("failed to initialize GormDatabase: %v", err)
	}
	if err = db.autoMigrate(); err != nil {
		return err
	}
	return db.DefaultsUsers()
}

// Create Defaults User with mock up Data
// TODO: Remove when User Register up and running
func (db Database) DefaultsUsers() (err error) {
	if err = db.FindOrCreateDefaultUsers("user_1", "pass_1", User1UUID, User1Exporters); err != nil {
		return err
	}
	if err = db.FindOrCreateDefaultUsers("user_2", "pass_2", User2UUID, User2Exporters); err != nil {
		return err
	}
	return nil
}

// Helper function to create default users
func (db Database) FindOrCreateDefaultUsers(login, pass, UUID string, exporters []string) (err error) {
	Uuid := uuid.UUID{}
	if Uuid, err = uuid.Parse(UUID); err != nil {
		return err
	}
	user := models.User{
		UUID:     Uuid,
		Login:    login,
		Password: pass,
	}
	for _, exporter := range exporters {
		if Uuid, err = uuid.Parse(exporter); err != nil {
			return err
		}
		user.Exporters = append(user.Exporters, models.Exporter{
			UUID: Uuid,
		})
	}
	if err = user.FindByUUID(db.DB); err != nil {
		if err = user.Create(db.DB); err != nil {
			return err
		}
	}
	return nil
}

// autoMigrate: Auto migrate models
func (db *Database) autoMigrate() (err error) {
	// GormDatabase migration
	if err = db.DB.Debug().AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err = db.DB.Debug().AutoMigrate(&models.Exporter{}); err != nil {
		return err
	}
	return nil
}

// getConfig: Load env from .env file (development) or deployment env
func (db *Database) getConfig() (func(*Database) error, error) {
	driverEnvVar := os.Getenv(DbDriverEnvVar)
	if driverEnvVar == "" {
		return nil, fmt.Errorf("environment variable missing : %s", DbDriverEnvVar)
	}

	driverHandler, ok := DriverMap[driverEnvVar]
	if !ok {
		return nil, fmt.Errorf("no driver found for : %s", DbDriverEnvVar)
	}

	db.Config = driverHandler.EnvVars
	var value string
	for envVar := range db.Config {
		if value = os.Getenv(envVar); value == "" {
			return nil, fmt.Errorf("environment variable missing : %s", envVar)
		}
		db.Config[envVar] = value
	}
	return driverHandler.InitFct, nil
}
