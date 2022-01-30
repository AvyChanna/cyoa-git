package cyoa

import (
	"fmt"
)

type ListCmd struct {
}

func (inp *ListCmd) Run(ctx *Context) error {
	fmt.Println("list")
	return nil
}
