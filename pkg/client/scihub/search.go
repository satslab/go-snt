package scihub

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://scihub.copernicus.eu"
	APIhubPath     = "apihub"
	DhusPath       = "dhus"
	oSearchPath    = "search"
)

type SearchClient struct {
	url      *url.URL
	username string
	password string
}

func NewSearchClient(username string, password string) *SearchClient {
	u, _ := url.Parse(strings.Join([]string{defaultBaseURL, APIhubPath, oSearchPath}, "/"))
	s := SearchClient{
		url:      u,
		username: username,
		password: password,
	}

	return &s
}

func (s *SearchClient) URL() *url.URL {
	fmt.Println(s.url)
	return s.url
}
