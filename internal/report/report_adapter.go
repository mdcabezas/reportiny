package report

import "fmt"

// NewReport crea una nueva instancia de Reporte.
func NewReport(r IRepository) IReport {
	fmt.Println("NewReport Contructor")
	return &Report{
		Repository: r,
	}
}

// func NewReport(db IDBSqlite) IReport {
// 	fmt.Println("NewReport Contructor")
// 	return &Report{
// 		DB: db,
// 	}
// }

// GetReport obtiene los reportes.
func (r *Report) GetReport() *[]ResponseDTO {

	all := r.Repository.ReadAll()
	// all := r.DB.ReadAll()

	var response []ResponseDTO

	for _, v := range all {
		response = append(response, ResponseDTO{
			ID:          v.ID,
			Description: v.Description,
		})
	}

	return &response
}

// PostReport agrega reporte(s).
func (r *Report) PostReport(b RequestDTO) (ResponseDTO, error) {

	v, err := r.Repository.Create(RepositoryModel{
		Image:       b.Image,
		Description: b.Description,
		User:        b.User,
		Latitude:    b.Latitude,
		Longitude:   b.Longitude,
	})

	// v, err := r.DB.Create(DBModel{
	// 	Image:       b.Image,
	// 	Description: b.Description,
	// 	User:        b.User,
	// 	Latitude:    b.Latitude,
	// 	Longitude:   b.Longitude,
	// })

	if err != nil {
		return ResponseDTO{}, err
	}

	res := ResponseDTO{
		ID:          v.ID,
		Description: v.Description,
	}

	return res, nil
}

func (r *Report) UpdateReport(id string, b RequestDTO) (ResponseDTO, error) {

	err := r.Repository.Update(id, RepositoryModel{
		Image:       b.Image,
		Description: b.Description,
		User:        b.User,
		Latitude:    b.Latitude,
		Longitude:   b.Longitude,
	})

	// err := r.DB.Update(id, DBModel{
	// 	Image:       b.Image,
	// 	Description: b.Description,
	// 	User:        b.User,
	// 	Latitude:    b.Latitude,
	// 	Longitude:   b.Longitude,
	// })

	if err != nil {
		return ResponseDTO{}, err
	}

	return ResponseDTO{ID: id, Description: b.Description}, nil
}

func (r *Report) DeleteReport(id string) error {
	return r.Repository.Delete(id)
	// return r.DB.Delete(id)
}
