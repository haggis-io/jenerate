package util

import (
	"errors"
	"github.com/haggis-io/registry/pkg/api"
)

var (
	CircularDependency = errors.New("circular dependency")
)

func ExtractVersionsFromDocumentSlice(docs []*api.Document) []string {
	var out []string
	for _, doc := range docs {
		out = append(out, doc.Version)
	}

	return out
}

func ConstructDocumentOrder(document *api.Document) ([]string, error) {

	if containsItself(document, document.GetDependencies()) {
		return nil, CircularDependency
	}

	var order []string

	if len(document.GetDependencies()) == 0 {
		return []string{document.GetSnippet().GetText()}, nil
	}

	return WalkDependencyTree(document, order), nil
}

func WalkDependencyTree(document *api.Document, currentOrder []string) []string {
	if len(document.GetDependencies()) == 0 {
		currentOrder = append(currentOrder, document.GetSnippet().Text)
	} else {
		for _, dep := range document.GetDependencies() {
			if containsItself(dep, dep.GetDependencies()) {
				continue
			}
			currentOrder = append(currentOrder, WalkDependencyTree(dep, currentOrder)...)
		}
		currentOrder = append(currentOrder, document.GetSnippet().Text)
	}

	return currentOrder
}

func containsItself(this *api.Document, dependencies []*api.Document) bool {

	if len(dependencies) <= 0 {
		return false
	}

	head := dependencies[0]

	if DocumentsEqual(this, head) {
		return true
	}

	return containsItself(this, dependencies[1:])

}

func DocumentsEqual(x, y *api.Document) bool {
	if x == y {
		return true
	}

	if x != nil && y != nil {
		return x.GetName() == y.GetName() && x.GetVersion() == y.GetVersion()
	}

	return false
}
