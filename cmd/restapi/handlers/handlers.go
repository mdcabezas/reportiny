package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/mdcabezas/reportiny/cmd/application"
	"github.com/mdcabezas/reportiny/internal/report"
	"github.com/mdcabezas/reportiny/pkg/web"
)

type Handlers struct {
	components *application.Components
}

func NewHandlers(components *application.Components) *Handlers {
	return &Handlers{components: components}
}

func (h *Handlers) HelloHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, fmt.Sprintf("%s, world!", r.URL.Path[1:]), http.StatusOK)
}

func (h *Handlers) GetReportHandler(w http.ResponseWriter, r *http.Request) error {

	reps := h.components.Report.GetReport()
	return web.EncodeJSON(w, reps, http.StatusOK)
}

func (h *Handlers) PostReportHandler(w http.ResponseWriter, r *http.Request) error {

	var body = report.RequestDTO{}

	err := web.DecodeJSON(r, &body)

	if err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	response, err := h.components.Report.PostReport(body)

	if err != nil {
		return web.EncodeJSON(w, err, http.StatusInternalServerError)
	}

	return web.EncodeJSON(w, response, http.StatusOK)
}

func (h *Handlers) UpdateReportHandler(w http.ResponseWriter, r *http.Request) error {

	var body = report.RequestDTO{}

	err := web.DecodeJSON(r, &body)

	if err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	id := chi.URLParam(r, "id")

	response, err := h.components.Report.UpdateReport(id, body)

	if err != nil {
		return web.EncodeJSON(w, err, http.StatusInternalServerError)
	}

	return web.EncodeJSON(w, response, http.StatusOK)
}

func (h *Handlers) DeleteReportHandler(w http.ResponseWriter, r *http.Request) error {

	// qParams, err := url.ParseQuery(r.URL.RawQuery)

	// if err != nil {
	// 	return web.NewError(http.StatusBadRequest, err.Error())
	// }

	// id := qParams.Get("id")

	id := chi.URLParam(r, "id")

	err := h.components.Report.DeleteReport(id)

	if err != nil {
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, fmt.Sprintf("Delete ID: %s", id), http.StatusOK)
}
