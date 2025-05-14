package routes

import (
	"ecommerce-back/controllers"
	"ecommerce-back/middlewares"
	"ecommerce-back/repositories"
	"ecommerce-back/services"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine) {
	orderRepo := repositories.NewOrderRepository()
	err := orderRepo.CreateTableForOrder()
	if err != nil {
		log.Fatalf("Error creaeting order table: %v", err)
	}
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	orders := r.Group("/orders")
	{
		orders.POST("/", middlewares.AuthMiddleware(), middlewares.RoleAuthMiddleware("consumer"), orderController.PlaceOrder)
		orders.GET("/", middlewares.AuthMiddleware(), middlewares.RoleAuthMiddleware("consumer"), orderController.GetMyOrders)
		orders.GET("/seller", middlewares.AuthMiddleware(), middlewares.RoleAuthMiddleware("seller"), orderController.GetSellerOrders)
	}
}
