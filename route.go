package main

import (
	"net/http"

	"github.com/JellyKuo/vid-dl-go/controller"
	"github.com/gin-gonic/gin"
)

func setRoute(r *gin.RouterGroup) {
	download := r.Group("task")
	download.GET("/list", controller.ListTask)
	download.POST("/create", controller.CreateTask)
	download.GET("/clear", controller.ClearTask)
	r.StaticFS("/view", http.Dir("downloads"))
}
