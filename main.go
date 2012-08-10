// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	download *bool   = flag.Bool("download", false, "Just download spec files from url.")
	specUrl  *string = flag.String("url", KhronosRegistryBaseURL, "OpenGL specification url.")
	specDir  *string = flag.String("dir", "khronos_specs", "OpenGL specification directory.")
	// TODO: add additional flags ...
)

func generatePackages(
	enumextFile string,
	specFiles []string,
	typeMapFiles []string,
	singlePackage string) error {

	fmt.Printf("Parsing %s file ...\n", enumextFile)
	enumCategories, err := ReadEnumsFromFile(filepath.Join(*specDir, enumextFile))
	if err != nil {
		return err
	}

	fmt.Printf("Parsing %v files ...\n", typeMapFiles)
	tmPaths := make([]string, len(typeMapFiles))
	for i, tmFile := range typeMapFiles {
		tmPaths[i] = filepath.Join(*specDir, tmFile)
	}
	typeMap, err := ReadTypeMapFromFiles(tmPaths)
	if err != nil {
		return err
	}

	fmt.Printf("Parsing %v files ...\n", specFiles)
	specPaths := make([]string, len(specFiles))
	for i, specFile := range specFiles {
		specPaths[i] = filepath.Join(*specDir, specFile)
	}
	funcCategories, funcInfo, err := ReadFunctionsFromFiles(specPaths)
	if err != nil {
		return err
	}

	fmt.Printf("Grouping extensions ...\n")
	var gf PackageGroupFunc
	if len(singlePackage) == 0 {
		gf = func(category string) (packageNames []string) {
			return GroupCategoriesByVendorAndVersion(
				category,
				funcInfo.Versions,
				funcInfo.DeprecatedVersions)
		}
	} else {
		gf = func(category string) (packageNames []string) {
			return GroupCategoriesIntoOnePackage(
				category,
				singlePackage)
		}
	}
	packages := GroupEnumsAndFunctions(enumCategories, funcCategories, gf)

	// TODO: This output is temporary for debugging
	if false {
		fmt.Println("Passthrus:")
		for cat, pts := range funcInfo.Passthru {
			fmt.Printf(" %s:\n", cat)
			for _, pt := range pts {
				fmt.Printf("  %s\n", pt)
			}
		}
	}
	if false {
		fmt.Println("Supported versions:")
		fmt.Println(funcInfo.Versions)
		fmt.Println("Deprecated versions:")
		fmt.Println(funcInfo.DeprecatedVersions)

		fmt.Println("Enums:")
		for category, enums := range enumCategories {
			fmt.Printf("  %v\n", category)
			for _, enum := range enums {
				fmt.Printf("    %v\n", enum)
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
				fmt.Printf("    %v\n", function)
			}
		}
	}

	fmt.Printf("Generating packages ...\n")
	if err := GeneratePackages(packages, funcInfo, typeMap); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Printf("OpenGL binding generator for Go. Copyright (c) 2011-2012 by Christoph Schunk.\n")
	flag.Parse()

	if *download {
		DownloadOpenGLSpecs(*specUrl, *specDir)
		return
	}
	var err error


	err = generatePackages(
		OpenGLEnumExtSpecFile,
		[]string{OpenGLSpecFile},
		[]string{OpenGLTypeMapFile},
		"")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// TODO: add glxext
	err = generatePackages(
		GLXEnumExtSpecFile,
		[]string{GLXSpecFile, GLXExtSpecFile},
		[]string{OpenGLTypeMapFile, GLXTypeMapFile},
		"glx")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// TODO: add wglext
	err = generatePackages(
		WGLEnumExtSpecFile,
		[]string{WGLSpecFile, WGLExtSpecFile},
		[]string{OpenGLTypeMapFile, WGLTypeMapFile},
		"wgl")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
