package scihub_test

import (
	"testing"

	"github.com/satslab/snt/pkg/client"
	"github.com/satslab/snt/pkg/client/scihub"
)

func TestNewSearchClient(t *testing.T) {
	a := client.BasicAuth{
		Username: "user",
		Password: "pwd",
	}

	s := scihub.NewSearchClient(a)

	if s == nil {
		t.Error("error while initializing search client")
	}

	expUrlStr := "https://scihub.copernicus.eu/apihub/search"
	if s.URL().String() != expUrlStr {
		t.Error("error building the base URL")
	}
}
