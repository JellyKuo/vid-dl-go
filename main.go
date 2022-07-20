package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	route := r.Group("/")
	setRoute(route)
	r.Run()
}
