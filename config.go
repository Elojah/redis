package redis

import (
	"github.com/elojah/services"
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
		return services.ErrEmptyNamespace{}
	}

	cAddr, ok := fconf["addr"]
	if !ok {
		return services.ErrMissingKey{Key: "address"}
	}

	if c.Addr, ok = cAddr.(string); !ok {
		return services.ErrInvalidType{
			Key:    "address",
			Expect: "string",
			Value:  cAddr,
		}
	}

	cPassword, ok := fconf["password"]
	if !ok {
		return services.ErrMissingKey{Key: "password"}
	}

	if c.Password, ok = cPassword.(string); !ok {
		return services.ErrInvalidType{
			Key:    "password",
			Expect: "string",
			Value:  cPassword,
		}
	}

	cDB, ok := fconf["db"]
	if !ok {
		return services.ErrMissingKey{Key: "db"}
	}

	cDBFloat64, ok := cDB.(float64)
	if !ok {
		return services.ErrInvalidType{
			Key:    "password",
			Expect: "int",
			Value:  cDBFloat64,
		}
	}

	c.DB = int(cDBFloat64)

	return nil
}
