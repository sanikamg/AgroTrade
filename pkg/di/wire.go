//go:build wireinject
// +build wireinject

package di

import (
	http "golang_project_ecommerce/pkg/api"
	"golang_project_ecommerce/pkg/api/handler"
	"golang_project_ecommerce/pkg/config"
	"golang_project_ecommerce/pkg/db"
	"golang_project_ecommerce/pkg/repository"
	"golang_project_ecommerce/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserhandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
