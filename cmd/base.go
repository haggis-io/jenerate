package cmd

import (
	"github.com/haggis-io/jenerate/pkg/render"
	"github.com/urfave/cli"
)

var BaseCommands = []cli.Command{
	{
		Name:        "create",
		Action:      HandleErr(CreateAction()),
		Usage:       "Creates a Jenkinsfile from documents",
		Description: "Creates a Jenkinsfile from document snippets",
		ArgsUsage:   "NAME:VERSION NAME:VERSION...",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Usage: "Output format `[json|plain]`",
				Value: render.Plain,
			},
		},
	},
	{
		Name:        "search",
		Action:      HandleErr(SearchAction()),
		Usage:       "Searches for a document",
		Description: "Prints a list of versions for a particular document in the registry",
		ArgsUsage:   "NAME",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Usage: "Output format `[json|plain]`",
				Value: render.Plain,
			},
		},
	},
	{
		Name:        "describe",
		Action:      HandleErr(DescribeAction()),
		Usage:       "Describes a document",
		Description: "Prints detailed information of a particular document in the registry",
		ArgsUsage:   "NAME VERSION",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Usage: "Output format `[json|plain]`",
				Value: render.Plain,
			},
		},
	},
}
