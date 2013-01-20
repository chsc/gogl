// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"fmt"
	"os"
	"strings"
)

type EnumGroupFunc func(category string) (packageNames []string)
type FunctionGroupFunc func(function *Function) (packageNames []string)

func GroupEnumsAndFunctions(enums EnumCategories, functions FunctionCategories, sortFuncEnums EnumGroupFunc, sortFuncFuncts FunctionGroupFunc) (packages Packages) {
	packages = make(Packages)
	for category, catEnums := range enums {
		if packageNames := sortFuncEnums(category); packageNames != nil {
			for _, packageName := range packageNames {
				if p, ok := packages[packageName]; ok {
					p.Enums[category] = catEnums
				} else {
					p := &Package{make(EnumCategories), make(FunctionCategories)}
					p.Enums[category] = catEnums
					packages[packageName] = p
				}
			}
		}
	}
	for _, catFunctions := range functions {
		for _, function := range catFunctions {
			if packageNames := sortFuncFuncts(function); packageNames != nil {
				for _, packageName := range packageNames {
					if p, ok := packages[packageName]; ok {
						if _, ok = p.Functions[function.Category]; ok {
							p.Functions[function.Category] = append(p.Functions[function.Category], function)
						} else {
							p.Functions[function.Category] = []*Function{function}
						}
					} else {
						p := &Package{make(EnumCategories), make(FunctionCategories)}
						p.Functions[function.Category] = []*Function{function}
						packages[packageName] = p
					}
				}
			}

		}
	}
	return
}

// Default grouping function.
// Groups packages by GL versions and vendor extensions:
//  EXT, ARB, ATI, NV -> ext, arb, ati, nv package
//  VERSION_1_0, VERSION_1_1, VERSION_1_2 -> gl12

func GroupEnumsByVendorAndVersion(category string, supportedVersions []Version) (packageName []string) {
	pc, err := ParseCategoryString(category)
	if err != nil {
		return nil
	}
	packages := make([]string, 0, 8)
	switch pc.Type {
	case CategoryExtension:
		packages = append(packages, strings.ToLower(pc.Vendor))
	case CategoryVersion:
		for _, ver := range supportedVersions {
			if ver.Compare(pc.Version) >= 0 {
				fmt.Println("++", ver, pc)
				packages = append(packages, fmt.Sprintf("gl%d%d", ver.Major, ver.Minor))
			}
		}
	case CategoryDepVersion:
		fmt.Fprintf(os.Stderr, "Deprecated categories are not longer supported. Update your spec file.")
		return nil // not supported anymore, spec outdated
	}
	return packages
}

func GroupFunctionsByVendorAndVersion(function *Function, supportedVersions []Version) (packageName []string) {
	pc, err := ParseCategoryString(function.Category)
	if err != nil {
		return nil
	}
	packages := make([]string, 0, 8)
	switch pc.Type {
	case CategoryExtension:
		packages = append(packages, strings.ToLower(pc.Vendor))
	case CategoryVersion:
		if function.DeprecatedVersion.Valid() { // is deprecated 
			for _, ver := range supportedVersions {
				if ver.Compare(pc.Version) >= 0 {
					if ver.Compare(function.DeprecatedVersion) < 0 {
						packages = append(packages, fmt.Sprintf("gl%d%d", ver.Major, ver.Minor))
					} else {
						packages = append(packages, fmt.Sprintf("gl%d%dc", ver.Major, ver.Minor))
					}
				}
			}
		} else {
			for _, ver := range supportedVersions {
				if ver.Compare(pc.Version) >= 0 {
					packages = append(packages, fmt.Sprintf("gl%d%d", ver.Major, ver.Minor))
				}
			}

		}
	case CategoryDepVersion:
		fmt.Fprintf(os.Stderr, "Deprecated categories are not longer supported. Update your spec file.")
		return nil // not supported anymore, spec outdated
	}
	return packages
}

// Put everything into one package.
// E.g. for wgl/glx.

func GroupAllEnumsIntoOnePackage(category, packageName string) (packageNames []string) {
	return []string{packageName}
}

func GroupAllFunctionsIntoOnePackage(functions *Function, packageName string) (packageNames []string) {
	return []string{packageName}
}
