package cyoa

import (
	"fmt"
)

type RemoveCmd struct {
	Tag string `arg:"" help:"tag name for chapter"`
}

func (inp *RemoveCmd) Run(ctx *Context) error {
	fmt.Println("remove", inp.Tag)
	return nil
}