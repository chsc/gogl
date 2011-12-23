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

func GeneratePackages(packages Packages, functsInfo *FunctionsInfo, typeMap TypeMap) error {
	for packageName, pak := range packages {
		fmt.Printf("Generating %s\n", packageName)
		if err := generatePackage(packageName, pak, functsInfo, typeMap); err != nil {
			return err
		}
	}
	return nil
}

func generatePackage(packageName string, pak *Package, functsInfo *FunctionsInfo, typeMap TypeMap) error {
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
	return writePackage(w, packageName, pak, functsInfo, typeMap)
}

func writePackage(w io.Writer, packageName string, pak *Package, functsInfo *FunctionsInfo, typeMap TypeMap) error {
	fmt.Fprintf(w, "// Automatically generated OpenGL binding.\n// \n")
	fmt.Fprintf(w, "// Categories in this package: \n// \n")
	for cat, _ := range pak.Functions {
		pc, _ := ParseCategoryString(cat)
		switch pc.Type {
		case CategoryExtension:
			fmt.Fprintf(w, "// %s: %s\n// \n", cat, MakeExtensionSpecDocUrl(pc.Vendor, pc.Name))
		case CategoryVersion:
			fmt.Fprintf(w, "// %s\n// \n", cat)
		case CategoryDepVersion:
			fmt.Fprintf(w, "// %s\n// \n", cat)
		}
	}
	fmt.Fprintf(w, "package %s\n\n", packageName)

	fmt.Fprintf(w, "// #cgo darwin  LDFLAGS: -framework OpenGL\n")
	fmt.Fprintf(w, "// #cgo linux   LDFLAGS: -lGL\n")
	fmt.Fprintf(w, "// #cgo windows LDFLAGS: -lopengl32\n// \n")

	fmt.Fprintf(w, "// #ifdef __APPLE__\n")
	fmt.Fprintf(w, "// // TODO: add context header\n")
	fmt.Fprintf(w, "// #elif defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)\n")
	fmt.Fprintf(w, "// #define WIN32_LEAN_AND_MEAN 1\n")
	fmt.Fprintf(w, "// #include <windows.h> // for wglGetProcAddress\n")
	fmt.Fprintf(w, "// #else\n")
	fmt.Fprintf(w, "// #include <X11/Xlib.h> // for glXGetProcAddress\n")
	fmt.Fprintf(w, "// #include <GL/glx.h> \n")
	fmt.Fprintf(w, "// #endif\n// \n")

	fmt.Fprintf(w, "// #ifndef APIENTRY\n")
	fmt.Fprintf(w, "// #define APIENTRY\n")
	fmt.Fprintf(w, "// #endif\n")
	fmt.Fprintf(w, "// #ifndef APIENTRYP\n")
	fmt.Fprintf(w, "// #define APIENTRYP APIENTRY *\n")
	fmt.Fprintf(w, "// #endif\n")
	fmt.Fprintf(w, "// #ifndef GLAPI\n")
	fmt.Fprintf(w, "// #define GLAPI extern\n")
	fmt.Fprintf(w, "// #endif\n// \n")

	fmt.Fprintf(w, "// typedef unsigned int GLenum;\n")
	fmt.Fprintf(w, "// typedef unsigned char GLboolean;\n")
	fmt.Fprintf(w, "// typedef unsigned int GLbitfield;\n")
	fmt.Fprintf(w, "// typedef signed char GLbyte;\n")
	fmt.Fprintf(w, "// typedef short GLshort;\n")
	fmt.Fprintf(w, "// typedef int GLint;\n")
	fmt.Fprintf(w, "// typedef int GLsizei;\n")
	fmt.Fprintf(w, "// typedef unsigned char GLubyte;\n")
	fmt.Fprintf(w, "// typedef unsigned short GLushort;\n")
	fmt.Fprintf(w, "// typedef unsigned int GLuint;\n")
	fmt.Fprintf(w, "// typedef unsigned short GLhalf;\n")
	fmt.Fprintf(w, "// typedef float GLfloat;\n")
	fmt.Fprintf(w, "// typedef float GLclampf;\n")
	fmt.Fprintf(w, "// typedef double GLdouble;\n")
	fmt.Fprintf(w, "// typedef double GLclampd;\n")
	fmt.Fprintf(w, "// typedef void GLvoid;\n// \n")

	for _, pt := range functsInfo.Passthru {
		fmt.Fprintf(w, "// %s\n", pt)
	}

	fmt.Fprintf(w, "// void* goglGetProcAddress(const char* name) { \n")
	fmt.Fprintf(w, "// #ifdef __APPLE__\n")
	fmt.Fprintf(w, "// 	return NULL; // TODO: Add get proc addr for Mac.\n")
	fmt.Fprintf(w, "// #elif _WIN32\n")
	fmt.Fprintf(w, "// 	return wglGetProcAddress((LPCSTR)name); // TODO: Not tested\n")
	fmt.Fprintf(w, "// #else\n")
	fmt.Fprintf(w, "// 	return glXGetProcAddress((const GLubyte*)name);\n")
	fmt.Fprintf(w, "// #endif\n")
	fmt.Fprintf(w, "// }\n// \n")

	if err := writeCFuncDefs(w, pak.Functions, typeMap); err != nil {
		return err
	}

	if err := writeCFuncDecls(w, pak.Functions, typeMap); err != nil {
		return err
	}

	if err := writeCFuncGetprocAddrs(w, pak.Functions); err != nil {
		return err
	}

	fmt.Fprintf(w, "import \"C\"\n")
	fmt.Fprintf(w, "import \"unsafe\"\n\n")

	fmt.Fprintf(w, "type (\n")
	fmt.Fprintf(w, "	Enum     C.GLenum\n")
	fmt.Fprintf(w, "	Boolean  C.GLboolean\n")
	fmt.Fprintf(w, "	Bitfield C.GLBitfield\n")
	fmt.Fprintf(w, "	Byte     C.GLbyte\n")
	fmt.Fprintf(w, "	Short    C.GLshort\n")
	fmt.Fprintf(w, "	Int      C.GLint\n")
	fmt.Fprintf(w, "	Sizei    C.GLsizei\n")
	fmt.Fprintf(w, "	Ubyte    C.GLubyte\n")
	fmt.Fprintf(w, "	Ushort   C.GLushort\n")
	fmt.Fprintf(w, "	Uint     C.GLuint\n")
	fmt.Fprintf(w, "	Half     C.GLhalf\n")
	fmt.Fprintf(w, "	Float    C.GLfloat\n")
	fmt.Fprintf(w, "	Clampf   C.GLclampf\n")
	fmt.Fprintf(w, "	Double   C.GLdouble\n")
	fmt.Fprintf(w, "	Clampd   C.GLclampd\n")
	fmt.Fprintf(w, "	Char     C.GLchar\n")
	fmt.Fprintf(w, "	Pointer  unsafe.Pointer\n")
	fmt.Fprintf(w, ")\n\n")

	writeGoEnumDefinitions(w, pak.Enums)

	if err := writeGoFunctionDefinitions(w, pak.Functions, typeMap); err != nil {
		return err
	}

	fmt.Fprintf(w, "// EOF")

	return nil
}

