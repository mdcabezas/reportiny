package sqlitedb

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdcabezas/reportiny/internal/report"
)

type DBSqlite struct{}

func NewDBSqlite() report.IRepository {
	return &DBSqlite{}
}

func (r *DBSqlite) Create(db report.RepositoryModel) (report.RepositoryModel, error) {
	fmt.Println("Create", db)
	dbConnect := GetConnection()

	pq := `INSERT INTO report (id, image, description, user, lat, lng, datetime) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := dbConnect.Prepare(pq)
	fmt.Println("err", err)

	if err != nil {
		return report.RepositoryModel{}, err
	}

	defer stmt.Close()
	db.ID = uuid.New().String()
	db.DateTime = time.Now().Format(time.RFC3339)

	sqlResult, sqlErr := stmt.Exec(db.ID, db.Image, db.Description, db.User, db.Latitude, db.Longitude, db.DateTime)

	fmt.Println("sqlResult", sqlResult)
	fmt.Println("sqlErr", sqlErr)
	return db, nil
}

func (r *DBSqlite) Read(id string) (report.RepositoryModel, error) {

	return report.RepositoryModel{}, nil
}

func (db *DBSqlite) ReadAll() []report.RepositoryModel {
	dbConnect := GetConnection()

	pq := `SELECT id, image, description, user, lat, lng, datetime FROM report`

	rows, err := dbConnect.Query(pq)
	if err != nil {
		return []report.RepositoryModel{}
	}

	defer rows.Close()

	dbResponse := []report.RepositoryModel{}

	for rows.Next() {
		var db report.RepositoryModel
		rows.Scan(&db.ID, &db.Image, &db.Description, &db.User, &db.Latitude, &db.Longitude, &db.DateTime)
		dbResponse = append(dbResponse, db)
	}

	return dbResponse
}

func (r *DBSqlite) Update(id string, db report.RepositoryModel) error {
	dbConnect := GetConnection()

	pq := `UPDATE report SET image=$1, description=$2, lat=$3, lng=$4, datetime=$5 WHERE id=$6`

	stmt, err := dbConnect.Prepare(pq)

	if err != nil {
		return err
	}

	defer stmt.Close()

	db.DateTime = time.Now().Format(time.RFC3339)

	stmt.Exec(db.Image, db.Description, db.Latitude, db.Longitude, db.DateTime, id)

	return nil
}

func (db *DBSqlite) Delete(id string) error {
	dbConnect := GetConnection()

	pq := `DELETE FROM report WHERE id=$1`

	stmt, err := dbConnect.Prepare(pq)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		fmt.Println("Error: ", err)
		return errors.New("Report not found")
	}

	return nil
}
