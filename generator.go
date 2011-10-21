package main

import (
	"os"
	"io"
	"fmt"
)

type FilterEnum     func(category string, enum *Enum) bool
type FilterFunction func(category string, function *Function) bool

func Generate(fname string, enums Enums, functions Functions, typeMap TypeMap, fenum FilterEnum, ffunctions FilterFunction) os.Error {
	w, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer w.Close()

	fmt.Fprintf(w, "// OpenGL\n")
	fmt.Fprintf(w, "package gl\n")
	
	// TODO: writeCFuncPtrDefinitions
	writeCFunctionDefinitions(functions, ffunctions, w)
	fmt.Fprintf(w, "import \"C\"\n\n")

	writeGoEnumDefinitions(enums, fenum, w)

	writeGoFunctionDefinitions(functions, ffunctions, w)

	fmt.Fprintf(w, "// EOF\n")
	return nil
}

func writeGoEnumDefinitions(enums Enums, fenum FilterEnum, w io.Writer) {
	fmt.Fprintf(w, "const (\n")
	for name, enums := range(enums) {
		if fenum(name, nil) {
			fmt.Fprintf(w, "\t// %s:\n", name)
			for _, e := range(enums) {
				fmt.Fprintf(w, "\t%s = %s\n", e.Name, e.Value)
			}
		}
	}
	fmt.Fprintf(w, ")\n")
}

func writeGoFunctionDefinitions(functions Functions, ffunctions FilterFunction, w io.Writer) {
	for cat, fs := range(functions) {
		fmt.Fprintf(w, "// %s\n", cat)
		for _, f := range(fs) {
			fmt.Fprintf(w, "func %s(", f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters) - 1 {
					fmt.Fprintf(w, "%s %s", f.Parameters[p].Name, f.Parameters[p].Type)
				} else {
					fmt.Fprintf(w, "%s %s, ", f.Parameters[p].Name, f.Parameters[p].Type)
				}
			}
			fmt.Fprintf(w, ") {\n")
			fmt.Fprintf(w, "\tC.gogl%s(", f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters) - 1 {
					fmt.Fprintf(w, "%s", f.Parameters[p].Name)
				} else {
					fmt.Fprintf(w, "%s, ", f.Parameters[p].Name)
				}
			}
			fmt.Fprintf(w, ")\n}\n")
		}
	}
}

func writeCFunctionDefinitions(functions Functions, ffunctions FilterFunction, w io.Writer) {
	for cat, fs := range(functions) {
		fmt.Fprintf(w, "// /* %s */  \n", cat)
		for _, f := range(fs) {
			fmt.Fprintf(w, "// %s gogl%s(", f.Return, f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters) - 1 {
					fmt.Fprintf(w, "%s %s", f.Parameters[p].Type, f.Parameters[p].Name)
				} else {
					fmt.Fprintf(w, "%s %s, ", f.Parameters[p].Type, f.Parameters[p].Name)
				}
			}
			fmt.Fprintf(w, ") {\n")
			fmt.Fprintf(w, "// \timpptr_gogl%s(", f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				if p == len(f.Parameters) - 1 {
					fmt.Fprintf(w, "%s", f.Parameters[p].Name)
				} else {
					fmt.Fprintf(w, "%s, ", f.Parameters[p].Name)
				}
			}
			fmt.Fprintf(w, ")\n// }\n")
		}
	}
}

/*
TODO: write function pointers & load them via getProcAddress ...

func (f *Function) WriteGetProcAddr(w io.Writer) {
	// TODO: Windows, Linux, ...
	fmt.Fprintf(w, "//\timpptr_%s = glxGetProcAddress(\"gl%s\");\n", f.Name, f.Name)
}

func (fg FunctionCategory) WriteGetProcAddrs(w io.Writer) {
	fmt.Fprintf(w, "// int init_%s_extensions() {\n", fg.Name)
	for _, f := range(fg.Functions) {
		f.WriteGetProcAddr(w)
	}
	fmt.Fprintf(w, "//\treturn 1;\n")
	fmt.Fprintf(w, "// }\n")
}*/