func writeCFuncDefs(w io.Writer, functions FunctionCategories, typeMap TypeMap) error {
	for cat, fs := range functions {
		fmt.Fprintf(w, "// //  %s\n", cat)
		for _, f := range fs {
			rtype, err := typeMap.Resolve(f.Return)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "// %s (APIENTRYP ptrgogl%s)(", rtype, f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				ptype, err := typeMap.Resolve(f.Parameters[p].Type)
				if err != nil {
					return err
				}
				fmt.Fprintf(w, "%s ", ptype)
				if f.Parameters[p].Out && !f.Parameters[p].Array {
					fmt.Fprint(w, "*")
				}
				if f.Parameters[p].Array {
					fmt.Fprintf(w, " *")
				}
				fmt.Fprintf(w, "%s", f.Parameters[p].Name)
				if p < len(f.Parameters)-1 {
					fmt.Fprintf(w, ", ")
				}
			}
			fmt.Fprintf(w, ");\n")
		}
	}
	fmt.Fprintf(w, "// \n")
	return nil
}

func writeCFuncDecls(w io.Writer, functions FunctionCategories, typeMap TypeMap) error {
	for cat, fs := range functions {
		fmt.Fprintf(w, "// //  %s\n", cat)
		for _, f := range fs {
			rtype, err := typeMap.Resolve(f.Return)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "// %s gogl%s(", rtype, f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				ptype, err := typeMap.Resolve(f.Parameters[p].Type)
				if err != nil {
					return err
				}
				fmt.Fprintf(w, "%s ", ptype)
				if f.Parameters[p].Out && !f.Parameters[p].Array {
					fmt.Fprint(w, "*")
				}
				if f.Parameters[p].Array {
					fmt.Fprintf(w, "*")
				}
				fmt.Fprintf(w, "%s", f.Parameters[p].Name)
				if p < len(f.Parameters)-1 {
					fmt.Fprintf(w, ", ")
				}
			}
			fmt.Fprintf(w, ") {\n")
			if rtype == "void" {
				fmt.Fprintf(w, "// 	(*ptrgogl%s)(", f.Name)
			} else {
				fmt.Fprintf(w, "// 	return (*ptrgogl%s)(", f.Name)
			}
			for p := 0; p < len(f.Parameters); p++ {
				fmt.Fprintf(w, "%s", f.Parameters[p].Name)
				if p < len(f.Parameters)-1 {
					fmt.Fprintf(w, ", ")
				}
			}
			fmt.Fprintf(w, ");\n// }\n")
		}
	}
	fmt.Fprintf(w, "// \n")
	return nil
}

