package main

import (
	"work6/microservice_test/protoc/api"
	"work6/microservice_test/protoc/service"
)

func main() {
	go service.Service()
	api.InitRouter()
}
