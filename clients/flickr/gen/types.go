package main

import (
	"flag"
	"os"

	. "github.com/attic-labs/noms/dbg"
	"github.com/attic-labs/noms/nomgen"
	"github.com/attic-labs/noms/types"
)

func main() {
	outFile := flag.String("o", "", "output file")
	flag.Parse()
	if *outFile == "" {
		flag.Usage()
		return
	}

	f, err := os.OpenFile(*outFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	Chk.NoError(err)
	ng := nomgen.New(f)

	photo := ng.AddType(types.NewMap(
		types.NewString("$type"), types.NewString("noms.StructDef"),
		types.NewString("$name"), types.NewString("Photo"),
		types.NewString("id"), types.NewString("string"),
		types.NewString("title"), types.NewString("string"),
		types.NewString("url"), types.NewString("string"),
		types.NewString("image"), types.NewString("blob")))

	photoSet := ng.AddType(types.NewMap(
		types.NewString("$type"), types.NewString("noms.SetDef"),
		types.NewString("elem"), photo))

	photoset := ng.AddType(types.NewMap(
		types.NewString("$type"), types.NewString("noms.StructDef"),
		types.NewString("$name"), types.NewString("Photoset"),
		types.NewString("id"), types.NewString("string"),
		types.NewString("title"), types.NewString("string"),
		types.NewString("photos"), photoSet))

	photosetSet := ng.AddType(types.NewMap(
		types.NewString("$type"), types.NewString("noms.SetDef"),
		types.NewString("elem"), photoset))

	ng.AddType(types.NewMap(
		types.NewString("$type"), types.NewString("noms.StructDef"),
		types.NewString("$name"), types.NewString("User"),
		types.NewString("id"), types.NewString("string"),
		types.NewString("name"), types.NewString("string"),
		types.NewString("oAuthToken"), types.NewString("string"),
		types.NewString("oAuthSecret"), types.NewString("string"),
		types.NewString("photosets"), photosetSet))

	ng.WriteGo("main")
}