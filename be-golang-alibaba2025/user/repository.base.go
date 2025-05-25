package user

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	FindUserByID(ctx context.Context, id int) (user *User, err error)
	FindUserPreferenceByID(ctx context.Context, id int) (userPreference *UserPreference, err error)
	UpsertUserPreference(ctx context.Context, userPreference *UserPreference) (err error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{DB: db}
}
