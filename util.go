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

func GLTypeToGoType(glType string, out, array bool) (ret string, err error) {
	// TODO: eval out, array...
	switch glType {
	case "void":
		ret = ""
	case "GLenum", "GLbitfield", "GLuint":
		ret = "uint32"
	case "GLsizei", "GLint":
		ret = "int32"
	case "GLushort", "GLhalf":
		ret = "uint16"
	case "GLshort":
		ret = "int16"
	case "GLboolean", "GLubyte":
		ret = "uint8"
	case "GLbyte":
		ret = "int8"
	case "GLchar":
		ret = "byte"
	case "GLfloat", "GLclampf":
		ret = "float32"
	case "GLdouble", "GLclampd":
		ret = "float64"
	case "GLintptr", "GLsizeiptr":
		ret = "int"
	case "GLvoid*":
		ret = "unsafe.Pointer"
	default:
		err = errors.New("Unknown GL type: " + glType)
	}
	return
}

func MakeExtensionSpecDocUrl(vendor, extension string) string {
	return fmt.Sprintf("http://www.opengl.org/registry/specs/%s/%s.txt", vendor, extension)
}

func MakeFuncDocUrl(majorVersion int, fName string) string {
	return fmt.Sprintf("http://www.opengl.org/sdk/docs/man%d/xhtml/gl%s.xml", majorVersion, fName)
}
