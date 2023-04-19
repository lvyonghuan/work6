package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"work6/microservice_test/protoc/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Plus(c *gin.Context) {
	a := c.PostForm("A")
	b := c.PostForm("B")
	a_num, err := strconv.ParseFloat(a, 32)
	if err != nil {
		return
	}
	b_num, err := strconv.ParseFloat(b, 32)
	if err != nil {
		return
	}
	a_num_32 := float32(a_num)
	b_num_32 := float32(b_num)
	con, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer con.Close()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	cli := rpc.NewServerClient(con)
	re, err := cli.Serv(context.Background(), &rpc.Request{
		Num_A: a_num_32,
		Num_B: b_num_32,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("rpc:", re.Result)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   re.Result,
	})
}
