package routes

import (
	"ecommerce-back/controllers"
	"ecommerce-back/repositories"
	"ecommerce-back/services"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	userRepo := repositories.NewUserRepo()

	err := userRepo.CreateTableIfNotExists()
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	auth := router.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
}