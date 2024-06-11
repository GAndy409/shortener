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

	if shortUrl, ok := s.listUrls[fmt.Sprintf("%s/%s", host, u)]; ok {
		return true, shortUrl
	} else {
		return false, ""
	}
}

func (s *Shortener) CheckShortKey(shortKey string) (bool, string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %s", err)
		}
	}()

	shortUrl := fmt.Sprintf("%s/%s", host, shortKey)
	for url := range s.listUrls {
		if s.listUrls[url] == shortUrl {
			return true, url
		}
	}
	return false, ""
}

func (s *Shortener) ShortUrl(u string) string {
	if search, shortUrl := s.CheckUrl(u); search {
		return shortUrl
	}

	result := fmt.Sprintf("%s/%s", host, randSeq(8))
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
