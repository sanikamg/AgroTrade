package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
)

type UserUsecase interface {
	Register(ctx context.Context, user domain.Users) (domain.Users, error)
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
}
