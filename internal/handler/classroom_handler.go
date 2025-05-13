package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/service"
)

type ClassroomHandler struct {
	ClassroomService service.ClassroomService
}

func NewClassroomHandler(classroomService service.ClassroomService) *ClassroomHandler {
	return &ClassroomHandler{ClassroomService: classroomService}
}

func (h *ClassroomHandler) CreateClassroom(c echo.Context) error {
	return nil
}

func (h *ClassroomHandler) GetClassroom(c echo.Context) error {
	return nil
}

func (h *ClassroomHandler) UpdateClassroom(c echo.Context) error {
	return nil
}

func (h *ClassroomHandler) DeleteClassroom(c echo.Context) error {
	return nil
}

func (h *ClassroomHandler) AddUserToClassroom(c echo.Context) error {
	return nil
}

func (h *ClassroomHandler) RemoveUserFromClassroom(c echo.Context) error {
	return nil
}

func (h *ClassroomHandler) ListUsersInClassroom(c echo.Context) error {
	return nil
}
