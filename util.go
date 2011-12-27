// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"errors"
	"fmt"
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
	Type     CategoryType
	Version  Version
	Vendor   string
	Name     string
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
		return r
	},
	n)
}

func RenameReservedWord(word string) string {
	switch word {
	case "func", "type", "struct", "range", "map":
		return fmt.Sprintf("%s_", word)
	}
	return word
}
// Converts C types to Go types.
func CTypeToGoType(glType string, out, array bool) (ret string, err error) {
	// special cases:
	switch glType{
	case "void":
		ret = ""
		if array {
			ret = "Pointer"
		}
		if out {
			err = errors.New("Unsupported void type.")
		}
		return
	case "GLvoid":
		ret = "Pointer"
		return
	case "GLvoid*", "GLvoid* const":
		ret = "Pointer"
		if array {
			ret = "*" + ret
		}
		// out can be ignored
		return
	case "const GLubyte *":
		ret = "*Ubyte"
		if array {
			ret = "*" + ret
		}
		if out {
			err = errors.New("Unsopported out parameter.")
		}
		return
	case "GLchar*", "GLcharARB*":
		ret = "*Char"
		if array {
			ret = "*" + ret
		}
		return
	case "GLboolean*":
		ret = "*Boolean"
		if array {
			ret = "*" + ret
		}
		return
	case "GLhandleARB":
		ret = "Uint" // handle is uint
		if array {
			ret = "*" + ret
		}
		return
	case "GLDEBUGPROCARB", "GLDEBUGPROCAMD":
		ret = "Pointer" // TODO: Debug callback support?
		if array || out {
			err = errors.New("Unsupported type.")
		}
		return
	case "GLvdpauSurfaceNV":
		ret = "Pointer" // TODO: Debug callback support?
		if array {
			ret = "*" + ret
		}
		return
	case "GLsync":
		ret = "Pointer"
		if array || out {
			err = errors.New("Unsupported type.")
		}
		return
	case "struct _cl_context *", "struct _cl_event *":
		ret = "Pointer" // TODO: OpenCL context, event support?
		if array || out {
			err = errors.New("Unsupported type.")
		}
		return
	}
	// standard cases for primitive data types:
	switch glType {
	// base types
	case "GLenum":
		ret = "Enum"
	case "GLboolean":
		ret = "Boolean"
	case "GLbitfield":
		ret = "Bitfield"
	case "GLbyte":
		ret = "Byte"
	case "GLshort":
		ret = "Short"
	case "GLint":
		ret = "Int"
	case "GLsizei":
		ret = "Sizei"
	case "GLubyte":
		ret = "Ubyte"
	case "GLushort":
		ret = "Ushort"
	case "GLuint":
		ret = "Uint"
	case "GLhalf", "GLhalfNV":
		ret = "Half"
	case "GLfloat":
		ret = "Float"
	case "GLclampf":
		ret = "Clampf"
	case "GLdouble":
		ret = "Double"
	case "GLclampd":
		ret = "Clampd"
	//  
	case "GLchar", "GLcharARB":
		ret = "Char"
	// 
	case "GLintptr", "GLintptrARB":
		ret = "Intptr"
	case "GLsizeiptr", "GLsizeiptrARB":
		ret = "Sizeiptr"
	// 
	case "GLint64", "GLint64EXT":
		ret = "Int64"
	case "GLuint64", "GLuint64EXT":
		ret = "Uint64"

	default:
		err = errors.New("Unknown GL type: " + glType)
		return
	}
	if array {
		ret = "*" + ret
	}
	if out && !array {
		ret = "*" + ret
	}
	return
}

func MakeExtensionSpecDocUrl(vendor, extension string) string {
	return fmt.Sprintf("http://www.opengl.org/registry/specs/%s/%s.txt", vendor, extension)
}

func MakeFuncDocUrl(majorVersion int, fName string) string {
	return fmt.Sprintf("http://www.opengl.org/sdk/docs/man%d/xhtml/gl%s.xml", majorVersion, fName)
}
