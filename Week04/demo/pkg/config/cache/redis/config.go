package redis

import (
	"github.com/golang/protobuf/ptypes/wrappers"
)

type Option func(c *Config) error

func (x *PbConfig) Options() []Option {
	return []Option{
		DialPassword(x.Password),
		DailReadTimeout(x.ReadTimeout),
	}
}

type Config struct {
	network string
	addr string
	database int32
	password string
	readTimeout int64
}

func DialPassword(v *wrappers.StringValue) func(c *Config) error {
	return func(c *Config) error {
		if v != nil {
			c.password = v.Value
		} else {
			c.password = "root"
		}
		return nil
	}
}

func DailReadTimeout(v *wrappers.Int64Value) func(c *Config) error {
	return func(c *Config) error {
		if v != nil {
			c.readTimeout = v.Value
		} else {
			c.readTimeout = 1000
		}
		return nil
	}
}

func NewConfig(x *PbConfig, options ...Option) (*Config, error) {
	c := &Config{
		network: x.Network,
		addr: x.Address,
		database: x.Database,
	}
	for _, opt := range options {
		err := opt(c)
		if err != nil {
			panic(err)
		}
	}
	return c, nil
}