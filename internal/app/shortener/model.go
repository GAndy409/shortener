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
	r       *rand.Rand
)

func init() {
	Shorts.listUrls = make(map[string]string)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
