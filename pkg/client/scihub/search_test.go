package scihub_test

import (
	"testing"

	"github.com/satslab/snt/pkg/scihub"
)

func TestSearchWithoutParams(t *testing.T) {
	s := scihub.NewSearcher("user", "pwd")

	if s == nil {
		t.Error("error while initializing search client")
	}
}
