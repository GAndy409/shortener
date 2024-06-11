package app

import (
	"encoding/json"
	"net/http"
)

func HandleInit(mux *http.ServeMux) {
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/", Hello)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	jsonPresenter(w, "Hello")
}

func jsonPresenter(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(v)
}
