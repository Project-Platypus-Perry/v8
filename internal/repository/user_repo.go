package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/project-platypus-perry/v8/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// Create a new user
func (r *userRepo) CreateUser(ctx context.Context, user *model.User) error {
	user.ID = uuid.New().String()
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Get a user by ID
func (r *userRepo) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update a user
func (r *userRepo) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
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
func (r *userRepo) DeleteUser(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Organization").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
