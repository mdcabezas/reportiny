package sqlitedb

import (
	"database/sql"
)

// Puntero a la estructura DB.
var db *sql.DB

func GetConnection() *sql.DB {
	// Evita una nueva conexión en cada llamada.
	if db != nil {
		return db
	}

	var err error
	// Conexión a la base de datos
	db, err = sql.Open("sqlite3", "./internal/report/report.sqlite3")
	if err != nil {
		panic(err)
	}
	return db
}
