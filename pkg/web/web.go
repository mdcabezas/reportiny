package web

import (
	"encoding/json"
	"net/http"
)

// EncodeJSON es una función para codificar datos en formato JSON y escribirlos en la respuesta HTTP.
func EncodeJSON(w http.ResponseWriter, data interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

// DecodeJSON es una función para decodificar datos JSON desde la solicitud HTTP.
func DecodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// NewError es una función para crear un nuevo error con un mensaje personalizado y un código de estado HTTP.
func NewError(statusCode int, message string) error {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// HTTPError es una estructura para representar errores HTTP.
type HTTPError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Message
}
