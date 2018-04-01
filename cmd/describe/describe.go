package describe

import (
	"github.com/haggis-io/jenerate/cmd/errors"
	"github.com/haggis-io/jenerate/pkg/render"
	"github.com/haggis-io/jenerate/pkg/service"
	"github.com/haggis-io/registry/pkg/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DescribeAction() cli.ActionFunc {

	return func(context *cli.Context) error {

		if context.NArg() < 1 {
			return errors.MissingDocumentDescribeArgErr
		}

		cc, err := grpc.Dial(context.GlobalString("registry"), grpc.WithInsecure())

		if err != nil {
			return err
		}

		defer cc.Close()

		var (
			registryClient  = api.NewRegistryClient(cc)
			name            = context.Args().First()
			version         = context.Args().Get(1)
			documentService = service.NewDocumentService(registryClient)
		)

		if version == "" {
			return errors.MissingVersionArgErr
		}

		doc, err := documentService.Get(name, version)

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

		render.GetRenderer(
			context.String("output")).
			PrettyPrint(doc)

		return nil

	}
}
