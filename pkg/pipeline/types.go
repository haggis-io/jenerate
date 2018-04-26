package pipeline

import "regexp"

var (
	LibraryRegex     = regexp.MustCompile(`(?s)(^@Library[\s.+]?\((.*)\)).*$`)
	StageRegex       = regexp.MustCompile(`(?s)(^stage[\s.+]?\{(.*)\}).*$`)
	PipelineTemplate = `{{ .Library }}
node {
{{range .Stages }}
    {{.}}
{{end}}
}`
)

type Pipeline struct {
	Library string
	Stages  []string
}
