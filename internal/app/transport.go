package app

import (
	"io"
	"net/http"

	shorts "github.com/GAndy409/shortener/internal/app/shortener"
	"github.com/gorilla/mux"
)

func RouterInit() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Hello)
	router.HandleFunc("/{id}", Hello)

	return router
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		responseString, err := rString(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		sUrl := shorts.Shorts.ShortUrl(responseString)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(sUrl))
	} else if r.Method == http.MethodGet {
		shortUrl := mux.Vars(r)["id"]
		search, fullUrl := shorts.Shorts.CheckShortKey(shortUrl)
		if search {
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Location", fullUrl)
			_, _ = w.Write([]byte(fullUrl))
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
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