package controllers

import (
	"ecommerce-back/models"
	"ecommerce-back/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type OrderController struct {
	service services.OrderService
}


func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (oc *OrderController) PlaceOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	claims := c.MustGet("user").(jwt.MapClaims)
	order.BuyerID = int64(claims["user_id"].(float64))
	if err := oc.service.PlaceOrder(c.Request.Context(), &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not place order"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order placed successfully"})
}

func (oc *OrderController) GetMyOrders(c *gin.Context) {
	claims := c.MustGet("user").(jwt.MapClaims)
	userID := int64(claims["user_id"].(float64))

	orders, err := oc.service.GetBuyerOrders(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}


// not working 
func (oc *OrderController) GetSellerOrders(c *gin.Context) {
	claims := c.MustGet("user").(jwt.MapClaims)
	sellerID := int64(claims["user_id"].(float64))

	orders, err := oc.service.GetSellerOrders(c.Request.Context(), sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch seller orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}