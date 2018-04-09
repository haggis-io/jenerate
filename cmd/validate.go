package cmd

import (
	"github.com/haggis-io/registry/pkg/api"
	"strings"
)

func SanitiseDocumentArg(s string) (string, string, error) {
	if s == "" {
		return "", "", InvalidDocumentArgErr
	}

	possiblyHasVersion := strings.Split(s, ":")

	if len(possiblyHasVersion) >= 2 {
		return possiblyHasVersion[0], possiblyHasVersion[1], nil
	}

	return s, "", nil
}

func ProcessDocumentArgs(args []string) (out []*api.Document, err error) {
	for _, arg := range args {
		name, ver, err := SanitiseDocumentArg(arg)

		if err != nil {
			return out, err
		}

		out = append(out, &api.Document{
			Name:    name,
			Version: ver,
		})

	}

	return
}
