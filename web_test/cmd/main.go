package main

import (
	"net/http"
	"work6/web_test/logic"
)

func main() {
	//http.HandleFunc("/", service.Index)
	//http.HandleFunc("/hello", service.Hello)
	//http.ListenAndServe(":8080", nil)
	engine := new(logic.Engine)
	http.ListenAndServe(":8080", engine)
}
