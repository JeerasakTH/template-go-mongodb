package main

import (
	"template-go-mongodb/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Router(r)

	r.Run() // default 8080
	// go func() {
	// 	if err := r.Run(":3000"); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// // Measure the elapsed time from the start of the program to the first HTTP request
	// start := time.Now()
	// resp, err := http.Get("http://localhost:3000/")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// elapsed := time.Since(start)
	// fmt.Printf("Time to first start: %v\n", elapsed)
}
