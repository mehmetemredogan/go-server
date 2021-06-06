package name

import (
	"encoding/json"
	"net/http"
)

type Output struct {
	Status	string	`json:"status"`
	Message	string	`json:"message"`
}

func NameService(w http.ResponseWriter, r *http.Request, name string) bool {
	output	:= Output {
		Status: "SUCCESS",
		Message: "Hi " + name,
	}

	out, err	:= json.Marshal(output)

	w.WriteHeader(200)
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	return true
}