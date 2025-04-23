package repository

import (
	"context"

	"github.com/gagan-gaurav/v8/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, id string, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

// Create a new user
func (r *userRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get a user by ID
func (r *userRepo) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update a user
func (r *userRepo) Update(ctx context.Context, id string, user *model.User) (*model.User, error) {
	var existingUser model.User
	err := r.db.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	err = r.db.Save(&existingUser).Error
	if err != nil {
		return nil, err
	}
	return &existingUser, nil
}

// Delete a user
func (r *userRepo) Delete(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
