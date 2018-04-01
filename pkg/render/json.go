package render

import (
	"encoding/json"
	"fmt"
)

var _ Renderer = &jsonRenderer{}

type jsonRenderer struct{}

func (r *jsonRenderer) Print(objs ...interface{}) {
	for _, obj := range objs {
		b, err := json.Marshal(obj)

		if err != nil {
			continue
		}

		fmt.Println(string(b))
	}

}

func (r *jsonRenderer) PrettyPrint(objs ...interface{}) {
	for _, obj := range objs {
		b, err := json.MarshalIndent(obj, "", "    ")

		if err != nil {
			continue
		}

		fmt.Println(string(b))
	}

}
