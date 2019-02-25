// Package main serves as the entrypoint of this API.
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Write API Starting..")

	router := gin.Default()

	registerRoutes(router)

	router.Run()
}
