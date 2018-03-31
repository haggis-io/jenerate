package cmd

import (
	"github.com/urfave/cli"

	"github.com/haggis-io/jenerate/cmd/search"
)

var BaseCommands = []cli.Command{
	{
		Name:        "search",
		Action:      search.SearchAction(),
		Usage:       "Searches for a document",
		Description: "Prints a list of versions for a particular document in the registry",
		ArgsUsage:   "NAME",
	},
}
