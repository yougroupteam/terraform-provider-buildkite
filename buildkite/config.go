package buildkite

import (
	"github.com/buildkite/go-buildkite/buildkite"
)

type Config struct {
	Token string
}

func (c *Config) Client() (*buildkite.Client, error) {
	config, err := buildkite.NewTokenConfig(c.Token, false)
	if err != nil {
		return nil, err
	}

	return buildkite.NewClient(config.Client()), nil
}
