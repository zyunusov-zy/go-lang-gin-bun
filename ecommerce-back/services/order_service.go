package services

import (
	"context"
	"ecommerce-back/models"
	"ecommerce-back/repositories"

)

type OrderService interface {
	PlaceOrder(ctx context.Context, order *models.Order) error 
	GetBuyerOrders(ctx context.Context, buyerID int64) ([]models.Order, error)
	GetSellerOrders(ctx context.Context, sellerID int64) ([]models.Order, error)
}

type orderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{orderRepo: repo}
}

func (s *orderService)	PlaceOrder(ctx context.Context, order *models.Order) error {
	order.TotalPrice = float64(order.Quantity) * order.TotalPrice
	return s.orderRepo.Create(ctx, order)
}

func (s *orderService)	GetBuyerOrders(ctx context.Context, buyerID int64) ([]models.Order, error) {
	return s.orderRepo.GetByBuyer(ctx, buyerID)
}

func (s *orderService)	GetSellerOrders(ctx context.Context, sellerID int64) ([]models.Order, error) {
	return s.orderRepo.GetBySeller(ctx, sellerID)
}