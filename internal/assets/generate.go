// +build generate

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"
	"go.tmthrgd.dev/spork/internal/assets"
)

func main() {
	if err := vfsgen.Generate(assets.FileSystem, vfsgen.Options{
		Filename:        "internal/assets/vfsdata.go",
		PackageName:     "assets",
		BuildTags:       "!dev",
		VariableName:    "FileSystem",
		VariableComment: "FileSystem contains project assets.",
	}); err != nil {
		log.Fatalln(err)
	}
}
