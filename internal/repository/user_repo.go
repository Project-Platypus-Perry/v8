package repository

import (
	"context"

	"github.com/gagan-gaurav/base/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*model.User, error)
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
func (r *userRepo) GetByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update a user
func (r *userRepo) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete a user
func (r *userRepo) Delete(ctx context.Context, id int) error {
	err := r.db.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// List all users
func (r *userRepo) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
