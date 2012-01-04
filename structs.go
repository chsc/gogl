// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Version

type Version struct {
	Major int
	Minor int
}

func MakeVersionFromString(version string) (Version, error) {
	split := strings.Split(version, ".")
	if len(split) != 2 {
		return Version{0, 0}, errors.New("Invalid version string.")
	}
	return MakeVersionFromMajorMinorString(split[0], split[1])
}

func MakeVersionFromMajorMinorString(major, minor string) (Version, error) {
	majorNumber, err := strconv.Atoi(major)
	if err != nil {
		return Version{0, 0}, errors.New("Invalid major version number.")
	}
	minorNumber, err := strconv.Atoi(minor)
	if err != nil {
		return Version{0, 0}, errors.New("Invalid minor version number.")
	}
	return Version{majorNumber, minorNumber}, nil
}

func (v Version) Compare(v2 Version) int {
	if v.Major < v2.Major {
		return -1
	} else if v.Major > v2.Major {
		return 1
	} else if v.Minor < v2.Minor {
		return -1
	} else if v.Minor > v2.Minor {
		return 1
	}
	return 0
}

func (v Version) Valid() bool {
	return v.Major != 0 || v.Minor != 0
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

// Functions

type ParamModifier int

const (
	ParamModifierValue ParamModifier = iota
	ParamModifierArray
	ParamModifierReference
)

type FunctionsInfo struct {
	Versions           []Version
	DeprecatedVersions []Version
	Categories         []string
	Passthru           []string
}

type Parameter struct {
	Name     string
	Type     string
	Out      bool
	Modifier ParamModifier
}

type Function struct {
	Name              string
	Parameters        []Parameter
	Return            string
	Version           Version
	DeprecatedVersion Version
	Category          string
	SubCategory       string
	Offset            string
	GlxRopCode        string
	GlxSingle         string
	WglFlags          string
	GlxFlags          string
	DlFlags           string
	GlfFlags          string
	GlxVendorPriv     string
	GlextMask         string
	Extension         string
	VectorEquiv       string
	GlxVectorEquiv    string
	Alias             string
	BeginEnd          string
}

type Functions []*Function
type FunctionCategories map[string]Functions

func (p *Parameter) String() string {
	out := ""
	modifier := ""
	if p.Out {
		out = "out "
	}
	switch p.Modifier {
	case ParamModifierArray:
		modifier = "[]"
	case ParamModifierReference:
		modifier = "&"
	}
	return fmt.Sprintf("%s %s%s%s", p.Type, out, modifier, p.Name)
}

func (f *Function) String() string {
	s := bytes.NewBufferString("")
	fmt.Fprintf(s, "%s %s(", f.Return, f.Name)
	for i, _ := range f.Parameters {
		p := &f.Parameters[i]
		fmt.Fprint(s, p)
		if i < len(f.Parameters)-1 {
			fmt.Fprint(s, ", ")
		}
	}
	fmt.Fprintf(s, ") (v%v, d%v)", f.Version, f.DeprecatedVersion)
	return string(s.Bytes())
}

func (fs Functions) Find(name string) *Function {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}

// Enums

type Enum struct {
	Name  string
	Value string
}

type Enums map[string]Enum
type EnumCategories map[string]Enums

func (e Enum) String() string {
	return fmt.Sprintf("%s = %s", e.Name, e.Value)
}

// Type maps

type TypeMap map[string]string

func (tm TypeMap) Resolve(t string) (string, error) {
	if t == "void" {
		return t, nil
	}
	if rt, ok := tm[t]; ok {
		return rt, nil
	}
	return t, fmt.Errorf("Unable to resolve type: %s.", t)
}

// Packages

type Package struct {
	Enums     EnumCategories
	Functions FunctionCategories
}

type Packages map[string]*Package
