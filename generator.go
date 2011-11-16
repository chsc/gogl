package main

import (
	"os"
	"io"
	"fmt"
)

func Generate(fname string, enums EnumCategories, functions FunctionCategories, typeMap TypeMap) os.Error {
	w, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer w.Close()

	fmt.Fprintf(w, "package gl\n")

	fmt.Fprintf(w, "// #cgo darwin  LDFLAGS: -framework OpenGL\n")
	fmt.Fprintf(w, "// #cgo linux   LDFLAGS: -lGL\n")
	fmt.Fprintf(w, "// #cgo windows LDFLAGS: -lopengl32\n")
	fmt.Fprintf(w, "// #ifdef __APPLE__\n")
	fmt.Fprintf(w, "// # include <OpenGL/gl.h>\n")
	fmt.Fprintf(w, "// #else\n")
	fmt.Fprintf(w, "// # include <GL/gl.h>\n")
	fmt.Fprintf(w, "// #endif\n")
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
