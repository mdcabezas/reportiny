package report

//go:generate mockgen -source=./repository.go -destination=mock/repository.go -package mock

type RepositoryModel struct {
	ID          string  `json:"id"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	User        string  `json:"user" validate:"required"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	DateTime    string  `json:"datetime"`
}

// Clase que representa un reporte
type Repository struct {
	Database []RepositoryModel `json:"reports"`
}

type IRepository interface {
	Create(repo RepositoryModel) (RepositoryModel, error)
	Read(id string) (RepositoryModel, error)
	Update(id string, repo RepositoryModel) error
	Delete(id string) error
	ReadAll() []RepositoryModel
}

func NewRepository() IRepository {
	return &Repository{}
}
