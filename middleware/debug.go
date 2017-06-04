package middleware

import (
	"net/http"
)

type Debug struct {
}

func NewDebug() *Debug {
	return &Debug{}
}

func (ua *Debug) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)
}