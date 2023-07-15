package main

import (
	"log"
	api "work6/kitex_demo/kitex_gen/api/pick"
)

func main() {
	svr := api.NewServer(new(PickImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
