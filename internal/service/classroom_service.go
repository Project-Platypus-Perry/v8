package service

import (
	"context"

	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/repository"
)

type ClassroomService interface {
	CreateClassroom(ctx context.Context, classroom *model.Classroom) error
	GetClassroom(ctx context.Context, id string) (*model.Classroom, error)
	UpdateClassroom(ctx context.Context, classroom *model.Classroom) error
	DeleteClassroom(ctx context.Context, id string) error
	AddUserToClassroom(ctx context.Context, userID, classroomID, organizationID string) error
	RemoveUserFromClassroom(ctx context.Context, userID, classroomID string) error
}

type classroomService struct {
	classroomRepo repository.ClassroomRepository
}

func NewClassroomService(classroomRepo repository.ClassroomRepository) ClassroomService {
	return &classroomService{classroomRepo: classroomRepo}
}

func (s *classroomService) CreateClassroom(ctx context.Context, classroom *model.Classroom) error {
	return s.classroomRepo.CreateClassroom(ctx, classroom)
}

func (s *classroomService) GetClassroom(ctx context.Context, id string) (*model.Classroom, error) {
	return s.classroomRepo.GetClassroom(ctx, id)
}

func (s *classroomService) UpdateClassroom(ctx context.Context, classroom *model.Classroom) error {
	return s.classroomRepo.UpdateClassroom(ctx, classroom)
}

func (s *classroomService) DeleteClassroom(ctx context.Context, id string) error {
	return s.classroomRepo.DeleteClassroom(ctx, id)
}

func (s *classroomService) AddUserToClassroom(ctx context.Context, userID, classroomID, organizationID string) error {
	return s.classroomRepo.AddUserToClassroom(ctx, userID, classroomID, organizationID)
}

func (s *classroomService) RemoveUserFromClassroom(ctx context.Context, userID, classroomID string) error {
	return s.classroomRepo.RemoveUserFromClassroom(ctx, userID, classroomID)
}
