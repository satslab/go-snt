package scihub

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/satslab/snt/pkg/client"
)

const (
	defaultBaseURL = "https://scihub.copernicus.eu"
	APIhubPath     = "apihub"
	DhusPath       = "dhus"
	oSearchPath    = "search"
)

type SearchClient struct {
	url    *url.URL
	client *http.Client
}

func NewSearchClient(a client.Auth) *SearchClient {
	u, _ := url.Parse(strings.Join([]string{defaultBaseURL, APIhubPath, oSearchPath}, "/"))
	s := SearchClient{
		url:    u,
		client: a.Client(),
	}

	return &s
}

func (s *SearchClient) URL() *url.URL {
	fmt.Println(s.url)
	return s.url
}
