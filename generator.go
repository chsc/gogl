// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GeneratePackages(packages Packages, functsInfo *FunctionsInfo, typeMap TypeMap) error {
	for packageName, pak := range packages {
		fmt.Printf("Generating package %s ...\n", packageName)
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
	highestMajorVersion := -1
	isGLPackage := false

	if strings.HasPrefix(packageName, "gl") && !strings.HasPrefix(packageName, "glx") {
		isGLPackage = true
	}

	fmt.Fprintf(w, "// Automatically generated OpenGL binding.\n// \n")
	fmt.Fprintf(w, "// Categories in this package: \n// \n")

	sortedFunctionCategories := make([]string, len(pak.Functions))
	i := 0
	for k, _ := range pak.Functions {
		sortedFunctionCategories[i] = k
		i++
	}
	sort.Strings(sortedFunctionCategories)

	sortedEnumCategories := make([]string, len(pak.Enums))
	i = 0
	for k, _ := range pak.Enums {
		sortedEnumCategories[i] = k
		i++
	}
	sort.Strings(sortedEnumCategories)

	for _, cat := range sortedFunctionCategories {
		pc, _ := ParseCategoryString(cat)
		switch pc.Type {
		case CategoryExtension:
			fmt.Fprintf(w, "// %s: %s\n// \n", cat, MakeExtensionSpecDocUrl(pc.Vendor, pc.Name))
		case CategoryVersion:
			fmt.Fprintf(w, "// %s\n// \n", cat)
			if pc.Version.Major > highestMajorVersion {
				highestMajorVersion = pc.Version.Major
			}
		case CategoryDepVersion:
			fmt.Fprintf(w, "// %s\n// \n", cat)
		}
	}
	if highestMajorVersion > 0 {
		fmt.Fprintf(w, "// %s\n// \n", MakeGLDocUrl(highestMajorVersion))
	}
	fmt.Fprintf(w, "package %s\n\n", packageName)

	fmt.Fprintf(w, "// #cgo darwin  LDFLAGS: -framework OpenGL\n")
	fmt.Fprintf(w, "// #cgo linux   LDFLAGS: -lGL\n")
	fmt.Fprintf(w, "// #cgo windows LDFLAGS: -lopengl32\n// \n")

	fmt.Fprintf(w, "// #include <stdlib.h>\n")
	fmt.Fprintf(w, "// #if defined(__APPLE__)\n")
	fmt.Fprintf(w, "// #include <dlfcn.h>\n")
	fmt.Fprintf(w, "// #elif defined(_WIN32)\n")
	fmt.Fprintf(w, "// #define WIN32_LEAN_AND_MEAN 1\n")
	fmt.Fprintf(w, "// #include <windows.h>\n")
	fmt.Fprintf(w, "// #else\n")
	fmt.Fprintf(w, "// #include <X11/Xlib.h>\n")
	fmt.Fprintf(w, "// #include <GL/glx.h>\n")
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

	for _, passthrus := range functsInfo.Passthru["global"] {
		fmt.Fprintf(w, "// %s\n", passthrus)
	}

	fmt.Fprintf(w, "// #ifdef _WIN32\n")
	fmt.Fprintf(w, "// static HMODULE opengl32 = NULL;\n")
	fmt.Fprintf(w, "// #endif\n// \n")

	fmt.Fprintf(w, "// static void* goglGetProcAddress(const char* name) { \n")
	fmt.Fprintf(w, "// #ifdef __APPLE__\n")
	fmt.Fprintf(w, "// 	return dlsym(RTLD_DEFAULT, name);\n")
	fmt.Fprintf(w, "// #elif _WIN32\n")
	fmt.Fprintf(w, "// 	void* pf = wglGetProcAddress((LPCSTR)name);\n")
	fmt.Fprintf(w, "// 	if(pf) {\n")
	fmt.Fprintf(w, "// 		return pf;\n")
	fmt.Fprintf(w, "// 	}\n")
	fmt.Fprintf(w, "// 	if(opengl32 == NULL) {\n")
	fmt.Fprintf(w, "// 		opengl32 = LoadLibraryA(\"opengl32.dll\");\n")
	fmt.Fprintf(w, "// 	}\n")
	fmt.Fprintf(w, "// 	return GetProcAddress(opengl32, (LPCSTR)name);\n")
	fmt.Fprintf(w, "// #else\n")
	fmt.Fprintf(w, "// 	return glXGetProcAddress((const GLubyte*)name);\n")
	fmt.Fprintf(w, "// #endif\n")
	fmt.Fprintf(w, "// }\n// \n")

	if err := writeCFuncDeclarations(w, sortedFunctionCategories, pak.Functions, typeMap); err != nil {
		return err
	}

	if err := writeCFuncDefinitions(w, sortedFunctionCategories, pak.Functions, typeMap); err != nil {
		return err
	}

	if err := writeCFuncGetProcAddrs(w, sortedFunctionCategories, pak.Functions); err != nil {
		return err
	}

	fmt.Fprintf(w, "import \"C\"\n")
	fmt.Fprintf(w, "import \"unsafe\"\n")
	fmt.Fprintf(w, "import \"errors\"\n\n")

	fmt.Fprintf(w, "type (\n")
	fmt.Fprintf(w, "	Enum     C.GLenum\n")
	fmt.Fprintf(w, "	Boolean  C.GLboolean\n")
	fmt.Fprintf(w, "	Bitfield C.GLbitfield\n")
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
	fmt.Fprintf(w, "	Sync     C.GLsync\n")
	fmt.Fprintf(w, "	Int64    C.GLint64\n")
	fmt.Fprintf(w, "	Uint64   C.GLuint64\n")
	fmt.Fprintf(w, "	Intptr   C.GLintptr\n")
	fmt.Fprintf(w, "	Sizeiptr C.GLsizeiptr\n")
	fmt.Fprintf(w, ")\n\n")

	writeGoEnumDefinitions(w, sortedEnumCategories, pak.Enums)

	if err := writeGoFuncDefinitions(w, sortedFunctionCategories, pak.Functions, typeMap, highestMajorVersion); err != nil {
		return err
	}

	writeGoInitDefinitions(w, sortedFunctionCategories, pak.Functions, isGLPackage)

	if isGLPackage {
		writeUtilityFunctions(w)
	}

	fmt.Fprintf(w, "// EOF")

	return nil
}

