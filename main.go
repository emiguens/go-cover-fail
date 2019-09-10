package main

import (
	"github.com/emiguens/go-cover-fail/pkg1"
	"github.com/emiguens/go-cover-fail/pkg2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		hh1 := pkg1.HasHeader(c, "x-ttl")
		hh2 := pkg2.HasHeader(c, "x-ttl")

		c.JSON(200, gin.H{
			"message": "pong",
			"hh1":     hh1,
			"hh2":     hh2,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
