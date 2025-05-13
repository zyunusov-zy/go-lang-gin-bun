package controllers

import (
	"ecommerce-back/models"
	"ecommerce-back/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ItemController struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (ic *ItemController) UploadItems(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userClaims := c.MustGet("user").(jwt.MapClaims)
	item.SellerID = int64(userClaims["user_id"].(float64))

	if err := ic.service.Upload(c.Request.Context(), &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": "Item uploaded successfully"})
}

func (ic *ItemController) ListItems(c *gin.Context) {
	items, err :=ic.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get items"})
		return
	}
	c.JSON(http.StatusOK, items)
}