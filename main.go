package main

import (
	"flag"

	"github.com/bagasjs/flux/core"
)

func main() {
    editor := core.NewDefaultFlux()
    file := flag.String("file", "", "File to edit")
    editor.Start(*file)
}
