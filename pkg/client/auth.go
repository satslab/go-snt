package client

import "net/http"

type Auth interface {
	Client() *http.Client
}

type BasicAuth struct {
	Username string
	Password string
}

func (a *BasicAuth) Client() *http.Client {
	c := http.Client{}

	return &c
}

// func cloneRequest(r *http.Request) *http.Request {
// 	// Shallow copy of the struct.
// 	r2 := new(http.Request)
// 	*r2 = *r
// 	// Deep copy of the Header.
// 	r2.Header = make(http.Header)
// 	for k, s := range r.Header {
// 		r2.Header[k] = s
// 	}
// 	return r2
// }
