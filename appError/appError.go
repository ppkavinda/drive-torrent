package appError

import (
	"encoding/json"
	"net/http"
)

// AppError : store errors, to send as json
type AppError struct {
	Error   error
	Message string
	Code    int
}

// WriteAsJSON : send the error as json
func WriteAsJSON(w http.ResponseWriter, err error, msg string, code, status int) {
	e := AppError{err, msg, code}
	js, _ := json.Marshal(e)
	if status != 0 {
		w.WriteHeader(status)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(js)
}
