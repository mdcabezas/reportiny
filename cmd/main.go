package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mdcabezas/reportiny/cmd/application"
	"github.com/mdcabezas/reportiny/cmd/restapi/handlers"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app := chi.NewRouter()

	component := application.NewComponents()
	h := handlers.NewHandlers(component)

	app.Post("/hello", adapt(h.HelloHandler))
	app.Get("/report", adapt(h.GetReportHandler))
	app.Post("/report", adapt(h.PostReportHandler))
	app.Delete("/report/{id}", adapt(h.DeleteReportHandler))
	app.Put("/report/{id}", adapt(h.UpdateReportHandler))

	return http.ListenAndServe(":8080", app)
}

// Adapt whith error
func adapt(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
