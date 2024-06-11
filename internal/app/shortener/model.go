package shortener

import (
	"math/rand"
	"time"
)

type Shortener struct {
	listUrls map[string]string
}

var (
	Shorts  Shortener
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	host    = "http://localhost:8080"
)

func init() {
	Shorts.listUrls = make(map[string]string)
	rand.Seed(time.Now().UnixNano())
}
