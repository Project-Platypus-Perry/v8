package constants

import "database/sql/driver"

type UserRole string

const (
	RoleAdmin      UserRole = "admin"
	RoleInstructor UserRole = "instructor"
	RoleStudent    UserRole = "student"
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
	RoleAdmin: {
		CreateUser,
		ReadUser,
		UpdateUser,
		DeleteUser,
	},
	RoleInstructor: {
		ReadUser,
		UpdateUser,
	},
	RoleStudent: {
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
