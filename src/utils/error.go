package utils

import (
	"html/template"
	"net/http"
)

func ThrowError(err error, w http.ResponseWriter) bool {
	if err != nil && err.Error() != NO_ROWS_RESULT_SET {
		w.WriteHeader(http.StatusBadRequest)
		t, _ := template.New("error").Parse(err.Error())
		_ = t.Execute(w, nil)
		return true
	}
	return false
}
