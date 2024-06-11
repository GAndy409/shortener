package main

import (
	"net/http"

	"github.com/GAndy409/shortener/internal/app"
)

func main() {
	router := app.RouterInit()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
