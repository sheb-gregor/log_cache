package modules

import (
	"log_cache/config"

	"github.com/lancer-kit/armory/initialization"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func Init(c *cli.Context) (*config.Cfg, error) {
	cfg, err := config.ReadConfig(c.GlobalString(config.FlagConfig))
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config")
	}

	err = getModules(cfg).InitAll()
	if err != nil {
		return nil, errors.Wrap(err, "modules initialization failed")
	}

	return &cfg, nil
}

func getModules(cfg config.Cfg) initialization.Modules {
	return initialization.Modules{

	}

}
