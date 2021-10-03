package main

import (
	"movieapp/server"

	"github.com/gin-gonic/gin"
)

func main()  {
	g := gin.Default()

	r := g.Group("/")
	server.NewRoute(r)

	g.Run(":9000")
}