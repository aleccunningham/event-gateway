package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/codefresh-io/hermes/pkg/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "dukectl"
	app.Authors = []cli.Author{{Name: "Alec Cunningham", Email: "aleccunningham96@gmail.com"}}
	app.Version = version.HumanVersion
	app.EnableBashCompletion = true
	app.Usage = "Glue apis together"
	app.UsageText = fmt.Sprintf(`Configure triggers for Codefresh pipeline execution or start trigger manager server. Process "normalized" events and run Codefresh pipelines with variables extracted from events payload.
%s
dukectl respects following environment variables:
   - GOOGLE_SERVICE_ACCOUNT   - path to GCP JSON credentials

Copyright © Codefresh.io`, version.ASCIILogo)
	app.Before = before

	app.Commands = []cli.Command{
		serverCommand,
		triggerCommand,
		runnerCommand,
		triggerEventCommand,
		triggerTypeCommand,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "service-account, s",
			Usage:  "GCP service account",
			EnvVar: "GOOGLE_SERVICE_ACCOUNT",
		},
		cli.StringFlag{
			Name:   "log-level, l",
			Usage:  "set log level (debug, info, warning(*), error, fatal, panic)",
			Value:  "warning",
			EnvVar: "LOG_LEVEL",
		},
		cli.BoolFlag{
			Name:  "dry-run, x",
			Usage: "do not execute commands, just log",
		},
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "produce log in JSON format: Logstash and Splunk friendly",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func before(c *cli.Context) error {
	// set debug log level
	switch level := c.GlobalString("log-level"); level {
	case "debug", "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "info", "INFO":
		log.SetLevel(log.InfoLevel)
	case "warning", "WARNING":
		log.SetLevel(log.WarnLevel)
	case "error", "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "fatal", "FATAL":
		log.SetLevel(log.FatalLevel)
	case "panic", "PANIC":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
	// set log formatter to JSON
	if c.GlobalBool("json") {
		log.SetFormatter(&log.JSONFormatter{})
	}

	return nil
}
