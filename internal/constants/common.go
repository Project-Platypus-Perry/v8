package constants

import "database/sql/driver"

type Environment string

const (
	EnvLocal       Environment = "local"
	EnvDevelopment Environment = "development"
	EnvStaging     Environment = "staging"
	EnvProduction  Environment = "production"
)

func (p *Environment) Scan(value interface{}) error {
	*p = Environment(value.([]byte))
	return nil
}

func (p Environment) Value() (driver.Value, error) {
	return string(p), nil
}
