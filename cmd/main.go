package cmd

import (
	"log_cache/config"
	"log_cache/workers"

	"github.com/lancer-kit/uwe/v2"
	"github.com/urfave/cli"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		serveCmd(),
		uwe.CliCheckCommand(config.AppInfo(), func(_ *cli.Context) []uwe.WorkerName {
			return []uwe.WorkerName{workers.APIServer}
		}),
	}
}

func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  config.FlagConfig + ", c",
			Value: "./config.yaml",
		},
	}
}
