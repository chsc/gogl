// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	funcCommentLineRE     = regexp.MustCompile("^#")
	funcRE                = regexp.MustCompile("^(\\w+)\\([\\w\\s,]*\\)")
	funcReturnRE          = regexp.MustCompile("^return\\s+(\\w+)")
	funcParamRE           = regexp.MustCompile("^param\\s+(\\w+)\\s+(\\w+)\\s+(in|out)\\s+(array|value|reference)")
	funcVectorEquivRE     = regexp.MustCompile("^vectorequiv\\s+(\\w+)")
	funcGlxVectorEquivRE  = regexp.MustCompile("^glxvectorequiv\\s+(\\w+)")
	funcCategoryRE        = regexp.MustCompile("^category\\s+(\\w+)")
	funcSubCategoryRE     = regexp.MustCompile("^subcategory\\s+(\\w+)")
	funcVersionRE         = regexp.MustCompile("^version\\s+([0-9]+)\\.([0-9]+)")
	funcDeprecatedRE      = regexp.MustCompile("^deprecated\\s+([0-9]+)\\.([0-9]+)")
	funcOffsetRE          = regexp.MustCompile("^offset\\s+([0-9\\*\\?]+)")
	funcGlxRopCodeRE      = regexp.MustCompile("^glxropcode\\s+([0-9\\*\\?]+)")
	funcGlxSingleRE       = regexp.MustCompile("^glxsingle\\s+([0-9\\?\\*]+)")
	funcWglFlagsRE        = regexp.MustCompile("^wglflags(\\s+([\\w\\*\\- ]+))")
	funcGlxFlagsRE        = regexp.MustCompile("^glxflags(\\s+([\\w\\*\\- ]+))")
	funcDlFlagsRE         = regexp.MustCompile("^dlflags\\s+([\\w\\*\\- ]+)")
	funcGlfFlagsRE        = regexp.MustCompile("^glfflags\\s+([\\w\\*\\- ]+)")
	funcGlxVendorPrivRE   = regexp.MustCompile("^glxvendorpriv\\s+([0-9\\?]+)")
	funcGlextMaskRE       = regexp.MustCompile("^glextmask\\s+(\\w+)")
	funcExtensionRE       = regexp.MustCompile("^extension$|^extension\\s+([\\w\\- ]*)")
	funcAliasRE           = regexp.MustCompile("^alias\\s+(\\w+)")
	funcBeginEndRE        = regexp.MustCompile("^beginend\\s+([\\w\\*\\-]+)")
	funcAllCategoriesRE   = regexp.MustCompile("^category:\\s*([\\w ]+)")
	funcAllVersionsRE     = regexp.MustCompile("^version:\\s*([0-9\\. ]+)")
	funcAllDeprVersionsRE = regexp.MustCompile("^deprecated:\\s*([0-9\\. ]+)")
	funcPassthruRE        = regexp.MustCompile("^passthru:\\s*(.*)")
	funcNewCategoryRE     = regexp.MustCompile("^newcategory:\\s*(\\w+)")
	//funcIgnoreRE          = regexp.MustCompile("^[a-z\\-]+:.*")
)

func ReadFunctionsFromFile(name string) (FunctionCategories, *FunctionsInfo, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	return ReadFunctions(file)
}

