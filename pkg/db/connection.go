package db

import (
	"fmt"
	"golang_project_ecommerce/pkg/config"
	"golang_project_ecommerce/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err := db.AutoMigrate(
		domain.Users{},
		domain.AdminDetails{},
		domain.ProductDetails{},
		domain.Category{},
		domain.Image{},
		domain.Address{},
		domain.Cart_item{},
		domain.Order{},
		domain.Coupon{},
		domain.PaymentMethod{},
		domain.OrderReturn{},
	); err != nil {
		return nil, err
	}

	return db, dbErr
}
