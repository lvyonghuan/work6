package logic

import (
	"fmt"
	"net/http"
	"work6/web_test/service"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		service.Index(w, r)
	case "/hello":
		service.Hello(w, r)
	default:
		fmt.Fprint(w, "404")
	}
}
