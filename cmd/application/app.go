package application

import (
	"log"

	"github.com/mdcabezas/reportiny/internal/report"
	"github.com/mdcabezas/reportiny/pkg/gormdb"
	"github.com/mdcabezas/reportiny/pkg/sqlitedb"
)

// Acá van todos los servicios que se van a utilizar en la aplicación
type Components struct {
	Report report.IReport
}

func NewComponents() *Components {
	return &Components{
		// Report: report.NewReport(getDBReport("gorm")),
		Report: report.NewReport(getDBReport("sqlite")),
	}
}

func getDBReport(s string) report.IRepository {

	switch s {
	case "sqlite":
		return sqlitedb.NewDBSqlite()

	case "memory":
		return report.NewRepository()

	case "gorm":
		return gormdb.NewGormSqlite()
	default:
		log.Panic("No se ha encontrado el tipo de base de datos:")
		return nil
	}
}
