package render

type Renderer interface {
	Print(...interface{})
	PrettyPrint(...interface{})
}
