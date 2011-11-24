package main

import (
	"fmt"
)

type Package struct {
	enums     EnumCategories
	functions FunctionCategories
}

type Packages map[string]*Package

func (p *Package) String() string {
	return fmt.Sprintf("%s %s", p.enums, p.functions)
}

type PackageSortFunc func(category string) (packageName string, ok bool)

func SortEnumsAndFunctions(enums EnumCategories, functions FunctionCategories, sortFunc PackageSortFunc) (packages Packages) {
	packages = make(Packages)
	for category, catEnums := range enums {
		if packageName, ok := sortFunc(category); ok {
			if p, ok := packages[packageName]; ok {
				p.enums[category] = catEnums
			} else {
				p := &Package{make(EnumCategories), make(FunctionCategories)}
				p.enums[category] = catEnums
				packages[packageName] = p
			}
		}
	}
	for category, catFunctions := range functions {
		if packageName, ok := sortFunc(category); ok {
			if p, ok := packages[packageName]; ok {
				p.functions[category] = catFunctions
			} else {
				p := &Package{make(EnumCategories), make(FunctionCategories)}
				p.functions[category] = catFunctions
				packages[packageName] = p
			}
		}
	}
	return
}

