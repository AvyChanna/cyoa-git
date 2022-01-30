package cyoa

import (
	"fmt"
)

type MoveCmd struct {
	Tag string `arg:"" help:"tag name for chapter"`
}

func (inp *MoveCmd) Run(ctx *Context) error {
	fmt.Println("move", inp.Tag)
	return nil
}
