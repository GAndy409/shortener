package app

import (
	"encoding/json"
	"io"
	"net/http"

	shorts "github.com/GAndy409/shortener/internal/app/shortener"
)

func HandleInit(mux *http.ServeMux) {
	mux.HandleFunc("/", Hello)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		responseString, err := rString(r)
		if err != nil {
			jsonPresenter(w, http.StatusBadRequest, err)
		}

		sUrl := shorts.Shorts.ShortUrl(responseString)
		jsonPresenter(w, http.StatusCreated, sUrl)
	} else if r.Method == http.MethodGet {
		responseString, err := rString(r)
		if err != nil {
			jsonPresenter(w, http.StatusBadRequest, err)
		}

		search, fullUrl := shorts.Shorts.CheckUrl(responseString)
		if search {
			jsonPresenter(w, http.StatusTemporaryRedirect, fullUrl)
		} else {
			jsonPresenter(w, http.StatusBadRequest, "not found")
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

func jsonPresenter(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(v)
}
