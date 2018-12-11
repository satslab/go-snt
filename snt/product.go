package snt

import (
	"io"

	"github.com/satslab/go-snt/snt/search"
)

type Product struct {
	ID   string
	Name string
}

type Searcher interface {
	Search(q search.Query) ([]Product, error)
}

type Downloader interface {
	Download(w io.Writer, p Product) (size int, err error)
}

func Search(s Searcher, q search.Query) ([]Product, error) {
	return nil, nil
}
