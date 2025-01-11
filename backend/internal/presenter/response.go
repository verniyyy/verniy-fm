package presenter

import (
	"encoding/json"
	"net/http"
)

type ClientError struct {
	Message string
	Errors  []error
}

func NewClientError(message string, errors ...error) *ClientError {
	return &ClientError{
		Message: message,
		Errors:  errors,
	}
}

func (e *ClientError) Error() string {
	return e.Message
}

type ServerError struct{}

func OK(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusOK, data)
}

func BadRequest(w http.ResponseWriter, err error) {
	writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
}

func InternalServerError(w http.ResponseWriter, err error) {
	writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}

func writeJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	e := json.NewEncoder(w)
	if err := e.Encode(data); err != nil {
		panic(err)
	}
}
