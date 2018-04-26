package cmd

import (
	"github.com/haggis-io/jenerate/pkg/render"
	"github.com/haggis-io/jenerate/pkg/service"
	"github.com/haggis-io/registry/pkg/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func DescribeAction() cli.ActionFunc {

	return func(context *cli.Context) (err error) {

		if context.NArg() < 1 {
			return MissingDocumentDescribeArgErr
		}

		cc, err := grpc.Dial(context.GlobalString(RegistryGlobalFlag), grpc.WithInsecure())

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
			return MissingVersionArgErr
		}

		doc, err := documentService.GetStrict(name, version)

		if err != nil {
			return err
		}

		render.GetRenderer(
			context.String(OutputFlag)).
			PrettyPrint(doc)

		return nil

	}
}
