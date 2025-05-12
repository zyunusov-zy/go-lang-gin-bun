package repositories

import (
	"context"
	"database/sql"
	"ecommerce-back/config"
	"ecommerce-back/models"
	"errors"
	"log"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	CreateTableIfNotExists() error

}

type userRepo struct{}


func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (r *userRepo) CreateTableIfNotExists() error {
	// Drop the users table if it exists
	_, err := config.DB.NewDropTable().Model((*models.User)(nil)).IfExists().Exec(context.Background())
	if err != nil {
		return err
	}

	// Now create the table
	_, err = config.DB.NewCreateTable().Model((*models.User)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	log.Println("Users table has been created")
	return nil
}


func (r *userRepo) Create(user *models.User) error {
	_, err := config.DB.NewInsert().Model(user).Exec(context.Background())
	return err
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := config.DB.NewSelect().Model(user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
    	if err == sql.ErrNoRows {
        	return nil, errors.New("user not found")
    	}
    	return nil, err
	}

	return user, err
}