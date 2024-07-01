package gormdb

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mdcabezas/reportiny/internal/report"
)

type GormSqlite struct{}

func NewGormSqlite() report.IRepository {
	return &GormSqlite{}
}

func (r *GormSqlite) Create(db report.RepositoryModel) (report.RepositoryModel, error) {
	fmt.Println("GORM Create", db)
	dbConnect := GetGormConnection()

	db.ID = uuid.New().String()
	db.DateTime = time.Now().Format(time.RFC3339)
	fmt.Println("GORM Pre-Create", db)

	dbConnect.Table("report").Create(map[string]interface{}{
		"id":          db.ID,
		"image":       db.Image,
		"description": db.Description,
		"user":        db.User,
		"lat":         db.Latitude,
		"lng":         db.Longitude,
		"datetime":    db.DateTime,
	})

	return db, nil
}

func (r *GormSqlite) Read(id string) (report.RepositoryModel, error) {
	fmt.Printf("GORM Read %s", id)
	dbConnect := GetGormConnection()
	dbConnect.First(&report.RepositoryModel{}, "id = ?", id)

	return report.RepositoryModel{}, nil
}

func (r *GormSqlite) ReadAll() []report.RepositoryModel {
	fmt.Printf("GORM ReadAll")
	dbConnect := GetGormConnection()
	var result []report.RepositoryModel
	dbConnect.Table("report").Find(&result)

	return result
}

func (r *GormSqlite) Update(id string, db report.RepositoryModel) error {
	fmt.Printf("GORM Update %s %+v", id, db)
	dbConnect := GetGormConnection()
	db.DateTime = time.Now().Format(time.RFC3339)

	dbConnect.Table("report").Where("id = ?", id).Updates(map[string]interface{}{
		"image":       db.Image,
		"description": db.Description,
		"lat":         db.Latitude,
		"lng":         db.Longitude,
		"datetime":    db.DateTime,
	})

	return nil
}

func (r *GormSqlite) Delete(id string) error {
	fmt.Printf("GORM Delete %s", id)
	dbConnect := GetGormConnection()
	dbConnect.Table("report").Where("id = ?", id).Delete(&report.RepositoryModel{})

	return nil
}
