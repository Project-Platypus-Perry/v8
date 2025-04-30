package model

import "github.com/project-platypus-perry/v8/internal/constants"

type UserInvite struct {
	Name           string             `json:"Name" validate:"required,min=2,max=50"`
	Email          string             `json:"Email" validate:"required,email"`
	Phone          string             `json:"Phone" validate:"required,e164"`
	Role           constants.UserRole `json:"Role" validate:"required"`
	OrganizationID string             `json:"OrganizationID" validate:"required"`
}

type UserInviteRequest struct {
	Users []UserInvite `json:"Users" validate:"required,min=1,dive"`
}

type PasswordResetRequest struct {
	Email string `json:"Email" validate:"required,email"`
}

type PasswordResetConfirm struct {
	Token       string `json:"Token" validate:"required"`
	NewPassword string `json:"NewPassword" validate:"required,min=8"`
}
