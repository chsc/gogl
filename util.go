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
	} else if len(category) > 0 {
		return ParsedCategory{CategoryExtension, Version{}, category, ""}, nil
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
		//if out {
		//	err = errors.New("Unsupported void type.")
		//}
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
	// glx
	case "Display":
		goType = "Pointer"
		cgoType = "C.Display"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			//goType = "*" + goType
			cgoType = "*" + cgoType
		}
		return
	case "Display *":
		goType = "Pointer"
		cgoType = "*C.Display"
		if mod == ParamModifierArray || mod == ParamModifierReference {
			//goType = "*" + goType
			cgoType = "*" + cgoType
		}
		return
	case "GLXFBConfig":
		goType = "Pointer"
		cgoType = "C.GLXFBConfig"
		return
	case "GLXFBConfig *":
		goType = "Pointer"
		cgoType = "C.GLXFBConfig *"
		return
	case "GLXContext", "const GLXContext":
		goType = "Pointer"
		cgoType = "C.GLXContext"
		return
	case "GLXContextID":
		goType = "uint32"
		cgoType = "C.GLXContextID"
		return
	case "GLXDrawable":
		goType = "Pointer"
		cgoType = "C.GLXDrawable"
		return
	case "XVisualInfo":
		goType = "Pointer"
		cgoType = "C.XVisualInfo"
		return
	case "XVisualInfo *":
		goType = "Pointer"
		cgoType = "*C.XVisualInfo"
		return
	case "Pixmap":
		goType = "Pointer"
		cgoType = "C.Pixmap"
		return
	case "GLXPixmap":
		goType = "Pointer"
		cgoType = "C.GLXPixmap"
		return
	case "Colormap":
		goType = "Pointer"
		cgoType = "C.Colormap"
		return
	case "GLXPbuffer":
		goType = "Pointer"
		cgoType = "C.GLXPbuffer"
		return
	case "GLXPbufferSGIX":
		goType = "Pointer"
		cgoType = "C.GLXPbufferSGIX"
		return
	case "GLXFBConfigSGIX":
		goType = "Pointer"
		cgoType = "C.GLXFBConfigSGIX"
		return
	case "GLXFBConfigSGIX *":
		goType = "Pointer"
		cgoType = "*C.GLXFBConfigSGIX"
		return
	case "DMparams":
		goType = "Pointer"
		cgoType = "C.DMparams"
		return
	case "DMbuffer":
		goType = "Pointer"
		cgoType = "C.DMbuffer"
		return
	case "__GLXextFuncPtr":
		goType = "Pointer"
		cgoType = "C.__GLXextFuncPtr"
		return
	case "GLXVideoDeviceNV":
		goType = "Pointer"
		cgoType = "C.GLXVideoDeviceNV"
		return
	case "GLXVideoCaptureDeviceNV":
		goType = "Pointer"
		cgoType = "C.GLXVideoCaptureDeviceNV"
		return
	case "GLXVideoCaptureDeviceNV *":
		goType = "Pointer"
		cgoType = "*C.GLXVideoCaptureDeviceNV"
		return
	case "GLXHyperpipeNetworkSGIX *":
		goType = "Pointer"
		cgoType = "*C.GLXHyperpipeNetworkSGIX"
		return
	case "GLXHyperpipeConfigSGIX *":
		goType = "Pointer"
		cgoType = "*C.GLXHyperpipeConfigSGIX"
		return
	case "GLXHyperpipeConfigSGIX":
		if mod == ParamModifierArray || mod == ParamModifierReference || out {
			goType = "Pointer"
			cgoType = "*C.GLXHyperpipeConfigSGIX"
		} else {
			err = errors.New("Unsupported GLXHyperpipeConfigSGIX type.")
		}
		return
	case "GLXVideoSourceSGIX":
		goType = "Pointer"
		cgoType = "*C.GGLXVideoSourceSGIX"
		return
	case "Window":
		goType = "Pointer"
		cgoType = "C.Window"
		return
	case "GLXWindow":
		goType = "Pointer"
		cgoType = "C.GLXWindow"
		return
	case "VLServer":
		goType = "Pointer"
		cgoType = "C.VLServer"
		return
	case "VLPath":
		goType = "Pointer"
		cgoType = "C.VLPath"
		return
	case "VLNode":
		goType = "Pointer"
		cgoType = "C.VLNode"
		return
	case "unsigned int *":
		goType = "*uint32"
		cgoType = "*C.uint32"
		return

	// wgl
	case "VOID":
		goType = ""
		cgoType = ""
		return
	case "LPVOID":
		goType = "Pointer"
		cgoType = "C.LPVOID"
		return
	case "void*":
		goType = "Pointer"
		cgoType = "C.LPVOID"
		return
	case "HANDLE":
		goType = "Pointer"
		cgoType = "C.HANDLE"
		return
	case "HDC":
		goType = "Pointer"
		cgoType = "C.HDC"
		return
	case "HGLRC":
		goType = "Pointer"
		cgoType = "C.HGLRC"
		return
	case "HPBUFFERARB":
		goType = "Pointer"
		cgoType = "C.HPBUFFERARB"
		return
	case "HPBUFFEREXT":
		goType = "Pointer"
		cgoType = "C.HPBUFFEREXT"
		return
	case "HGPUNV":
		goType = "Pointer"
		cgoType = "C.HGPUNV"
		return
	case "PGPU_DEVICE":
		goType = "Pointer"
		cgoType = "C.PGPU_DEVICE"
		return
	case "HVIDEOOUTPUTDEVICENV":
		goType = "Pointer"
		cgoType = "C.HVIDEOOUTPUTDEVICENV"
		return
	case "HVIDEOINPUTDEVICENV":
		goType = "Pointer"
		cgoType = "C.HVIDEOINPUTDEVICENV"
		return
	case "HPVIDEODEV":
		goType = "Pointer"
		cgoType = "C.HPVIDEODEV"
		return
	case "COLORREF":
		goType = "Pointer"
		cgoType = "C.COLORREF"
		return
	case "LAYERPLANEDESCRIPTOR":
		goType = "[]byte"
		cgoType = "C.LAYERPLANEDESCRIPTOR"
		return
	case "LPCSTR":
		goType = "*byte"
		cgoType = "C.LPCSTR"
		return
	case "PIXELFORMATDESCRIPTOR":
		goType = "[]byte"
		cgoType = "C.PIXELFORMATDESCRIPTOR"
		return
	case "const char *":
		goType = "*byte"
		cgoType = "*C.char"
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
	// glx
	case "int":
		goType = "Int"
		cgoType = "C.int"
	case "int32_t":
		goType = "int32"
		cgoType = "C.int32_t"
	case "int64_t":
		goType = "int64"
		cgoType = "C.int64_t"
	case "Bool":
		goType = "int"
		cgoType = "C.int"
	case "Status":
		goType = "int"
		cgoType = "C.int"
	case "long":
		goType = "int32"
		cgoType = "C.long"

	//wgl
	case "UINT":
		goType = "uint32"
		cgoType = "C.UINT"
	case "INT":
		goType = "int32"
		cgoType = "C.INT"
	case "INT32":
		goType = "int32"
		cgoType = "C.INT32"
	case "INT64":
		goType = "int64"
		cgoType = "C.INT64"
	case "USHORT":
		goType = "uint16"
		cgoType = "C.USHORT"
	case "DWORD":
		goType = "uint32"
		cgoType = "C.DWORD"
	case "FLOAT":
		goType = "float32"
		cgoType = "C.FLOAT"
	case "BOOL":
		goType = "int32"
		cgoType = "C.BOOL"
	case "PROC":
		goType = "Pointer"
		cgoType = "C.PROC"
	case "float":
		goType = "float32"
		cgoType = "C.float32"
	case "unsigned int":
		goType = "uint32"
		cgoType = "C.uint32"
	case "unsigned long":
		goType = "uint32"
		cgoType = "C.unsigned_long"
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
	return fmt.Sprintf("https://www.opengl.org/registry/specs/%s/%s.txt", vendor, extension)
}

func MakeGLDocUrl(majorVersion int) string {
	manVer := ""
	if majorVersion >= 3 {
		manVer = strconv.Itoa(majorVersion)
	}
	return fmt.Sprintf("https://www.opengl.org/sdk/docs/man%s", manVer)
}

func MakeFuncDocUrl(majorVersion int, fName string) string {
	manVer := ""
	if majorVersion >= 3 {
		manVer = strconv.Itoa(majorVersion)
	}
	return fmt.Sprintf("https://www.opengl.org/sdk/docs/man%s/xhtml/gl%s.xml", manVer, fName)
}
