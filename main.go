package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

const (
	govecPreamble = "_govec"
	govecFuncPreamble = govecPreamble
	govecstructPreamble = govecPreamble
)

// typedef struct parserConfig{
var f *os.File
var fset *token.FileSet
var err error
// }

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func parseFuncDecl(node *ast.FuncDecl) {
	fmt.Println("")
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
		fmt.Println("Usage: my_tool <filenama.go>")
		os.Exit(1)
	}

	fileName := os.Args[1]

	f, err = os.Create(fileName[:(len(os.Args[1])-2)] + "cpp")
	check(err)
	defer f.Close()

	_, err = f.WriteString("#include<bits/stdc++.h>\n")
	check(err)


	fset = token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	check(err)

	/*
	fo, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	bufferedWriter := bufio.NewWriter(fo)
	*/

	ast.Inspect(node, collectGoVecFunctions)

}
