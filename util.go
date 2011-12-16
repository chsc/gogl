// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

var (
	versionRE    = regexp.MustCompile("^VERSION_([0-9]+)_([0-9]+)$")
	versionDepRE = regexp.MustCompile("^VERSION_([0-9]+)_([0-9]+)_DEPRECATED$")
	extensionRE  = regexp.MustCompile("^([A-Za-z0-9]+)_([A-Za-z0-9_]+)$")
)

type CategoryType int32

const (
	CategoryInvalid CategoryType = iota
	CategoryVersion
	CategoryDeprecatedVersion
	CategoryExtension
)

type ParsedCategoryString struct {
	CategoryType CategoryType
	Version      Version
	Vendor       string
	Name         string
}

func ParseCategoryString(category string) (ParsedCategoryString, error) {
	if versionDep := versionDepRE.FindStringSubmatch(category); versionDep != nil {
		v, err := MakeVersionFromMajorMinorString(versionDep[1], versionDep[2])
		if err != nil {
			return ParsedCategoryString{CategoryInvalid, Version{0, 0}, "", ""}, err
		}
		return ParsedCategoryString{CategoryDeprecatedVersion, v, "", ""}, nil
	} else if version := versionRE.FindStringSubmatch(category); version != nil {
		v, err := MakeVersionFromMajorMinorString(version[1], version[2])
		if err != nil {
			return ParsedCategoryString{CategoryInvalid, Version{0, 0}, "", ""}, err
		}
		return ParsedCategoryString{CategoryVersion, v, "", ""}, nil
	} else if extension := extensionRE.FindStringSubmatch(category); extension != nil {
		return ParsedCategoryString{CategoryExtension, Version{0, 0}, extension[1], extension[2]}, nil
	}
	return ParsedCategoryString{CategoryInvalid, Version{0, 0}, "", ""}, errors.New("Unable to parse category.")
}

// Converts strings with underscores to Go-like names. e.g.: bla_blub_foo -> BlaBlubFoo
func GoName(n string) string {
	prev := '_'
	return strings.Map(func(r int) int {
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
