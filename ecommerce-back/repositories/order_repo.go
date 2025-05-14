package repositories

import (
	"context"
	"ecommerce-back/config"
	"ecommerce-back/models"
	"log"
)

type OrderRepository interface {
	CreateTableForOrder() error
	Create(ctx context.Context, order *models.Order) error
	GetByBuyer(ctx context.Context, buyerID int64) ([]models.Order, error)
	GetBySeller(ctx context.Context, sellerID int64) ([]models.Order, error)
}

type orderRepo struct{}

func NewOrderRepository() OrderRepository{
	return &orderRepo{}
}

func (r *orderRepo) CreateTableForOrder() error {
	_, err := config.DB.NewCreateTable().Model((*models.Order)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	log.Println("Users table has been created")
	return nil
}


func (r *orderRepo)	Create(ctx context.Context, order *models.Order) error{
	_, err := config.DB.NewInsert().Model(order).Exec(ctx)
	return err
}

func (r *orderRepo)	GetByBuyer(ctx context.Context, buyerID int64) ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.NewSelect().Model(&orders).Where("buyer_id = ?", buyerID).Order("created_at DESC").Scan(ctx)
	return orders, err
}

func (r *orderRepo)	GetBySeller(ctx context.Context, sellerID int64) ([]models.Order, error) {
	var orders []models.Order
	query := `
		SELECT o. * FROM orders o
		JOIN items i ON o.item_id = i.id
		WHERE i.seller_id = ?
		ORDER BY 0.created_at DESC 
	`
	err := config.DB.NewRaw(query, sellerID).Scan(ctx, orders)
	return orders, err
}