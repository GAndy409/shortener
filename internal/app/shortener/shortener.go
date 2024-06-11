package shortener

import (
	"fmt"
	"math/rand"
)

func (s *Shortener) CheckUrl(u string) (bool, string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %s", err)
		}
	}()

	if v, ok := s.listUrls[u]; ok {
		return true, v
	} else {
		return false, ""
	}
}

func (s *Shortener) ShortUrl(u string) (string) {
	result := randSeq(8)
	s.listUrls[u] = result
	return result
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
