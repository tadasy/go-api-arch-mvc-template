package models

import (
	"errors"
	"fmt"
	"go-api-arch-mvc-template/configs"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	InstanceSqlite int = iota
	InstanceMysql
)

var (
	DB                            *gorm.DB
	errInvalidSQLDatabaseInstance = errors.New("invalid SQL database instance")
)

func GetModels() []interface{} {
	return []interface{}{&Album{}, &Category{}}
}

func NewDatabaseSQLFactory(instance int) (db *gorm.DB, err error) {
	switch instance {
	case InstanceSqlite:
		db, err = gorm.Open(sqlite.Open(configs.Config.DBName), &gorm.Config{})
	case InstanceMysql:
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
			configs.Config.DBUser,
			configs.Config.DBPassword,
			configs.Config.DBHost,
			configs.Config.DBPort,
			configs.Config.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
	return db, err
}

func SetDatabase(instance int) (err error) {
	db, err := NewDatabaseSQLFactory(instance)
	if err != nil {
		return err
	}
	DB = db
	return err
}
