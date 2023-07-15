package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
	"work6/kitex_demo/kitex_gen/api"
	"work6/kitex_demo/kitex_gen/api/pick"
)

func main() {
	c, err := pick.NewClient("example-server", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	var str = "111"
	req := &api.Request{
		Count:    0,
		Name:     "hello",
		Age:      11,
		NickName: &str,
	}
	resp, err := c.Pick(context.Background(), req, callopt.WithConnectTimeout(30*time.Second))
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
