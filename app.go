package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

var portPtr = flag.String("port", "5001", "Listen and serve port.")

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func init() {
	flag.Parse()
}

func main() {

	// http.Handle("/foo", fooHandler)

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("%s %s %s\n", blue(r.Method), r.URL.Path, blue(r.Proto))

		for k, v := range r.Header {
			fmt.Printf("%s: %s\n", k, cyan(v[0]))
		}
		fmt.Println()
		buf := make([]byte, 1024)
		n, _ := r.Body.Read(buf)
		fmt.Println(string(buf[0:n]))
		fmt.Println()
		fmt.Fprintf(w, "ok")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *portPtr), nil))

	// gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.POST("/post", func(c *gin.Context) {
	// 	buf := make([]byte, 1024)
	// 	n, _ := c.Request.Body.Read(buf)
	// 	fmt.Println(string(buf[0:n]))
	// 	c.JSON(http.StatusOK, string(buf[0:n]))
	// })

	// fmt.Println(fmt.Sprintf("\nSimple HTTP Echo Server V0.1 running at: http://localhost:%s\n", *portPtr))
	// fmt.Println("GET /ping \n\t - Echo pong for testing.")
	// fmt.Println("POST /post \n\t - Show post content.")

	// r.Run(fmt.Sprintf(":%s", *portPtr)) // listen and serve on 0.0.0.0:8080

}
