package main

import (
	"log"
	"work6/gotuber-test-demo/config"
	nlp "work6/gotuber-test-demo/kitex_gen/nlp/nlp"
)

func main() {
	svr := nlp.NewServer(new(NlpImpl))

	config.InitGPTConfig()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
