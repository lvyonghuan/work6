package main

import (
	"work6/protoc/api"
	"work6/protoc/service"
)

func main() {
	go service.Service()
	api.InitRouter()
}
