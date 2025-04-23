package constants

import "database/sql/driver"

type UserRole string

const (
	RoleAdmin      UserRole = "admin"
	RoleInstructor UserRole = "instructor"
	RoleStudent    UserRole = "student"
)

func (p *UserRole) Scan(value interface{}) error {
	*p = UserRole(value.([]byte))
	return nil
}

func (p UserRole) Value() (driver.Value, error) {
	return string(p), nil
}
