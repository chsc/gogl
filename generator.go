// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func MakeSpecDocUrl(vendor, extension string) string {
	return fmt.Sprintf("http://www.opengl.org/registry/specs/%s/%s.txt", vendor, extension)
}

func GeneratePackages(packages Packages, typeMap TypeMap) error {
	for packageName, pak := range packages {
		//fmt.Printf("Generating %s ...\n",  packageName);
		if err := generatePackage(packageName, pak, typeMap); err != nil {
			return err
		}
	}
	return nil
}

func generatePackage(packageName string, pak *Package, typeMap TypeMap) error {
	absPath, err := filepath.Abs(packageName)
	if err != nil {
		return err
	}
	err = os.MkdirAll(absPath, 0755)
	if err != nil {
		return err
	}
	w, err := os.Create(filepath.Join(absPath, fmt.Sprintf("%s.go", packageName)))
	if err != nil {
		return err
	}
	defer w.Close()
	return writePackage(w, packageName, pak, typeMap);
}

func writePackage(w io.Writer, packageName string, pak *Package, typeMap TypeMap) error {
	fmt.Fprintf(w, "// OpenGL.\n")
	fmt.Fprintf(w, "// \n")
	fmt.Fprintf(w, "//  %s \n", MakeSpecDocUrl("ARB", "multisample"))	
	fmt.Fprintf(w, "package %s\n", packageName)

	fmt.Fprintf(w, "\n/*\n")
	fmt.Fprintf(w, "#cgo darwin  LDFLAGS: -framework OpenGL\n")
	fmt.Fprintf(w, "#cgo linux   LDFLAGS: -lGL\n")
	fmt.Fprintf(w, "#cgo windows LDFLAGS: -lopengl32\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "#if defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)\n")
	fmt.Fprintf(w, "#define WIN32_LEAN_AND_MEAN 1\n")
	fmt.Fprintf(w, "#include <windows.h>\n")
	fmt.Fprintf(w, "#endif\n")
	fmt.Fprintf(w, "#ifndef APIENTRY\n")
	fmt.Fprintf(w, "#define APIENTRY\n")
	fmt.Fprintf(w, "#endif\n")
	fmt.Fprintf(w, "#ifndef APIENTRYP\n")
	fmt.Fprintf(w, "#define APIENTRYP APIENTRY *\n")
	fmt.Fprintf(w, "#endif\n")
	fmt.Fprintf(w, "#ifndef GLAPI\n")
	fmt.Fprintf(w, "#define GLAPI extern\n")
	fmt.Fprintf(w, "#endif\n\n")
	
	fmt.Fprintf(w, "typedef unsigned int GLenum;\n")
	fmt.Fprintf(w, "typedef unsigned char GLboolean;\n")
	fmt.Fprintf(w, "typedef unsigned int GLbitfield;\n")
	fmt.Fprintf(w, "typedef signed char GLbyte;\n")
	fmt.Fprintf(w, "typedef short GLshort;\n")
	fmt.Fprintf(w, "typedef int GLint;\n")
	fmt.Fprintf(w, "typedef int GLsizei;\n")
	fmt.Fprintf(w, "typedef unsigned char GLubyte;\n")
	fmt.Fprintf(w, "typedef unsigned short GLushort;\n")
	fmt.Fprintf(w, "typedef unsigned int GLuint;\n")
	fmt.Fprintf(w, "typedef unsigned short GLhalf;\n")
	fmt.Fprintf(w, "typedef float GLfloat;\n")
	fmt.Fprintf(w, "typedef float GLclampf;\n")
	fmt.Fprintf(w, "typedef double GLdouble;\n")
	fmt.Fprintf(w, "typedef double GLclampd;\n")
	fmt.Fprintf(w, "typedef void GLvoid;\n\n")

	//write

	fmt.Fprintf(w, "*/\n")
	fmt.Fprintf(w, "import \"C\"\n")

	// write passthru

	return nil
}

func generate(fname string, enums EnumCategories, functions FunctionCategories, typeMap TypeMap) error {
	w, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer w.Close()

	writeCFuncPtrDefinitions(functions, w)
	writeCFunctionDeclarations(functions, w)
	writeGetProcAddrsDeclarations(functions, w)
	fmt.Fprintf(w, "import \"C\"\n\n")

	//writeGoEnumDefinitions(enums, w)
	//writeGoFunctionDefinitions(functions, w)

	fmt.Fprintf(w, "// EOF\n")
	return nil
}

func writeGoEnumDefinitions(enums EnumCategories, w io.Writer) {
	fmt.Fprintf(w, "const (\n")
	for name, enums := range enums {
		fmt.Fprintf(w, "\t// %s:\n", name)
		for _, e := range enums {
			fmt.Fprintf(w, "\t%s = %s\n", e.Name, e.Value)
		}
	}
	fmt.Fprintf(w, ")\n")
}

func writeGoFunctionDefinitions(functions FunctionCategories, w io.Writer) {
	for cat, fs := range functions {
		fmt.Fprintf(w, "// %s\n", cat)
		for _, f := range fs {
			fmt.Fprintf(w, "func %s(", f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters)-1 {
					fmt.Fprintf(w, "%s %s", f.Parameters[p].Name, f.Parameters[p].Type)
				} else {
					fmt.Fprintf(w, "%s %s, ", f.Parameters[p].Name, f.Parameters[p].Type)
				}
			}
			fmt.Fprintf(w, ") {\n")
			fmt.Fprintf(w, "\tC.gogl%s(", f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters)-1 {
					fmt.Fprintf(w, "%s", f.Parameters[p].Name)
				} else {
					fmt.Fprintf(w, "%s, ", f.Parameters[p].Name)
				}
			}
			fmt.Fprintf(w, ")\n}\n")
		}
	}
}

