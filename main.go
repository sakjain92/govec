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
	goprinter "go/printer"
)

const (
	govecPreamble = "_govec"
	govecNoExportPreamble = "__govec"
	uniformPreamble = "Uniform"
	programPreamble = "Program"
	rangeFunction = "Range"
	govecFuncPreamble = govecPreamble
	govecstructPreamble = govecPreamble
	govecBuildDirName = "govec_build"
	govecToolName = "govectool"
)

var f *os.File
var h *os.File
var shadow *os.File
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

// Write a function wrapper in shadow
func writeShadowFunctionWrapper(shadow *os.File, funcDecl, funcBody string) {
	s := funcDecl + " {\n" + "\t" + funcBody + "\n}\n";
	_, err := shadow.Write([]byte(s))
	check(err, "Couldn't write to shadow file")
}

// This function modifys the function ast node
func getCFunctionDecl(node *ast.FuncDecl, f *token.FileSet) string {

	var b bytes.Buffer
	var s string

	node.Body = nil
	node.Name.Name = strings.Replace(node.Name.Name, "govec", "", -1)

	goprinter.Fprint(&b, f, node)

	s = b.String()

	// Replacements
	s = strings.Replace(s, "govec.UniformFloat32", "float32", -1)
	s = strings.Replace(s, "govec.UniformInt", "int", -1)
	s = strings.Replace(s, "_govec", "", -1)

	return s
}

func parseFuncDecl(node *ast.FuncDecl) {

	shadowStr, err := Fprint(f, h, fset, node)
	check(err)

	writeShadowFunctionWrapper(shadow, getCFunctionDecl(node, fset),
					shadowStr)
}

func parseFuncDeclNoExport(node *ast.FuncDecl) {

	_, err := Fprint(f, h, fset, node)
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
		if !strings.HasPrefix(node.Name.Name, govecNoExportPreamble) {
			return false
		}

		parseFuncDeclNoExport(node)
		node.Name.Name = node.Name.Name[7:]
		return false

		// Not a function we are interested
	}

	node.Name.Name = node.Name.Name[1:]
	parseFuncDecl(node)

	// We already inspected inside of function
	return false
}

func initShadowFile(shadow *os.File, shadowfilepath, pkg,
			libfilename, headerfilename string) {

	// Write package name
	s := "package " + pkg + "\n\n"

	// Add warning
	s = s + "/* DON'T MODIFY THIS FILE." +
		" CREATED AUTOMATICALLY BY GOVEC TOOL */\n\n\n"

	// Include govec build dir
	s = s + "// #cgo CFLAGS: -I" + govecBuildDirName + "\n"

	// Include library
	s = s + "// #cgo LDFLAGS: " + govecBuildDirName + "/" +
			libfilename + "\n"

	// Include header file
	s = s + "// #include " + "<" + headerfilename + ">\n"

	// Import C
	s = s + "import \"C\"" + "\n\n"

	s = s + "import \"unsafe\"" + "\n\n"

	s = s + "var unused_" + headerfilename[:len(headerfilename) - 2] +
			" unsafe.Pointer\n\n"

	_, err := shadow.Write([]byte(s))
	check(err, "Couldn't write to file:", shadowfilepath)
}

func main() {

	if len(os.Args) < 2 {
		exit("Usage:", govecToolName, "<file.go>")
	}

	fileName := os.Args[1]

	_, err := os.Stat(fileName)
	check(err, "Usage:", govecToolName, "<file.go>")

	// Extra / for incase dir is "."
	dir := filepath.Dir(fileName) + "/"
	file := filepath.Base(fileName)

	filesuffix := file[(len(file)-3):]
	if filesuffix != ".go" {
		exit("Usage:", govecToolName, "<file.go>")
	}

	newdir := dir + "/" + govecBuildDirName + "/"
	_, err = os.Stat(newdir);

	if os.IsNotExist(err) {
	    os.Mkdir(newdir, 0700)
	} else {
		check(err, "Couldn't make govec directory", newdir)
	}

	fileprefix := file[:(len(file)-3)]
	newfilepath := newdir + fileprefix + ".ispc"

	headerfile := fileprefix + ".h"
	headerfilepath := newdir + headerfile

	objfilepath := newdir + fileprefix + ".o"

	libfile := "lib" + fileprefix + ".a"
	libfilepath := newdir + libfile

	shadowfile := "govec" + fileprefix + ".go"
	shadowfilepath := dir + shadowfile

	f, err = os.Create(newfilepath)
	check(err, "Couldn't create file:", newfilepath)
	defer f.Close()

	h, err = os.Create(headerfilepath)
	check(err, "Couldn't create file:", headerfilepath)
	defer h.Close()

	shadow, err = os.Create(shadowfilepath)
	check(err, "Couldn't create file:", shadowfilepath)
	defer shadow.Close()

	fset = token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	check(err)

	initShadowFile(shadow, shadowfilepath, node.Name.Name,
			libfile, headerfile)

	ast.Inspect(node, collectGoVecFunctions)

	// Create the object file and the library
	runCommand("Couldn't compile with ispc",
		   "ispc", "-O3", "--target=sse4-x2",
			"--arch=x86-64", newfilepath, "-o",
			objfilepath)

	runCommand("Couldn't name static library",
			"ar", "rcs", libfilepath, objfilepath)
}
