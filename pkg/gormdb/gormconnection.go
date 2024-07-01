package gormdb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Puntero a la estructura DB.
var dbGorm *gorm.DB

func GetGormConnection() *gorm.DB {

	if dbGorm != nil {
		return dbGorm
	}
	var err error
	// Conexión a la base de datos
	dbGorm, err = gorm.Open(sqlite.Open("./internal/report/report.sqlite3"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return dbGorm
}
