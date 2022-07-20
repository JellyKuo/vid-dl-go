package controller

import (
	"os"

	"github.com/JellyKuo/vid-dl-go/dl"
	"github.com/gin-gonic/gin"
)

func ListTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Task list",
	})
}

type CreateTaskRequest struct {
	Url string `json:"url"`
}

func CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := dl.Download(req.Url)
	c.Data(200, "text/plain", []byte(id.String()))
}

func ClearTask(c *gin.Context) {
	os.RemoveAll("downloads")
	c.JSON(200, gin.H{
		"message": "All task cleared",
	})
}
