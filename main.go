package main

import (
	"JsonCompareGolang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.TestRouter(router)

	router.Run("localhost:5000")
}
