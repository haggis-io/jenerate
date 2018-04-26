package cmd

import (
	"github.com/haggis-io/jenerate/pkg/service"
	"github.com/haggis-io/jenerate/pkg/util"
	"github.com/haggis-io/registry/pkg/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func CreateAction() cli.ActionFunc {

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
			documentService = service.NewDocumentService(registryClient)
		)

		var documents []*api.Document
		for _, arg := range context.Args() {
			name, ver, err := SanitiseDocumentArg(arg)

			if err != nil {
				return DocumentCreateFormatArgError
			}

			documents = append(documents, &api.Document{
				Name:    name,
				Version: ver,
			})

		}

		docs, err := documentService.GetAll(documents...)

		if err != nil {
			return err
		}

		var snippetOrder []string
		for _, doc := range docs {
			snippets, err := util.ConstructDocumentOrder(doc)

			if err != nil {
				return err
			}
			snippetOrder = append(snippetOrder, snippets...)
		}

		return util.PrintPipeline(snippetOrder)
	}
}
