package modules

import (
	"log_cache/config"
	"log_cache/metrics"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func Init(c *cli.Context) (*config.Cfg, error) {
	cfg, err := config.ReadConfig(c.GlobalString(config.FlagConfig))
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config")
	}

	metrics.Init()
	return &cfg, nil
}
