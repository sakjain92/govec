package main

import (
	"go/token"
	"go/parser"
	"os"
	"log"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "saxpy.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	/*
	fo, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	bufferedWriter := bufio.NewWriter(fo)
	*/

	err = Fprint(os.Stdout, fset, node)
	if err != nil {
	    log.Fatal(err)
	}
}
