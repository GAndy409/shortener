package app

import (
	"io"
	"net/http"

	shorts "github.com/GAndy409/shortener/internal/app/shortener"
	"github.com/gorilla/mux"
)

func RouterInit() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", ShortingPost).Methods(http.MethodPost)
	router.HandleFunc("/{id}", ShortingGet).Methods(http.MethodGet)

	return router
}

func ShortingPost(w http.ResponseWriter, r *http.Request) {
	responseString, err := rString(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	sUrl := shorts.Shorts.ShortUrl(responseString)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(sUrl))
}

func ShortingGet(w http.ResponseWriter, r *http.Request) {
	shortUrl := mux.Vars(r)["id"]
	search, fullUrl := shorts.Shorts.CheckShortKey(shortUrl)
	if search {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", fullUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func rString(r *http.Request) (string, error) {
	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	responseString := string(responseData)
	return responseString, nil
}
