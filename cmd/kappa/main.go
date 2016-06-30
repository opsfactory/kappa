package main

import (
	"io/ioutil"
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
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatalf("Unable to read config file %s: %v.", configFile, err)
		}
		c, err := config.Parse(data)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		eng, err := engine.NewEngine(c)
		if err != nil {
			log.Fatalf("Unexpected error starting Kappa with the given configuration: %v", err)
		}
		go eng.Run()

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
