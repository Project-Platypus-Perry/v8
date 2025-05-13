package constants

import "database/sql/driver"

type UserRole string

const (
	AdminRole      UserRole = "admin"
	InstructorRole UserRole = "instructor"
	StudentRole    UserRole = "student"
)

type Permission string

const (
	// User permissions
	CreateUser Permission = "create:user"
	ReadUser   Permission = "read:user"
	UpdateUser Permission = "update:user"
	DeleteUser Permission = "delete:user"

	// Organization permissions
	CreateOrganization Permission = "create:organization"
	ReadOrganization   Permission = "read:organization"
	UpdateOrganization Permission = "update:organization"
	DeleteOrganization Permission = "delete:organization"

	// Classroom permissions
	CreateClassroom Permission = "create:classroom"
	ReadClassroom   Permission = "read:classroom"
	UpdateClassroom Permission = "update:classroom"
	DeleteClassroom Permission = "delete:classroom"

	// Other resource permissions can be added here
)

// RolePermissions maps roles to their allowed permissions
var RolePermissions = map[UserRole][]Permission{
	AdminRole: {
		CreateUser,
		ReadUser,
		UpdateUser,
		DeleteUser,
		CreateOrganization,
		ReadOrganization,
		UpdateOrganization,
		DeleteOrganization,
		CreateClassroom,
		ReadClassroom,
		UpdateClassroom,
		DeleteClassroom,
	},
	InstructorRole: {
		ReadUser,
		UpdateUser,
		CreateClassroom,
		ReadClassroom,
		UpdateClassroom,
		DeleteClassroom,
	},
	StudentRole: {
		ReadUser,
		UpdateUser,
		ReadClassroom,
	},
}

func (p *UserRole) Scan(value interface{}) error {
	*p = UserRole(value.(string))
	return nil
}

func (p UserRole) Value() (driver.Value, error) {
	return string(p), nil
}
