package search

import (
	"fmt"
	"github.com/haggis-io/jenerate/cmd/errors"
	"github.com/haggis-io/jenerate/pkg/render"
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
			return errors.MissingDocumentSearchArgErr
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
				return errors.RegistryUnavaliableErr
			case codes.NotFound:
				return errors.DocumentNotFoundErr
			default:
				return errors.GenericExitErr(err)
			}
		}

		fmt.Printf("Document: %s\n", name)
		fmt.Println("Avaliable versions:")

		render.GetRenderer(
			context.String("output")).
			PrettyPrint(docVers)

		return nil

	}
}
