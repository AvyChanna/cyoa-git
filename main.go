package main

import "github.com/alecthomas/kong"
import "github.com/AvyChanna/cyoa-git/cyoa"

func main() {
	ctx := kong.Parse(
		&cyoa.Commands,
		kong.ShortUsageOnError(),
	)
	err := ctx.Run(&cyoa.Context{})
	ctx.FatalIfErrorf(err)
}
