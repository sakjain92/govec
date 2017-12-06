package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"strings"
	"bytes"
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

func exit(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(-1)
}

func check(err error, a ...interface{}) {
	if err != nil {
		exit(err, a)
	}
}

func parseFuncDecl(node *ast.FuncDecl) {
	err := Fprint(f, h, fset, node)
	check(err)

}

func runCommand(s string, command string, a ...string) {
	var out bytes.Buffer

	// Create the object file and the library
	cmd := exec.Command(command, a...)
	cmd.Stderr = &out

	err := cmd.Run()

	check(err, s, ":\n", out.String())
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
		exit("Usage: govec <filename.go>")
	}

	fileName := os.Args[1]

	_, err := os.Stat(fileName)
	check(err, "Usage: govec <filename.go>")

	dir := filepath.Dir(fileName)
	file := filepath.Base(fileName)

	filesuffix := file[(len(file)-3):]
	if filesuffix != ".go" {
		exit("Usage: govec <filename.go>")
	}

	newdir := dir + "/govec_build/"
	_, err = os.Stat(newdir);

	if os.IsNotExist(err) {
	    os.Mkdir(newdir, 0700)
	} else {
		check(err, "Couldn't make govec directory", newdir)
	}

	fileprefix := file[:(len(file)-3)]
	newfilename := newdir + fileprefix + ".ispc"
	headerfilename := newdir + fileprefix + ".h"
	objfilename := newdir + fileprefix + ".o"
	libfilename := newdir + "lib" + fileprefix + ".a"

	f, err = os.Create(newfilename)
	check(err, "Couldn't create file:", newfilename)
	defer f.Close()

	h, err = os.Create(headerfilename)
	check(err, "Couldn't create file:", headerfilename)
	defer h.Close()


	fset = token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	check(err)

	ast.Inspect(node, collectGoVecFunctions)

	// Create the object file and the library
	runCommand("Couldn't compile with ispc",
		   "ispc", "-O3", "--target=sse4-x2",
			"--arch=x86-64", newfilename, "-o",
			objfilename)

	runCommand("Couldn't name static library",
			"ar", "rcs", libfilename, objfilename)
}
