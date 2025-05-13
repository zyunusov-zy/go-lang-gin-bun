package services

import (
	"context"
	"ecommerce-back/models"
	"ecommerce-back/repositories"
)

type ItemService interface{
	Upload(ctx context.Context, item *models.Item) error
	List(ctx context.Context) ([]models.Item, error)
}

type itemService struct {
	itemRepo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{itemRepo: repo}
}

func (s *itemService)	Upload(ctx context.Context,item *models.Item) error {
	return s.itemRepo.Create(ctx, item)
}

func (s *itemService) 	List(ctx context.Context) ([]models.Item, error) {
	return s.itemRepo.FetchAll(ctx)
}
