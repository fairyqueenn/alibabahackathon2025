package user

import "context"

type Service interface {
	GetUserByID(ctx context.Context, id int) (user *User, err error)
	GetUserPreferenceByID(ctx context.Context, id int) (userPreference *UserPreference, err error)
	UpdateUserPreferenceRequest(ctx context.Context, id int, req UpdateUserPreferenceRequest) (err error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) *service {
	return &service{repo: repository}
}
