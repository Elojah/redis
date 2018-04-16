package redis

import (
	"errors"
)

// Config is redis structure config.
type Config struct {
	Addr     string
	Password string
	DB       int
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return (c.Addr == rhs.Addr &&
		c.Password == rhs.Password &&
		c.DB == rhs.DB)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cVal, ok := fconf["addr"]
	if !ok {
		return errors.New("missing key addr")
	}
	if c.Addr, ok = cVal.(string); !ok {
		return errors.New("key addr invalid. must be string")
	}
	cVal, ok = fconf["password"]
	if !ok {
		return errors.New("missing key password")
	}
	if c.Password, ok = cVal.(string); !ok {
		return errors.New("key password invalid. must be string")
	}
	cVal, ok = fconf["db"]
	if !ok {
		return errors.New("missing key db")
	}
	if c.DB, ok = cVal.(int); !ok {
		return errors.New("key db invalid. must be int")
	}

	return nil
}
