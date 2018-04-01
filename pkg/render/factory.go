package render

import "strings"

var (
	jsonRenderer_  = &jsonRenderer{}
	plainRenderer_ = &plainRender{}
)

const (
	JSON  = "json"
	Plain = "plain"
)

func GetRenderer(name string) Renderer {
	switch strings.ToLower(name) {

	case JSON:
		return jsonRenderer_

	default:
		return plainRenderer_
	}
}
