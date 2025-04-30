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

	// Other resource permissions can be added here
)

// RolePermissions maps roles to their allowed permissions
var RolePermissions = map[UserRole][]Permission{
	AdminRole: {
		CreateUser,
		ReadUser,
		UpdateUser,
		DeleteUser,
	},
	InstructorRole: {
		ReadUser,
		UpdateUser,
	},
	StudentRole: {
		ReadUser,
	},
}

func (p *UserRole) Scan(value interface{}) error {
	*p = UserRole(value.(string))
	return nil
}

func (p UserRole) Value() (driver.Value, error) {
	return string(p), nil
}
