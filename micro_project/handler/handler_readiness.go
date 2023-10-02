package handler

import "net/http"

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	ResponseWithJSON(w, 200, struct{}{})
}
