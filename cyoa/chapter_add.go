package cyoa

import (
	"fmt"
)

type AddCmd struct {
	Tag string `arg:"" help:"tag name for chapter"`
}

func (inp *AddCmd) Run(ctx *Context) error {
	fmt.Println("add", inp.Tag)
	return nil
}
