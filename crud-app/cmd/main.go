package main

import (
	"crud-app/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default() // create a Gin instance

	// routes for CRUD operations

	routes.SetupRoutes(r)

	// Run the server on the port 8080
	r.Run(":8080")


}