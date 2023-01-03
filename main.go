package main

import (
	"FindMeATutor_User_Service/API"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	basePath := server.Group("/v1")
	API.RegisterThreadRoutes(basePath)
	err := server.Run()
	if err != nil {
		return
	}
}