func writeCFuncDeclarations(w io.Writer, sortedFunctionCategories []string, functions FunctionCategories, typeMap TypeMap) error {
	for _, cat := range sortedFunctionCategories {
		fs := functions[cat]
		fmt.Fprintf(w, "// //  %s\n", cat)
		for _, f := range fs {
			if err := writeCFuncDeclaration(w, f, typeMap, false); err != nil {
				return err
			}
		}
	}
	fmt.Fprintf(w, "// \n")
	return nil
}

func writeCFuncDeclaration(w io.Writer, f *Function, typeMap TypeMap, prototype bool) error {
	rtype, err := typeMap.Resolve(f.Return)
	if err != nil {
		return err
	}
	if prototype {
		fmt.Fprintf(w, "// GLAPI %s APIENTRY gl%s(", rtype, f.Name)
	} else {
		fmt.Fprintf(w, "// %s (APIENTRYP ptrgl%s)(", rtype, f.Name)
	}
	for i, _ := range f.Parameters {
		p := &f.Parameters[i]
		ptype, err := typeMap.Resolve(p.Type)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s", ptype)
		if p.Out || (p.Modifier == ParamModifierReference || p.Modifier == ParamModifierArray) {
			fmt.Fprint(w, "*")
		}
		fmt.Fprintf(w, " %s", p.Name)
		if i < len(f.Parameters)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ");\n")
	return nil
}

func writeCFuncDefinitions(w io.Writer, sortedFunctionCategories []string, functions FunctionCategories, typeMap TypeMap) error {
	for _, cat := range sortedFunctionCategories {
		fs := functions[cat]
		fmt.Fprintf(w, "// //  %s\n", cat)
		for _, f := range fs {
			if err := writeCFuncDefinition(w, f, typeMap, false); err != nil {
				return err
			}
		}
	}
	fmt.Fprintf(w, "// \n")
	return nil
}

