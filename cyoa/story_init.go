package cyoa

import (
	"os"

	"github.com/go-git/go-git/v5"
)

type InitCmd struct{}

func (inp *InitCmd) Run(ctx *Context) error {
	ParseTags("2")
	ParseTags("2.3")
	ParseTags("2333...9")
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = git.PlainInit(cwd, false)
	return err
}
