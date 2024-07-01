package report

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) Create(repo RepositoryModel) (RepositoryModel, error) {
	repo.ID = uuid.New().String()
	repo.DateTime = time.Now().Format(time.RFC3339)
	r.Database = append(r.Database, repo)
	return repo, nil
}

func (r *Repository) Read(id string) (RepositoryModel, error) {
	for _, v := range r.Database {
		if v.ID == id {
			return v, nil
		}
	}
	return RepositoryModel{}, errors.New("Report not found")
}

func (r *Repository) ReadAll() []RepositoryModel {
	fmt.Println("ReadAll", r.Database)
	return r.Database
}

func (r *Repository) Update(id string, repo RepositoryModel) error {

	for i := range r.Database {
		if r.Database[i].ID == id {
			r.Database[i].Image = repo.Image
			r.Database[i].Description = repo.Description
			r.Database[i].Latitude = repo.Latitude
			r.Database[i].Longitude = repo.Longitude
			r.Database[i].DateTime = time.Now().Format(time.RFC3339)
		}
	}

	return nil
}

func (r *Repository) Delete(id string) error {

	for i, v := range r.Database {
		if v.ID == id {
			r.Database = append(r.Database[:i], r.Database[i+1:]...)
			return nil
		}

	}

	return errors.New("Report not found")
}
