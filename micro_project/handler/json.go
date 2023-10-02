package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	ResponseWithJSON(w, code, errResponse{
		Error: msg,
	})

}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to maeshal JSON response %v\n", payload)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Write(dat)
}