func writeCFuncDefinition(w io.Writer, f *Function, typeMap TypeMap, prototype bool) error {
	rtype, err := typeMap.Resolve(f.Return)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "// %s gogl%s(", rtype, f.Name)
	for i, _ := range f.Parameters {
		p := &f.Parameters[i]
		ptype, err := typeMap.Resolve(p.Type)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s", ptype)
		if p.Out || (p.Modifier == ParamModifierReference || p.Modifier == ParamModifierArray) {
			fmt.Fprint(w, "*")
		}
		fmt.Fprintf(w, " %s", RenameIfReservedWord(p.Name))
		if i < len(f.Parameters)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ") {\n")
	if rtype == "void" {
		fmt.Fprintf(w, "// 	")
	} else {
		fmt.Fprintf(w, "// 	return ")
	}
	if prototype {
		fmt.Fprintf(w, "gl%s(", f.Name)
	} else {
		fmt.Fprintf(w, "(*ptrgl%s)(", f.Name)
	}
	for i, _ := range f.Parameters {
		p := &f.Parameters[i]
		fmt.Fprintf(w, "%s", RenameIfReservedWord(p.Name))
		if i < len(f.Parameters)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ");\n// }\n")
	return nil
}

func writeCFuncGetProcAddrs(w io.Writer, sortedFunctionCategories []string, functions FunctionCategories) error {
	for _, cat := range sortedFunctionCategories {
		fs := functions[cat]
		fmt.Fprintf(w, "// int init_%s() {\n", cat)
		for _, f := range fs {
			fmt.Fprintf(w, "// 	ptrgl%s = goglGetProcAddress(\"gl%s\");\n", f.Name, f.Name)
			fmt.Fprintf(w, "// 	if(ptrgl%s == NULL) return 1;\n", f.Name)
		}
		fmt.Fprintf(w, "// \treturn 0;\n")
		fmt.Fprintf(w, "// }\n")
	}
	fmt.Fprintf(w, "// \n")
	return nil
}

func writeGoEnumDefinitions(w io.Writer, sortedEnumCategories []string, enumCats EnumCategories) {
	for _, cat := range sortedEnumCategories {
		enums := enumCats[cat]
		fmt.Fprintf(w, "// %s\n", cat)
		fmt.Fprintf(w, "const (\n")

		sortedEnums := make([]string, len(enums))
		i := 0
		for k, _ := range enums {
			//if strings.HasPrefix(k, "COPY_READ") {
			//	fmt.Println(k)
			//}
			sortedEnums[i] = k
			i++
		}
		sort.Strings(sortedEnums)

		for _, e := range sortedEnums {
			v := enums[e]
			if !enumCats.IsAlreadyDefined(e, cat) {
				fmt.Fprintf(w, "\t%s = %s\n", CleanEnumName(e), v)
			}
		}
		fmt.Fprintf(w, ")\n")
	}
}

