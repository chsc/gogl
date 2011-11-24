package main

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
