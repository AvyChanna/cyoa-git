package cyoa

import (
	"fmt"
)

type UpdateCmd struct {
	Tag string `arg:"" help:"tag name for chapter"`
}

func (inp *UpdateCmd) Run(ctx *Context) error {
	fmt.Println("update", inp.Tag)
	return nil
}
