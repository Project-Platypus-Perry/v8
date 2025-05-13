package repository

import (
	"context"

	"github.com/project-platypus-perry/v8/internal/model"
	"gorm.io/gorm"
)

type ClassroomRepository interface {
	CreateClassroom(ctx context.Context, classroom *model.Classroom) error
	GetClassroom(ctx context.Context, id string) (*model.Classroom, error)
	UpdateClassroom(ctx context.Context, classroom *model.Classroom) error
	DeleteClassroom(ctx context.Context, id string) error
	AddUserToClassroom(ctx context.Context, userID, classroomID string, organizationID string) error
	RemoveUserFromClassroom(ctx context.Context, userID, classroomID string) error
}

type classroomRepository struct {
	db *gorm.DB
}

func NewClassroomRepository(db *gorm.DB) ClassroomRepository {
	return &classroomRepository{db: db}
}

func (r *classroomRepository) CreateClassroom(ctx context.Context, classroom *model.Classroom) error {
	return r.db.Create(classroom).Error
}

func (r *classroomRepository) GetClassroom(ctx context.Context, id string) (*model.Classroom, error) {
	var classroom model.Classroom
	if err := r.db.Where("id = ?", id).First(&classroom).Error; err != nil {
		return nil, err
	}
	return &classroom, nil
}

func (r *classroomRepository) UpdateClassroom(ctx context.Context, classroom *model.Classroom) error {
	return r.db.Save(classroom).Error
}

func (r *classroomRepository) DeleteClassroom(ctx context.Context, id string) error {
	return r.db.Delete(&model.Classroom{}, id).Error
}

func (r *classroomRepository) AddUserToClassroom(ctx context.Context, userID, classroomID, organizationID string) error {
	return r.db.Create(&model.UsersClassrooms{
		UserID:         userID,
		ClassroomID:    classroomID,
		OrganizationID: organizationID,
	}).Error
}

func (r *classroomRepository) RemoveUserFromClassroom(ctx context.Context, userID, classroomID string) error {
	return r.db.Delete(&model.UsersClassrooms{}, "user_id = ? AND classroom_id = ?", userID, classroomID).Error
}
