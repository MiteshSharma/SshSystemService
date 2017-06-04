package middleware

import (
	"net/http"
)

type Request struct {
}

func NewRequest() *Request {
	return &Request{}
}

func (ua *Request) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)
}