func writeCFuncGetprocAddrs(w io.Writer, functions FunctionCategories) error {
	for cat, fs := range functions {
		fmt.Fprintf(w, "// int init_%s() {\n", cat)
		for _, f := range fs {
			fmt.Fprintf(w, "// 	ptrgogl%s = goglGetProcAddress(\"gl%s\");\n", f.Name, f.Name)
			fmt.Fprintf(w, "// 	if(ptrgogl%s == NULL) return 1;\n", f.Name)
		}
		fmt.Fprintf(w, "// \treturn 0;\n")
		fmt.Fprintf(w, "// }\n")
	}
	fmt.Fprintf(w, "// \n")
	return nil
}

func writeGoEnumDefinitions(w io.Writer, enums EnumCategories) {
	for cat, enums := range enums {
		fmt.Fprintf(w, "// %s\n", cat)
		fmt.Fprintf(w, "const (\n")
		for _, e := range enums {
			fmt.Fprintf(w, "	%s = %s\n", e.Name, e.Value)
		}
		fmt.Fprintf(w, ")\n")
	}
}

func writeGoFunctionDefinitions(w io.Writer, functions FunctionCategories, typeMap TypeMap) error {
	for cat, fs := range functions {
		fmt.Fprintf(w, "// %s\n", cat)
		for _, f := range fs {
			fmt.Fprintf(w, "func %s(", f.Name)
			for p := 0; p < len(f.Parameters); p++ {
				fmt.Fprintf(w, "%s ", RenameReservedWord(f.Parameters[p].Name))
				ptype, err := typeMap.Resolve(f.Parameters[p].Type)
				if err != nil {
					return err
				}
				goptype, err := GLTypeToGoType(ptype, f.Parameters[p].Out, f.Parameters[p].Array)
				if err != nil {
					return err
				}
				if f.Parameters[p].Out && !f.Parameters[p].Array {
					fmt.Fprint(w, "*")
				}
				if f.Parameters[p].Array {
					fmt.Fprintf(w, "*")
				}
				fmt.Fprintf(w, "%s", goptype)
				if p < len(f.Parameters)-1 {
					fmt.Fprintf(w, ", ")
				}
			}
			rtype, err := typeMap.Resolve(f.Return)
			if err != nil {
				return err
			}
			gortype, err := GLTypeToGoType(rtype, false, false)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, ") %s {\n", gortype)
			if rtype == "void" {
				fmt.Fprintf(w, "	C.gogl%s(", f.Name)
			} else {
				fmt.Fprintf(w, "	return C.gogl%s(", f.Name)
			}
			for p := 0; p < len(f.Parameters); p++ {
				ptype, err := typeMap.Resolve(f.Parameters[p].Type)
				if err != nil {
					return err
				}
				fmt.Fprintf(w, "C.%s(%s)", ptype, RenameReservedWord(f.Parameters[p].Name))
				if p < len(f.Parameters)-1 {
					fmt.Fprintf(w, ", ")
				}
			}
			fmt.Fprintf(w, ")\n}\n")
		}
	}
	return nil
}
