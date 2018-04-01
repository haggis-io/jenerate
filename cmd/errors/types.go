package errors

import (
	"fmt"
	"github.com/urfave/cli"
)

var (
	MissingDocumentSearchArgErr   = cli.NewExitError("Please specify the document you would like to search for.", 1)
	MissingDocumentDescribeArgErr = cli.NewExitError("Please specify the document you would like to describe.", 1)
	MissingVersionArgErr          = cli.NewExitError("Missing version.", 1)
	RegistryUnavaliableErr        = cli.NewExitError("Can't connect to the registry.", 1)
	DocumentNotFoundErr           = cli.NewExitError("Document not found.", 1)
	GenericExitErr                = func(err error) error { return cli.NewExitError(fmt.Sprintf("something went wrong, %s", err), 1) }
)
