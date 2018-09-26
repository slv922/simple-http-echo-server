package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	portPtr := flag.String("port", "5001", "Listen and serve port.")

	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))
		c.JSON(http.StatusOK, JSON(buf[0:n]))
	})

	fmt.Println(fmt.Sprintf("\nSimple HTTP Echo Server V0.1 running at: http://localhost:%s\n", *portPtr))
	fmt.Println("GET /ping \n\t - Echo pong for testing.")
	fmt.Println("POST /post \n\t - Show post content.")

	r.Run(fmt.Sprintf(":%s", *portPtr)) // listen and serve on 0.0.0.0:8080
}
