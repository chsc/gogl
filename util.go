// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	versionPrefix    = "VERSION_"
	depVersionSuffix = "_DEPRECATED"
	verPrefixLen     = len(versionPrefix)
	depVerSuffixLen  = len(depVersionSuffix)
)

type CategoryType int32

const (
	CategoryInvalid CategoryType = iota
	CategoryVersion
	CategoryDepVersion
	CategoryExtension
)

type ParsedCategory struct {
	Type    CategoryType
	Version Version
	Vendor  string
	Name    string
}

func (pc ParsedCategory) String() string {
	switch pc.Type {
	case CategoryExtension:
		return fmt.Sprintf("Extension: vendor: %s, name: %s", pc.Vendor, pc.Name)
	case CategoryVersion:
		return fmt.Sprintf("Version: %v %s", pc.Version, pc.Vendor)
	case CategoryDepVersion:
		return fmt.Sprintf("Deprecated version: %v %s", pc.Version, pc.Vendor)
	}
	return "Invalid category"
}

func ParseCategoryString(category string) (ParsedCategory, error) {
	if strings.HasPrefix(category, versionPrefix) {
		ver := category[verPrefixLen:]
		ctype := CategoryVersion
		if strings.HasSuffix(ver, depVersionSuffix) {
			ver = ver[:len(ver)-depVerSuffixLen]
			ctype = CategoryDepVersion
		}
		if splitVer := strings.Split(ver, "_"); len(splitVer) == 2 {
			version, err := MakeVersionFromMajorMinorString(splitVer[0], splitVer[1])
			if err != nil {
				return ParsedCategory{}, errors.New("Invalid category string.")
			}
			return ParsedCategory{ctype, version, "", ""}, nil
		}
		return ParsedCategory{}, errors.New("Invalid category string.")
	} else if splitCat := strings.SplitN(category, "_", 2); len(splitCat) == 2 {
		return ParsedCategory{CategoryExtension, Version{}, splitCat[0], splitCat[1]}, nil
	}
	return ParsedCategory{}, errors.New("Invalid category string.")
}

// Converts strings with underscores to Go-like names. e.g.: bla_blub_foo -> BlaBlubFoo
func GoName(n string) string {
	prev := '_'
	return strings.Map(func(r rune) rune {
		if r == '_' {
			prev = r
			return -1
		}
		if prev == '_' {
			prev = r
			return unicode.ToTitle(r)
		}
		prev = r
		return unicode.ToLower(r)
	},
		n)
}

func RenameIfReservedWord(word string) string {
	switch word {
	case "func", "type", "struct", "range", "map", "string", "near", "far":
		return fmt.Sprintf("%s_", word)
	}
	return word
}

// append X suffix if enum name starts with a digit
func CleanEnumName(enum string) string {
	if strings.IndexAny(enum, "0123456789") == 0 {
		return fmt.Sprintf("X%s", enum)
	}
	return enum
}

