package service

import (
	"github.com/haggis-io/registry/pkg/api"
)

func ExtractVersionsFromDocumentSlice(docs []*api.Document) []string {
	var out []string
	for _, doc := range docs {
		out = append(out, doc.Version)
	}

	return out
}
