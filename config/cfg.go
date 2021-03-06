package config

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const FlagConfig = "config"

// Cfg main structure of the app configuration.
type Cfg struct {
	API        api.Config  `yaml:"api"`
	Monitoring api.Config  `yaml:"monitoring"`
	Log        log.NConfig `yaml:"log"`
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (cfg Cfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.API, validation.Required),
		validation.Field(&cfg.Monitoring, validation.Required),
		validation.Field(&cfg.Log, validation.Required),
	)
}

func ReadConfig(path string) (Cfg, error) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return Cfg{}, errors.Wrap(err, "unable to read config file")
	}

	config := new(Cfg)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		return Cfg{}, errors.Wrap(err, "unable to unmarshal config file")
	}

	err = config.Validate()
	if err != nil {
		return Cfg{}, errors.Wrap(err, "invalid configuration")
	}

	_, err = log.Init(log.Config{
		AppName:  config.Log.AppName,
		Level:    config.Log.Level.Get(),
		Sentry:   config.Log.Sentry.Get(),
		AddTrace: config.Log.AddTrace,
		JSON:     config.Log.JSON,
	})
	if err != nil {
		return Cfg{}, errors.Wrap(err, "unable to init log")
	}

	return *config, nil
}