// Converts C types to Go types.
func CTypeToGoType(cType string, out bool, mod ParamModifier) (goType, cgoType string, err error) {
	// special cases:
	switch cType {
	case "void":
		goType = ""
		cgoType = ""
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "Pointer"
			cgoType = "unsafe.Pointer"
		}
		if out {
			err = errors.New("Unsupported void type.")
		}
		return
	case "GLvoid":
		if mod == ParamModifierArray || mod == ParamModifierReference || out {
			goType = "Pointer"
			cgoType = "unsafe.Pointer"
		} else {
			err = errors.New("Unsupported GLvoid type.")
		}
		return
	case "GLvoid*", "GLvoid* const":
		goType = "Pointer"
		cgoType = "unsafe.Pointer"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "*" + goType
			cgoType = "*" + cgoType
		}
		// out can be ignored
		return
	case "const GLubyte *":
		goType = "*Ubyte"
		cgoType = "*C.GLubyte"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "*" + goType
			cgoType = "*" + cgoType
		}
		if out {
			err = errors.New("Unsupported out parameter.")
		}
		return
	case "GLchar*", "GLchar* const":
		goType = "*Char"
		cgoType = "*C.GLchar"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "*" + goType
			cgoType = "*" + cgoType
		}
		return
	case "GLcharARB*":
		goType = "*Char"
		cgoType = "*C.GLcharARB"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "*" + goType
			cgoType = "*" + cgoType
		}
		return
	case "GLboolean*":
		goType = "*Boolean"
		cgoType = "*C.GLboolean"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "*" + goType
			cgoType = "*" + cgoType
		}
		return
	case "GLhandleARB":
		goType = "Uint" // handle is uint
		cgoType = "C.GLhandleARB"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "*" + goType
			cgoType = "*" + cgoType
		}
		return
	case "GLDEBUGPROCARB", "GLDEBUGPROCAMD":
		goType = "Pointer" // TODO: Debug callback support?
		cgoType = "*[0]byte"
		if mod == ParamModifierArray || mod == ParamModifierReference || out {
			err = errors.New("Unsupported type.")
		}
		return
	case "GLvdpauSurfaceNV":
		goType = "Intptr" // TODO: vdpau support?
		cgoType = "C.GLvdpauSurfaceNV"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			goType = "*" + goType
			cgoType = "*" + cgoType
		}
		return
	case "GLsync":
		goType = "Sync"
		cgoType = "C.GLsync"
		if mod == ParamModifierArray || mod == ParamModifierReference || out {
			err = errors.New("Unsupported type.")
		}
		return
	case "struct _cl_context *", "struct _cl_event *":
		goType = "Pointer" // TODO: OpenCL context, event support?
		cgoType = "*[0]byte"
		if mod == ParamModifierArray || mod == ParamModifierReference || out {
			err = errors.New("Unsupported type.")
		}
		return
	}
	// standard cases for primitive data types:
	switch cType {
	// base types
	case "GLenum":
		goType = "Enum"
		cgoType = "C.GLenum"
	case "GLboolean":
		goType = "Boolean"
		cgoType = "C.GLboolean"
	case "GLbitfield":
		goType = "Bitfield"
		cgoType = "C.GLbitfield"
	case "GLbyte":
		goType = "Byte"
		cgoType = "C.GLbyte"
	case "GLshort":
		goType = "Short"
		cgoType = "C.GLshort"
	case "GLint":
		goType = "Int"
		cgoType = "C.GLint"
	case "GLsizei":
		goType = "Sizei"
		cgoType = "C.GLsizei"
	case "GLubyte":
		goType = "Ubyte"
		cgoType = "C.GLubyte"
	case "GLushort":
		goType = "Ushort"
		cgoType = "C.GLushort"
	case "GLuint":
		goType = "Uint"
		cgoType = "C.GLuint"
	case "GLhalf":
		goType = "Half"
		cgoType = "C.GLhalf"
	case "GLhalfNV":
		goType = "Half"
		cgoType = "C.GLhalfNV"
	case "GLfloat":
		goType = "Float"
		cgoType = "C.GLfloat"
	case "GLclampf":
		goType = "Clampf"
		cgoType = "C.GLclampf"
	case "GLdouble":
		goType = "Double"
		cgoType = "C.GLdouble"
	case "GLclampd":
		goType = "Clampd"
		cgoType = "C.GLclampd"
	//  
	case "GLchar":
		goType = "Char"
		cgoType = "C.GLchar"
	case "GLcharARB":
		goType = "Char"
		cgoType = "C.GLcharARB"
	// 
	case "GLintptr":
		goType = "Intptr"
		cgoType = "C.GLintptr"
	case "GLintptrARB":
		goType = "Intptr"
		cgoType = "C.GLintptrARB"
	case "GLsizeiptr":
		goType = "Sizeiptr"
		cgoType = "C.GLsizeiptr"
	case "GLsizeiptrARB":
		goType = "Sizeiptr"
		cgoType = "C.GLsizeiptrARB"
	// 
	case "GLint64":
		goType = "Int64"
		cgoType = "C.GLint64"
	case "GLint64EXT":
		goType = "Int64"
		cgoType = "C.GLint64EXT"
	case "GLuint64":
		goType = "Uint64"
		cgoType = "C.GLuint64"
	case "GLuint64EXT":
		goType = "Uint64"
		cgoType = "C.GLuint64EXT"
	default:
		err = errors.New("Unknown GL type: " + cType)
		return
	}
	if mod == ParamModifierArray || mod == ParamModifierReference || out {
		goType = "*" + goType
		cgoType = "*" + cgoType
	}
	return
}

func MakeExtensionSpecDocUrl(vendor, extension string) string {
	return fmt.Sprintf("http://www.opengl.org/registry/specs/%s/%s.txt", vendor, extension)
}

func MakeGLDocUrl(majorVersion int) string {
	manVer := ""
	if majorVersion >= 3 {
		manVer = strconv.Itoa(majorVersion)
	}
	return fmt.Sprintf("http://www.opengl.org/sdk/docs/man%s", manVer)
}

func MakeFuncDocUrl(majorVersion int, fName string) string {
	manVer := ""
	if majorVersion >= 3 {
		manVer = strconv.Itoa(majorVersion)
	}
	return fmt.Sprintf("http://www.opengl.org/sdk/docs/man%s/xhtml/gl%s.xml", manVer, fName)
}
