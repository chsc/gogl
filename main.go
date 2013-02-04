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

func generatePackages(
	specsDir string,
	enumextFile string,
	specFiles []string,
	typeMapFiles []string,
	singlePackage, prefix string) error {

	fmt.Printf("Parsing %s file ...\n", enumextFile)
	enumCategories, err := ReadEnumsFromFile(filepath.Join(specsDir, enumextFile))
	if err != nil {
		return err
	}

	fmt.Printf("Parsing %v files ...\n", typeMapFiles)
	tmPaths := make([]string, len(typeMapFiles))
	for i, tmFile := range typeMapFiles {
		tmPaths[i] = filepath.Join(specsDir, tmFile)
	}
	typeMap, err := ReadTypeMapFromFiles(tmPaths)
	if err != nil {
		return err
	}

	fmt.Printf("Parsing %v files ...\n", specFiles)
	specPaths := make([]string, len(specFiles))
	for i, specFile := range specFiles {
		specPaths[i] = filepath.Join(specsDir, specFile)
	}
	funcCategories, funcInfo, err := ReadFunctionsFromFiles(specPaths)
	if err != nil {
		return err
	}

	fmt.Printf("Grouping extensions ...\n")
	var egf EnumGroupFunc
	var fgf FunctionGroupFunc
	if len(singlePackage) == 0 { // create multiple packages
		egf = func(category string) (packageNames []string) {
			return GroupEnumsByVendorAndVersion(
				category,
				funcInfo.Versions)
		}
		fgf = func(function *Function) (packageNames []string) {
			return GroupFunctionsByVendorAndVersion(
				function,
				funcInfo.Versions)
		}
	} else { // create single package
		egf = func(category string) (packageNames []string) {
			return GroupAllEnumsIntoOnePackage(
				category,
				singlePackage)
		}
		fgf = func(function *Function) (packageNames []string) {
			return GroupAllFunctionsIntoOnePackage(
				function,
				singlePackage)
		}
	}
	packages := GroupEnumsAndFunctions(enumCategories, funcCategories, egf, fgf)

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
	if err := GeneratePackages(packages, funcInfo, typeMap, prefix); err != nil {
		return err
	}

	return nil
}

func generateGoPackages(specsDir string) {
	// generate gl and ext packages
	if err := generatePackages(
		specsDir,
		OpenGLEnumExtSpecFile,
		[]string{OpenGLSpecFile},
		[]string{OpenGLTypeMapFile},
		"", "gl"); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// generate wgl package
	if err := generatePackages(
		specsDir,
		GLXEnumExtSpecFile,
		[]string{GLXSpecFile, GLXExtSpecFile},
		[]string{OpenGLTypeMapFile, GLXTypeMapFile},
		"glx", "glx"); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// generate wgl package
	if err := generatePackages(
		specsDir,
		WGLEnumExtSpecFile,
		[]string{WGLSpecFile, WGLExtSpecFile},
		[]string{OpenGLTypeMapFile, WGLTypeMapFile},
		"wgl", "wgl"); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func downloadSpec(name string, args []string) {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	src := fs.String("src", "alfonse", "Source URL or 'khronos' or 'alfonse'.")
	odir := fs.String("odir", "glspecs", "Output directory for spec files.")
	fs.Parse(args)
	fmt.Println("Downloading specs ...")
	switch *src {
	case "alfonse":
		DownloadOpenGLSpecs(AlfonseSpecsBaseURL, *odir)
	case "khronos":
		DownloadOpenGLSpecs(KhronosRegistryBaseURL, *odir)
	default:
		DownloadOpenGLSpecs(*src, *odir)
	}
}

func downloadDoc(name string, args []string) {
	fmt.Println("Download docs not implemented.")
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	fs.String("src", "", "Source URL.")
	fs.String("odir", "gldocs", "Output directory for doc files.")
	fs.Parse(args)
	fmt.Println("Downloading docs ...")
	// TODO: download docs
}

func generate(name string, args []string) {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	sdir := fs.String("sdir", "glspecs", "OpenGL spec directory.")
	_ = fs.String("ddir", "gldocs", "Documentation directory (currently not used).")
	fs.Parse(args)
	fmt.Println("Generate Bindings ...")
	generateGoPackages(*sdir)
}

func printUsage(name string) {
	fmt.Printf("Usage:     %s command [arguments]\n", name)
	fmt.Println("Commands:")
	fmt.Println(" dlspec    Download spec files.")
	fmt.Println(" dldoc     Download documentation files.")
	fmt.Println(" gen       Generate bindings.")
	fmt.Printf("Type %s <command> -help for a detailed command description.\n", name)
}

func main() {
	fmt.Println("OpenGL binding generator for the Go programming language (http://golang.org).")
	fmt.Println("Copyright (c) 2011-2013 by Christoph Schunk. All rights reserved.")

	name := os.Args[0]
	args := os.Args[1:]

	if len(args) < 1 {
		printUsage(name)
		os.Exit(-1)
	}

	command := args[0]

	switch command {
	case "dlspec":
		downloadSpec("dlspec", args[1:])
	case "dldoc":
		downloadDoc("dldoc", args[1:])
	case "gen":
		generate("gen", args[1:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command '%s'.", command)
		printUsage(name)
		os.Exit(-1)
	}
}
