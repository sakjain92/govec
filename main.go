package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"path/filepath"
)

const (
	govecPreamble = "_govec"
	uniformPreamble = "uniform_"
	rangeFunction = "govecRange"
	govecFuncPreamble = govecPreamble
	govecstructPreamble = govecPreamble
)

var f *os.File
var h *os.File
var fset *token.FileSet
var err error

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func parseFuncDecl(node *ast.FuncDecl) {
	err := Fprint(f, fset, node)
	check(err)

}

func collectGoVecFunctions(n ast.Node) bool {

	if n == nil {
		// Reached end
		return false
	}


	_, ok := n.(*ast.File)
	if ok {
		// Have to explore to find functions
		return true
	}

	ret, ok := n.(*ast.FuncDecl)
	if !ok {
		// Not a function
		// TODO: Currently not supporting function inside functions
		// Also they are anonymous
		return false
	}

	node := ret

	if !strings.HasPrefix(node.Name.Name, govecFuncPreamble) {
		// Not a function we are interested
		return false
	}

	node.Name.Name = node.Name.Name[1:]
	parseFuncDecl(node)

	// We already inspected inside of function
	return false
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: govec <filename.go>")
		os.Exit(1)
	}

	fileName := os.Args[1]

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		panic("File doesn't exist")
	}

	dir := filepath.Dir(fileName)
	file := filepath.Base(fileName)

	newdir := dir + "/govec/"
	if _, err := os.Stat(newdir); os.IsNotExist(err) {
	    os.Mkdir(newdir, 0700)
	}

	newfilename := newdir + file[:(len(file)-2)] + "ispc"
	headerfilename := newdir + file[:(len(file)-2)] + "h"

	f, err = os.Create(newfilename)
	check(err)
	defer f.Close()

	h, err = os.Create(headerfilename)
	check(err)
	defer h.Close()


	fset = token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	check(err)

	ast.Inspect(node, collectGoVecFunctions)

}