func writeGoFuncDefinitions(w io.Writer, sortedFunctionCategories []string, functions FunctionCategories, typeMap TypeMap, majorVersion int) error {
	for _, cat := range sortedFunctionCategories {
		fs := functions[cat]
		fmt.Fprintf(w, "// %s\n\n", cat)
		for _, f := range fs {
			if err := writeGoFuncDefinition(w, f, typeMap, majorVersion, false); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeGoFuncDefinition(w io.Writer, f *Function, typeMap TypeMap, majorVersion int, prototype bool) error {
	if majorVersion > 0 {
		fmt.Fprintf(w, "// %s\n", MakeFuncDocUrl(majorVersion, f.Name))
	}
	fmt.Fprintf(w, "func %s(", f.Name)
	for i, _ := range f.Parameters {
		p := &f.Parameters[i]
		fmt.Fprintf(w, "%s ", RenameIfReservedWord(p.Name))
		ptype, err := typeMap.Resolve(p.Type)
		if err != nil {
			return err
		}
		goType, _, err := CTypeToGoType(ptype, p.Out, p.Modifier)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s", goType)
		if i < len(f.Parameters)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	rtype, err := typeMap.Resolve(f.Return)
	if err != nil {
		return err
	}
	goType, _, err := CTypeToGoType(rtype, false, ParamModifierValue)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, ") %s {\n", goType)
	if rtype == "void" || rtype == "VOID" {
		if prototype {
			fmt.Fprintf(w, "	C.gl%s(", f.Name)
		} else {
			fmt.Fprintf(w, "	C.gogl%s(", f.Name)
		}
	} else {
		if prototype {
			fmt.Fprintf(w, "	return (%s)(C.gl%s(", goType, f.Name)
		} else {
			fmt.Fprintf(w, "	return (%s)(C.gogl%s(", goType, f.Name)
		}
	}
	for i, _ := range f.Parameters {
		p := &f.Parameters[i]
		ptype, err := typeMap.Resolve(p.Type)
		if err != nil {
			return err
		}
		_, cgoType, err := CTypeToGoType(ptype, p.Out, p.Modifier)
		if err != nil {
			return err
		}
		// HACK: handling pointers to pointers is tricky. We use unsafe.Pointer. Any better solution?
		if strings.HasPrefix(cgoType, "**") {
			fmt.Fprintf(w, "(%s)(unsafe.Pointer(%s))", cgoType, RenameIfReservedWord(p.Name))
		} else {
			fmt.Fprintf(w, "(%s)(%s)", cgoType, RenameIfReservedWord(p.Name))
		}
		if i < len(f.Parameters)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	if rtype == "void" || rtype == "VOID" {
		fmt.Fprintf(w, ")\n}\n")
	} else {
		fmt.Fprintf(w, "))\n}\n")
	}
	return nil
}

func writeGoInitDefinitions(w io.Writer, sortedFunctionCategories []string, functions FunctionCategories, initAll bool) error {
	for _, cat := range sortedFunctionCategories {
		fmt.Fprintf(w, "func Init%s() error {\n", GoName(cat))
		fmt.Fprintf(w, "\tvar ret C.int\n")
		fmt.Fprintf(w, "\tif ret = C.init_%s(); ret != 0 {\n", cat)
		fmt.Fprintf(w, "\t\treturn errors.New(\"unable to initialize %s\")\n", cat)
		fmt.Fprintf(w, "\t}\n")
		fmt.Fprintf(w, "\treturn nil\n")
		fmt.Fprintf(w, "}\n")
	}
	if initAll {
		fmt.Fprintf(w, "func Init() error {\n")
		fmt.Fprintf(w, "\tvar err error\n")
		for _, cat := range sortedFunctionCategories {
			fmt.Fprintf(w, "\tif err = Init%s(); err != nil {\n", GoName(cat))
			fmt.Fprintf(w, "\t\treturn err\n")
			fmt.Fprintf(w, "\t}\n")
		}
		fmt.Fprintf(w, "\treturn nil\n")
		fmt.Fprintf(w, "}\n")
	}
	return nil
}

func writeUtilityFunctions(w io.Writer) {
	fmt.Fprintln(w, `//Go bool to GL boolean.
func GLBool(b bool) Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

// GL boolean to Go bool.
func GoBool(b Boolean) bool {
	return b == TRUE
}

// Go string to GL string.
func GLString(str string) *Char {
	return (*Char)(C.CString(str))
}

// Allocates a GL string.
func GLStringAlloc(length Sizei) *Char {
	return (*Char)(C.malloc(C.size_t(length)))
}

// Frees GL string.
func GLStringFree(str *Char) {
	C.free(unsafe.Pointer(str))
}

// GL string (GLchar*) to Go string.
func GoString(str *Char) string {
	return C.GoString((*C.char)(str))
}

// GL string (GLubyte*) to Go string.
func GoStringUb(str *Ubyte) string {
	return C.GoString((*C.char)(unsafe.Pointer(str)))
}

// GL string (GLchar*) with length to Go string.
func GoStringN(str *Char, length Sizei) string {
	return C.GoStringN((*C.char)(str), C.int(length))
}

// Converts a list of Go strings to a slice of GL strings.
// Usefull for ShaderSource().
func GLStringArray(strs ...string) []*Char {
	strSlice := make([]*Char, len(strs))
	for i, s := range strs {
		strSlice[i] = (*Char)(C.CString(s))
	}
	return strSlice
}

// Free GL string slice allocated by GLStringArray().
func GLStringArrayFree(strs []*Char) {
	for _, s := range strs {
		C.free(unsafe.Pointer(s))
	}
}

// Add offset to a pointer. Usefull for VertexAttribPointer, TexCoordPointer, NormalPointer, ... 
func Offset(p Pointer, o uintptr) Pointer {
	return Pointer(uintptr(p) + o)
}
`)
}
