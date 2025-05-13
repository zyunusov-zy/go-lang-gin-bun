package repositories

import (
	"context"
	"ecommerce-back/config"
	"ecommerce-back/models"
	"log"
)

type ItemRepository interface{
	Create(ctx context.Context, item *models.Item) error
	FetchAll(ctx context.Context) ([]models.Item,error)
	CreateTableItemsIfNotExists() error 
}

type itemRepo struct{}

func NewItemRepo() ItemRepository{
	return &itemRepo{}
}

func (r *itemRepo) CreateTableItemsIfNotExists() error {
	_, err := config.DB.NewCreateTable().Model((*models.Item)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}
	log.Println("Items table has been created")
	return err
}

func (r *itemRepo) Create(ctx context.Context, item *models.Item) error {
	_, err := config.DB.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *itemRepo) FetchAll(ctx context.Context) ([]models.Item,error) {
	var items []models.Item
	err := config.DB.NewSelect().Model(&items).Order("created_at DESC").Scan(ctx)
	return items, err
}

