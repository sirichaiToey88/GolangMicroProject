package handler

import "net/http"

func HandleErr(w http.ResponseWriter, r *http.Request) {
	ResponseWithError(w, 400, "Someting went wrong")
}
