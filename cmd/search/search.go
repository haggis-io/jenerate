package search

import (
	"fmt"
	"github.com/haggis-io/jenerate/pkg/service"
	"github.com/haggis-io/registry/pkg/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SearchAction() cli.ActionFunc {

	return func(context *cli.Context) error {

		if context.NArg() < 1 {
			return cli.NewExitError("Please specify the document you would like to search for.", 1)
		}

		cc, err := grpc.Dial(context.GlobalString("registry"), grpc.WithInsecure())

		if err != nil {
			return err
		}

		defer cc.Close()

		var (
			registryClient  = api.NewRegistryClient(cc)
			name            = context.Args().First()
			documentService = service.NewDocumentService(registryClient)
		)

		docVers, err := documentService.GetAll(name)

		if err != nil {
			switch status.Code(err) {
			case codes.Unavailable:
				return cli.NewExitError("Can't connect to the registry.", 1)

			default:
				return cli.NewExitError(fmt.Sprintf("something went wrong, %s", err), 1)
			}
		}

		if len(docVers) == 0 {
			return cli.NewExitError("No document exist with that name or the document is hasn't been accepted yet...", 1)
		}

		fmt.Printf("Document: %s\n", name)
		fmt.Println("Avaliable versions:")

		for _, ver := range docVers {
			fmt.Println(ver)
		}

		return nil

	}
}
