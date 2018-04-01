package cmd

import (
	"github.com/urfave/cli"

	"github.com/haggis-io/jenerate/cmd/describe"
	"github.com/haggis-io/jenerate/cmd/search"
	"github.com/haggis-io/jenerate/pkg/render"
)

var BaseCommands = []cli.Command{
	{
		Name:        "search",
		Action:      search.SearchAction(),
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
		Action:      describe.DescribeAction(),
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
