package render

import (
	"fmt"
)

var _ Renderer = &plainRender{}

type plainRender struct{}

func (r *plainRender) Print(objs ...interface{}) {
	fmt.Printf("%+v\n", objs...)
}

func (r *plainRender) PrettyPrint(objs ...interface{}) {
	for _, obj := range objs {
		fmt.Printf("%+v\n", obj)
	}

}
