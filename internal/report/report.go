package report

//go:generate mockgen -source=./report.go -destination=mock/report.go -package mock

// Definición de la estructura Reporte
type RequestDTO struct {
	Image       string  `json:"image"`
	Description string  `json:"description"`
	User        string  `json:"user" validate:"required"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
}

type ResponseDTO struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// IReport es una interfaz que define los métodos de un Reporte
type IReport interface {
	GetReport() *[]ResponseDTO
	PostReport(RequestDTO) (ResponseDTO, error)
	DeleteReport(id string) error
	UpdateReport(id string, b RequestDTO) (ResponseDTO, error)
}

type Report struct {
	Repository IRepository
	// DB         IDBSqlite
}
