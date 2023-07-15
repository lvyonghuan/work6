package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Index(resp http.ResponseWriter, req *http.Request) {
	msg, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(msg))
	resp.Header().Set("Content-Type", "text")
	response := "yes"
	_, err = fmt.Fprint(resp, response)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("ok")
}

func Hello(resp http.ResponseWriter, req *http.Request) {
	log.Println("hello world")
}