func writeCFuncPtrDefinitions(functions FunctionCategories, w io.Writer) {
	fmt.Fprintf(w, "// /* function pointers: */\n")
	for cat, fs := range functions {
		fmt.Fprintf(w, "// /* %s */  \n", cat)
		for _, f := range fs {
			fmt.Fprintf(w, "// %s (*ptrgogl%s)(", f.Return, f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters)-1 {
					fmt.Fprintf(w, "%s %s", f.Parameters[p].Type, f.Parameters[p].Name)
				} else {
					fmt.Fprintf(w, "%s %s, ", f.Parameters[p].Type, f.Parameters[p].Name)
				}
			}
			fmt.Fprintf(w, ");\n")
		}
	}
	fmt.Fprintf(w, "// \n")
}

func writeCFunctionDeclarations(functions FunctionCategories, w io.Writer) {
	fmt.Fprintf(w, "// /* function declarations */\n")
	for cat, fs := range functions {
		fmt.Fprintf(w, "// /* %s */  \n", cat)
		for _, f := range fs {
			fmt.Fprintf(w, "// %s gogl%s(", f.Return, f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters)-1 {
					fmt.Fprintf(w, "%s %s", f.Parameters[p].Type, f.Parameters[p].Name)
				} else {
					fmt.Fprintf(w, "%s %s, ", f.Parameters[p].Type, f.Parameters[p].Name)
				}
			}
			fmt.Fprintf(w, ") {\n")
			fmt.Fprintf(w, "// \tptrgogl%s(", f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters)-1 {
					fmt.Fprintf(w, "%s", f.Parameters[p].Name)
				} else {
					fmt.Fprintf(w, "%s, ", f.Parameters[p].Name)
				}
			}
			fmt.Fprintf(w, ")\n// }\n")
		}
	}
}

func writeGetProcAddrsDeclarations(functions FunctionCategories, w io.Writer) {
	fmt.Fprintf(w, "// /* extension loading */\n")
	fmt.Fprintf(w, "// void* goglGetProcAddress(const char* name) { \n")
	fmt.Fprintf(w, "// #ifdef __APPLE__\n")
	fmt.Fprintf(w, "// \treturn NULL;\n")
	fmt.Fprintf(w, "// #elif WIN32\n")
	fmt.Fprintf(w, "// \treturn wglGetProcAddress((LPCSTR)name);\n")
	fmt.Fprintf(w, "// #else\n")
	fmt.Fprintf(w, "// \treturn glXGetProcddress((const GLubyte*)name);\n")
	fmt.Fprintf(w, "// #endif\n")
	fmt.Fprintf(w, "// }\n")
	for cat, fs := range functions {
		fmt.Fprintf(w, "// int init_%s() {\n", cat)
		for _, f := range fs {
			// TODO: Windows, Linux, ...
			fmt.Fprintf(w, "//\tptrgogl%s = goglGetProcAddress(\"gl%s\");\n", f.Name, f.Name)
		}
		fmt.Fprintf(w, "//\treturn 1;\n")
		fmt.Fprintf(w, "// }\n")
	}
}
