package scihub_test

import (
	"testing"

	"github.com/satslab/snt/pkg/client/scihub"
)

func TestNewSearchClient(t *testing.T) {
	s := scihub.NewSearchClient("user", "pwd")

	if s == nil {
		t.Error("error while initializing search client")
	}

	expUrlStr := "https://scihub.copernicus.eu/apihub/search"
	if s.URL().String() != expUrlStr {
		t.Error("error building the base URL")
	}
}
