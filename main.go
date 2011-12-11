// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var (
	download *bool   = flag.Bool("download", false, "Just download spec files from url.")
	specUrl  *string = flag.String("url", KhronosRegistryBaseURL, "OpenGL specification location.")
	specDir  *string = flag.String("dir", "khronos_specs", "OpenGL specification directory.")
	// TODO: add additional flags ...
)

func main() {
	fmt.Printf("OpenGL binding generator for Go. Copyright (c) 2011 by Christoph Schunk.\n")
	flag.Parse()

	if *download {
		DownloadOpenGLSpecs(*specUrl, *specDir)
		return
	}

	fmt.Printf("Parsing enumext.spec file...\n")
	enumCategories, err := ReadEnumsFromFile(filepath.Join(*specDir, OpenGLEnumExtSpecFile))
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Parsing gl.tm file ...\n")
	typeMap, err := ReadTypeMapFromFile(filepath.Join(*specDir, OpenGLTypeMapFile))
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Parsing gl.spec file ...\n")
	funcCategories, funcInfo, err := ReadFunctionsFromFile(filepath.Join(*specDir, OpenGLSpecFile))
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Sorting extensions ...\n")
	packages := GroupEnumsAndFunctions(enumCategories, funcCategories,
		func(category string) (packageNames []string) {
			fmt.Printf("%s\n", category)
			return GroupPackagesByVendorFunc(category, funcInfo.Versions, funcInfo.DeprecatedVersions)
		})

	fmt.Printf("Generating packages %v  %v...\n", packages, funcInfo)
	if err := GeneratePackages(packages, typeMap); err != nil {
		panic(err.Error())
	}

	if false {
		// TODO: This output is temporary for debugging
		fmt.Println("Supported versions:")
		fmt.Println(funcInfo.Versions)
		fmt.Println("Deprecated versions:")
		fmt.Println(funcInfo.DeprecatedVersions)

		fmt.Println("Enums:")
		for category, enums := range enumCategories {
			fmt.Printf("  %v\n", category)
			for _, enum := range enums {
				fmt.Printf("    %v = %v\n", enum.Name, enum.Value)
			}
		}
		fmt.Println("Types:")
		for abstractType, cType := range typeMap {
			fmt.Printf("  %v -> %v\n", abstractType, cType)
		}
		fmt.Println("Functions:")
		for category, functions := range funcCategories {
			fmt.Printf("  %v\n", category)
			for _, function := range functions {
				fmt.Printf("    %v\n", function.Name)
				if function.Version.Valid() {
					fmt.Printf("      Version: %v\n", function.Version)
				}
				if function.DeprecatedVersion.Valid() {
					fmt.Printf("      Deprecated Version: %v\n", function.DeprecatedVersion)
				}
				fmt.Printf("      Return Type: %v\n", function.Return)
				if len(function.Parameters) > 0 {
					fmt.Printf("      Parameters:\n")
					for _, param := range function.Parameters {
						if param.InArray {
							fmt.Printf("        %v %v in array\n", param.Name, param.Type)
						} else {
							fmt.Printf("        %v %v\n", param.Name, param.Type)
						}
					}
				} else {
					fmt.Printf("      0 Parameters\n")
				}
			}
		}
	}
}
