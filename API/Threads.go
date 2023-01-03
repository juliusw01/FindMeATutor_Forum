package API

import (
	"FindMeATutor_User_Service/MongoDB"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllThreads(ctx *gin.Context) {
	users, err := MongoDB.GetAllThreads()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func CreateThread(ctx *gin.Context) {
	var thread MongoDB.Thread
	if err := ctx.ShouldBindJSON(&thread); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := MongoDB.CreateThread(&thread)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Thread created"})
}

func GetThread(ctx *gin.Context) {
	email := ctx.Param("email")
	thread, err := MongoDB.GetThread(&email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, thread)
}

func UpdateThread(ctx *gin.Context) {
	var thread MongoDB.Thread
	if err := ctx.ShouldBindJSON(&thread); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := MongoDB.UpdateThread(&thread)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Thread updated"})
}

func RegisterThreadRoutes(router *gin.RouterGroup) {
	router.POST("/createThread", CreateThread)
	router.GET("/getThread/:_id", GetThread)
	router.GET("/getAllThreads", GetAllThreads)
	router.PATCH("/updateThread", UpdateThread)
}
