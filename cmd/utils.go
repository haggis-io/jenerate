package cmd

import (
	"github.com/haggis-io/jenerate/pkg/util"
	"github.com/urfave/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleErr(actionFunc cli.ActionFunc) cli.ActionFunc {
	return func(context *cli.Context) error {
		err := actionFunc(context)

		if err != nil {
			switch err {
			case util.CircularDependency:
				return CircularDependencyErr
			}
			switch status.Code(err) {
			case codes.Unavailable:
				return RegistryUnavaliableErr

			case codes.NotFound:
				return DocumentNotFoundErr
			default:
				return GenericExitErr(err)
			}
		}

		return nil

	}
}
