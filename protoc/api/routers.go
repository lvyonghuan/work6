package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.GET("/plus", Plus)
	r.Run()
}
