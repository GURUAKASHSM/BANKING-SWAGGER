package main

import (
	"fmt"
	"log"
	"mongoapi/router"
)

// @title Customer
// @version 1.0
// @description Your API Description
// @host localhost:8080
// @BasePath /api
func main() {
	fmt.Println("Started Running")
	r := router.Router()
	log.Fatal(r.Run(":8080"))
	fmt.Println("Listening At PORT ... ")
}
