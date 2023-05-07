package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
)

type UserRepository interface {
	Addusers(ctx context.Context, user domain.Users) (domain.Users, error)
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)
}
