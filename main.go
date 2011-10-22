// OpenGL Binding generator

package main

import (
	"fmt"
	"flag"
)

var (
	download    *bool   = flag.Bool("download", false, "Download spec files from Khronos registry first.")
	version     *string = flag.String("version", "2.1", "OpenGL version")
	deprecation *bool   = flag.Bool("deprecate", false, "Exclude all deprecated features.")
	outGLFile   *string = flag.String("outgl", "gl.go", "Output file name")
	// TODO: add flags for additional extensions
)

func enumFilter(category string, enum *Enum) bool {
	// TODO: filter by version/extension use flags
	return true
}

func functionFilter(category string, function *Function) bool {
	// TODO: filter by version/extension use flags
	return true
}

func main() {
	fmt.Printf("OpenGL binding generator for Go. Copyright (c) 2011 by Christoph Schunk\n")
	flag.Parse()

	if *download {
		DownloadOpenGLSpecs()
	}

	fmt.Printf("Parsing enumext.spec file...\n")
	enums, err := ReadEnumsFromFile(OpenGLEnumExtSpecFile)
	if err != nil {
		panic(err.String())
	}

	fmt.Printf("Parsing type map file ...\n")
	typeMap, err := ReadTypeMapFromFile(OpenGLTypeMapFile)
	if err != nil {
		panic(err.String())
	}

	fmt.Printf("Parsing gl.spec file ...\n")
	functions, err := ReadFunctionsFromFile(OpenGLSpecFile)
	if err != nil {
		panic(err.String())
	}

	// TODO: just a test, do real unit testing
	functions["Cat1"] = []Function{
		Function{Name: "Foo1", Parameters: []Parameter{Parameter{"p1", "int"}, Parameter{"p2", "int"}}, Return: "void"},
		Function{Name: "Foo2", Parameters: []Parameter{Parameter{"p1", "int"}, Parameter{"p2", "int"}, Parameter{"p3", "float"}}, Return: "void"},
		Function{Name: "Foo3", Parameters: []Parameter{Parameter{"p1", "int"}, Parameter{"p2", "int"}, Parameter{"p3", "float"}}, Return: "void"},
	}

	Generate(*outGLFile, enums, functions, typeMap, enumFilter, functionFilter)
}
