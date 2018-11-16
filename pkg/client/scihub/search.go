package scihub

type SearchClient struct {
	url string
}

func NewSearcher(u string, p string) *SearchClient {
	s := SearchClient{
		url: "a url",
	}

	return &s
}
