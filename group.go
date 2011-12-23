// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"fmt"
	"strings"
)

type PackageGroupFunc func(category string) (packageNames []string)

func GroupEnumsAndFunctions(enums EnumCategories, functions FunctionCategories, sortFunc PackageGroupFunc) (packages Packages) {
	packages = make(Packages)
	for category, catEnums := range enums {
		if packageNames := sortFunc(category); packageNames != nil {
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
	for category, catFunctions := range functions {
		if packageNames := sortFunc(category); packageNames != nil {
			for _, packageName := range packageNames {
				if p, ok := packages[packageName]; ok {
					p.Functions[category] = catFunctions
				} else {
					p := &Package{make(EnumCategories), make(FunctionCategories)}
					p.Functions[category] = catFunctions
					packages[packageName] = p
				}
			}
		}
	}
	return
}

// Default grouping function.
// Currently only one deprecation level supported.
func GroupPackagesByVendorFunc(category string, supportedVersions []Version, deprecatedVersions []Version) (packageNames []string) {
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
				packages = append(packages, fmt.Sprintf("gl%d%d", ver.Major, ver.Minor))
				if ver.Compare(deprecatedVersions[0]) >= 0 {
					packages = append(packages, fmt.Sprintf("gl%d%dc", ver.Major, ver.Minor))
				}
			}
		}
	case CategoryDepVersion:
		for _, ver := range supportedVersions {
			if ver.Compare(pc.Version) >= 0 {
				if ver.Compare(deprecatedVersions[0]) < 0 {
					packages = append(packages, fmt.Sprintf("gl%d%d", ver.Major, ver.Minor))
				} else {
					packages = append(packages, fmt.Sprintf("gl%d%dc", ver.Major, ver.Minor))
				}
			}
		}
	}
	return packages
}
