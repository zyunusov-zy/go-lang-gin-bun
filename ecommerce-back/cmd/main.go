package main

import (
	"ecommerce-back/config"
	"ecommerce-back/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	routes.AuthRoutes(r)

	r.Run(":" + config.GetPort())
}