package main

import (
	"net/http"

	"github.com/GAndy409/shortener/internal/app"
)

func main() {
	mux := http.NewServeMux()

	app.HandleInit(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
