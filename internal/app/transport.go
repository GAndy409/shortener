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
			jsonPresenter(w, err)
		}

		sUrl := shorts.Shorts.ShortUrl(responseString)
		jsonPresenter(w, sUrl)
	} else if r.Method == http.MethodGet {
		responseString, err := rString(r)
		if err != nil {
			jsonPresenter(w, err)
		}

		search, fullUrl := shorts.Shorts.CheckUrl(responseString)
		if search {
			jsonPresenter(w, fullUrl)
		} else {
			jsonPresenter(w, "not found")
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

func jsonPresenter(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(v)
}
