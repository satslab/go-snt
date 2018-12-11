package client_test

import (
	"testing"

	"github.com/satslab/snt/pkg/client"
)

func TestBasicAuth(t *testing.T) {
	a := client.BasicAuth{
		Username: "user",
		Password: "pwd",
	}
}
