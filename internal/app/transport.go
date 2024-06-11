package app

import (
	"encoding/json"
	"net/http"
)

func HandleInit(mux *http.ServeMux) {
	mux.HandleFunc("/", Hello)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		jsonPresenter(w, "Hello Post")
	} else if r.Method == http.MethodGet {
		
	}
}

func jsonPresenter(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(v)
}
