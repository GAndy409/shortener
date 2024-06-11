package shortener

type Shortener struct {
	ListUrls map[string]string
}

var Shorts Shortener

func init() {
	Shorts.ListUrls = make(map[string]string)
}