func ReadFunctions(r io.Reader) (FunctionCategories, *FunctionsInfo, error) {
	var currentFunction *Function = nil
	currentCategory := "global"
	functions := make(FunctionCategories)
	finfo := &FunctionsInfo{make([]Version, 0, 8), make([]Version, 0, 8), make([]string, 0, 8), make(map[string][]string)}

	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}
		line = strings.Trim(line, "\t\n\r ")

		if len(line) == 0 || enumCommentLineRE.MatchString(line) {
			//fmt.Printf("Empty or comment line %s\n", line)
			continue
		}

		if f := funcRE.FindStringSubmatch(line); f != nil {
			currentFunction = new(Function)
			currentFunction.Name = f[1]
			currentFunction.Parameters = make([]Parameter, 0, 4)
		} else if ret := funcReturnRE.FindStringSubmatch(line); ret != nil {
			currentFunction.Return = ret[1]
		} else if param := funcParamRE.FindStringSubmatch(line); param != nil {
			out := false
			if param[3] == "out" {
				out = true
			}
			modifier := ParamModifierValue
			if param[4] == "array" {
				modifier = ParamModifierArray
			} else if param[4] == "reference" {
				modifier = ParamModifierReference
			}
			currentFunction.Parameters = append(currentFunction.Parameters, Parameter{param[1], param[2], out, modifier})
		} else if category := funcCategoryRE.FindStringSubmatch(line); category != nil {
			currentCategory = category[1]
			currentFunction.Category = currentCategory
			if functions[currentCategory] == nil {
				functions[currentCategory] = make([]*Function, 0, 4)
			}
			functions[currentCategory] = append(functions[currentCategory], currentFunction)
		} else if subcategory := funcSubCategoryRE.FindStringSubmatch(line); subcategory != nil {
			currentFunction.SubCategory = subcategory[1]
		} else if version := funcVersionRE.FindStringSubmatch(line); version != nil {
			v, err := MakeVersionFromMajorMinorString(version[1], version[2])
			if err != nil {
				return functions, finfo, fmt.Errorf("Unable to parse version: '%s'", line)
			}
			currentFunction.Version = v
		} else if deprecated := funcDeprecatedRE.FindStringSubmatch(line); deprecated != nil {
			v, err := MakeVersionFromMajorMinorString(deprecated[1], deprecated[2])
			if err != nil {
				return functions, finfo, fmt.Errorf("Unable to parse version: '%s'", line)
			}
			currentFunction.DeprecatedVersion = v
		} else if offset := funcOffsetRE.FindStringSubmatch(line); offset != nil {
			currentFunction.Offset = offset[1]
		} else if glxropcode := funcGlxRopCodeRE.FindStringSubmatch(line); glxropcode != nil {
			currentFunction.GlxRopCode = glxropcode[1]
		} else if glxsingle := funcGlxSingleRE.FindStringSubmatch(line); glxsingle != nil {
			currentFunction.GlxSingle = glxsingle[1]
		} else if wglflags := funcWglFlagsRE.FindStringSubmatch(line); wglflags != nil {
			currentFunction.WglFlags = wglflags[1]
		} else if glxflags := funcGlxFlagsRE.FindStringSubmatch(line); glxflags != nil {
			currentFunction.GlxFlags = glxflags[1]
		} else if dlflags := funcDlFlagsRE.FindStringSubmatch(line); dlflags != nil {
			currentFunction.DlFlags = dlflags[1]
		} else if glfflags := funcGlfFlagsRE.FindStringSubmatch(line); glfflags != nil {
			currentFunction.GlfFlags = glfflags[1]
		} else if glxvendorpriv := funcGlxVendorPrivRE.FindStringSubmatch(line); glxvendorpriv != nil {
			currentFunction.GlxVendorPriv = glxvendorpriv[1]
		} else if glextmask := funcGlextMaskRE.FindStringSubmatch(line); glextmask != nil {
			currentFunction.GlextMask = glextmask[1]
		} else if extension := funcExtensionRE.FindStringSubmatch(line); extension != nil {
			if len(extension) >= 2 {
				currentFunction.Extension = extension[1]
			}
		} else if vectorequiv := funcVectorEquivRE.FindStringSubmatch(line); vectorequiv != nil {
			currentFunction.VectorEquiv = vectorequiv[1]
		} else if glxvectorequiv := funcGlxVectorEquivRE.FindStringSubmatch(line); glxvectorequiv != nil {
			currentFunction.GlxVectorEquiv = glxvectorequiv[1]
		} else if alias := funcAliasRE.FindStringSubmatch(line); alias != nil {
			currentFunction.Alias = alias[1]
		} else if beginend := funcBeginEndRE.FindStringSubmatch(line); beginend != nil {
			currentFunction.BeginEnd = beginend[1]
		} else if allVersions := funcAllVersionsRE.FindStringSubmatch(line); allVersions != nil {
			split := strings.Split(allVersions[1], " ")
			for _, verString := range split {
				v, err := MakeVersionFromString(verString)
				if err != nil {
					return functions, finfo, fmt.Errorf("Unable to parse version: '%s'", line)
				}
				finfo.Versions = append(finfo.Versions, v)
			}
		} else if allDeprVersions := funcAllDeprVersionsRE.FindStringSubmatch(line); allDeprVersions != nil {
			split := strings.Split(allDeprVersions[1], " ")
			for _, verString := range split {
				v, err := MakeVersionFromString(verString)
				if err != nil {
					return functions, finfo, fmt.Errorf("Unable to parse version: '%s'", line)
				}
				finfo.DeprecatedVersions = append(finfo.DeprecatedVersions, v)
			}
		} else if allCategories := funcAllCategoriesRE.FindStringSubmatch(line); allCategories != nil {
			finfo.Categories = strings.Split(allCategories[1], " ")
		} else if passthru := funcPassthruRE.FindStringSubmatch(line); passthru != nil {
			if _, ok := finfo.Passthru[currentCategory]; !ok {
				finfo.Passthru[currentCategory] = make([]string, 0, 8)
			}
			finfo.Passthru[currentCategory] = append(finfo.Passthru[currentCategory], passthru[1])
		} else if newcategory := funcNewCategoryRE.FindStringSubmatch(line); newcategory != nil {
			currentCategory = newcategory[1]
			functions[currentCategory] = make([]*Function, 0)
			//} else if funcIgnoreRE.MatchString(line) {
			//	// ignore
		} else {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to parse line '%s' (Ignoring)\n", line)
			//return functions, finfo, fmt.Errorf("Unable to parse line: '%s'", line)
		}
	}

	// HACK: This is really ugly:
	// Some ARB extensions are now part of the GL core (starting from version 3.0).
	// Parse the passthru C comments for every "VERSION_*" category.
	// If we found an extension in the comment add them to a version category.
	for cat, pts := range finfo.Passthru {
		if strings.HasPrefix(cat, "VERSION") {
			for _, pt := range pts {
				strippedCat := strings.Trim(pt, "\t */")
				if foundFuncts, ok := functions[strippedCat]; ok {
					fmt.Printf("Adding extension %s to %s\n", strippedCat, cat)
					functions[cat] = append(functions[cat], foundFuncts...)
				}
			}
		}
	}

	return functions, finfo, nil
}
