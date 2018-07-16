package main

import (
	"flag"
)

func main() {
	var mount = flag.String("mount", "/resources", "Defines the path to the DXFs")
	var output = flag.String("output", "/output/output.wol", "Defines the path for the output")
	flag.Parse()
	files := getFiles(*mount, ".dxf")
	plates := getPlatesDetails(files)
	getPlatesDetailsTable(plates)
	generateWOL(plates, *output)
}
