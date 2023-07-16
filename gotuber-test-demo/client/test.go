package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
	"work6/gotuber-test-demo/kitex_gen/nlp"
	service "work6/gotuber-test-demo/kitex_gen/nlp/nlp"
)

func main() {
	c, err := service.NewClient("nlp", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	request := &nlp.Request{
		Choice:         1,
		RequestMessage: "hello world!",
	}
	resp, err := c.Handel(context.Background(), request, callopt.WithConnectTimeout(30*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.ResponseMessage)
}
