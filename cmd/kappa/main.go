package main

import (
	"os"

	"github.com/opsfactory/kappa/config"
	"github.com/opsfactory/kappa/engine"
	"github.com/opsfactory/kappa/version"

	log "github.com/Sirupsen/logrus"

	"github.com/urfave/cli"
)

var (
	configFile string
)

func main() {
	app := cli.NewApp()
	app.Name = "kappa"
	app.Version = version.FullVersion()
	app.Author = "@opsfactory"
	app.Usage = "Native docker autoscaling for the most popular orchestration frameworks."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, C",
			Usage: "Configuration file",
		},
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug logging",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		if c.String("config") != "" {
			configFile = c.String("config")
		} else {
			log.Fatal("You must provide a configuration file.")
		}

		return nil
	}
	app.Action = func(ctx *cli.Context) error {
		log.Infof("Reading config from %s.", configFile)
		cfg, err := config.NewConfigFromFile(configFile)
		if err != nil {
			log.Fatalf("Unexpected error parsing configuration: %v", err)
		}

		eng, err := engine.NewEngine(cfg)
		if err != nil {
			log.Fatalf("Unexpected error starting Kappa with the given configuration: %v", err)
		}
		return eng.Run()
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
