package shortener

import (
	"fmt"
)

func (s *Shortener) CheckUrl(u string) (bool, string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %s", err)
		}
	}()

	if shortURL, ok := s.listUrls[u]; ok {
		return true, shortURL
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

	shortURL := fmt.Sprintf("%s/%s", host, shortKey)
	for url := range s.listUrls {
		if s.listUrls[url] == shortURL {
			return true, url
		}
	}
	return false, ""
}

func (s *Shortener) ShortURL(u string) string {
	if search, shortURL := s.CheckUrl(u); search {
		return shortURL
	}

	result := fmt.Sprintf("%s/%s", host, randSeq(8))
	s.listUrls[u] = result
	return result
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
