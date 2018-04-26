package main

import (
	"fmt"
	"github.com/haggis-io/jenerate/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const (
	name        = "jenerate"
	description = "Generates Jenkinsfiles from documents stored in a registry"
)

var (
	commit  string
	version string
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Commands = cmd.BaseCommands
	app.Version = fmt.Sprintf("%s (%s)", version, commit)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   cmd.CreateFlagName(cmd.RegistryGlobalFlag, cmd.RegistryGlobalShortFlag),
			Usage:  "Registry address",
			Value:  "127.0.0.1:8080",
			EnvVar: "REGISTRY_ADDR",
		},
	}

	app.Run(os.Args)
}
