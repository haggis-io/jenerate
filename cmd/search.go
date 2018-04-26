package cmd

import (
	"fmt"
	"github.com/haggis-io/jenerate/pkg/render"
	"github.com/haggis-io/jenerate/pkg/service"
	"github.com/haggis-io/jenerate/pkg/util"
	"github.com/haggis-io/registry/pkg/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func SearchAction() cli.ActionFunc {

	return func(context *cli.Context) error {

		all := context.Bool("all")

		if context.NArg() < 1 && !all {
			return MissingDocumentSearchArgErr
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

		if all {
			docs, err := documentService.GetAll()

			if err != nil {
				return err
			}

			render.GetRenderer(
				context.String("output")).
				PrettyPrint(util.JustNameAndVersionFromDocuments(docs))
			return nil
		}

		docs, err := documentService.Get(name)

		if err != nil {
			return err
		}

		docVers := util.ExtractVersionsFromDocumentSlice(docs)

		fmt.Printf("Document: %s\n", name)
		fmt.Println("Avaliable versions:")

		render.GetRenderer(
			context.String("output")).
			PrettyPrint(docVers)

		return nil

	}
}
