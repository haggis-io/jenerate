package cmd

import (
	"github.com/haggis-io/jenerate/pkg/render"
	"github.com/urfave/cli"
)

const (
	RegistryGlobalFlag      = "registry"
	RegistryGlobalShortFlag = "r"
	OutputFlag              = "output"
	OutputShortFlag         = "o"
	OutputFlagUsage         = "Output format `[json|plain]`"
	AllFlag                 = "all"
	AllShortFlag            = "a"
)

var BaseCommands = []cli.Command{
	{
		Name:        "create",
		Action:      HandleErr(CreateAction()),
		Usage:       "Creates a Jenkinsfile from documents",
		Description: "Creates a Jenkinsfile from document snippets",
		ArgsUsage:   "NAME:VERSION NAME:VERSION...",
	},
	{
		Name:        "search",
		Action:      HandleErr(SearchAction()),
		Usage:       "Searches for a document",
		Description: "Prints a list of versions for a particular document in the registry",
		ArgsUsage:   "NAME",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  CreateFlagName(OutputFlag, OutputShortFlag),
				Usage: OutputFlagUsage,
				Value: render.Plain,
			},
			cli.BoolFlag{
				Name:  CreateFlagName(AllFlag, AllShortFlag),
				Usage: "Print list of all documents stored in the registry",
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
				Name:  CreateFlagName(OutputFlag, OutputShortFlag),
				Usage: OutputFlagUsage,
				Value: render.Plain,
			},
		},
	},
}
