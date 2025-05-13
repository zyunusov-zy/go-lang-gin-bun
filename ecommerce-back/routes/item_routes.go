package routes

import (
	"ecommerce-back/controllers"
	"ecommerce-back/middlewares"
	"ecommerce-back/repositories"
	"ecommerce-back/services"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(r *gin.Engine) {
	itemRepo := repositories.NewItemRepo()

	err := itemRepo.CreateTableItemsIfNotExists()
	if err != nil {
		log.Fatalf("Error creating items table: %v", err)
	}

	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)
	
	items := r.Group("/items")
	items.GET("/", itemController.ListItems)
	items.POST("/upload", middlewares.AuthMiddleware(), middlewares.RoleAuthMiddleware("seller"), itemController.UploadItems)


}