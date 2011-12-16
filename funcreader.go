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
	funcEmptyLineRE       = regexp.MustCompile("^\\s*(#.*)?$")
	funcRE                = regexp.MustCompile("^(\\w+)\\([\\w\\s,]*\\)")
	funcReturnRE          = regexp.MustCompile("^\\s+return\\s+(\\w+)")
	funcParamRE           = regexp.MustCompile("^\\s+param\\s+(\\w+)\\s+(\\w+)\\s+(in|out)\\s+(array|value|reference)")
	funcVectorEquivRE     = regexp.MustCompile("^\\s+vectorequiv\\s+(\\w+)")
	funcGlxVectorEquivRE  = regexp.MustCompile("^\\s+glxvectorequiv\\s+(\\w+)")
	funcCategoryRE        = regexp.MustCompile("^\\s+category\\s+(\\w+)")
	funcSubCategoryRE     = regexp.MustCompile("^\\s+subcategory\\s+(\\w+)")
	funcVersionRE         = regexp.MustCompile("^\\s+version\\s+([0-9]+)\\.([0-9]+)")
	funcDeprecatedRE      = regexp.MustCompile("^\\s+deprecated\\s+([0-9]+)\\.([0-9]+)")
	funcOffsetRE          = regexp.MustCompile("^\\s+offset(\\s+([0-9\\*\\?]+))?")
	funcGlxRopCodeRE      = regexp.MustCompile("^\\s+glxropcode(\\s+([0-9\\*\\?]+))?")
	funcGlxSingleRE       = regexp.MustCompile("^\\s+glxsingle\\s+([0-9\\?\\*]+)")
	funcWglFlagsRE        = regexp.MustCompile("^\\s+wglflags(\\s+([\\w\\*\\- ]+))?")
	funcGlxFlagsRE        = regexp.MustCompile("^\\s+glxflags(\\s+([\\w\\*\\- ]+))?")
	funcDlFlagsRE         = regexp.MustCompile("^\\s+dlflags(\\s+([\\w\\*\\- ]+))?")
	funcGlfFlagsRE        = regexp.MustCompile("^\\s+glfflags(\\s+([\\w\\*\\- ]+))?")
	funcGlxVendorPrivRE   = regexp.MustCompile("^\\s+glxvendorpriv\\s+([0-9\\?]+)")
	funcGlextMaskRE       = regexp.MustCompile("^\\s+glextmask(\\s+(\\w+))?")
	funcExtensionRE       = regexp.MustCompile("^\\s+extension(\\s+([\\w\\- ]+))?")
	funcAliasRE           = regexp.MustCompile("^\\s+alias\\s+(\\w+)")
	funcBeginEndRE        = regexp.MustCompile("^\\s+beginend\\s+([\\w\\*\\-]+)")
	funcAllCategoriesRE   = regexp.MustCompile("^category:\\s*([\\w ]+)")
	funcAllVersionsRE     = regexp.MustCompile("^version:\\s*([0-9\\. ]+)")
	funcAllDeprVersionsRE = regexp.MustCompile("^deprecated:\\s*([0-9\\. ]+)")
	funcPassthruRE        = regexp.MustCompile("^passthru:\\s*(.*)")
	funcNewCategoryRE     = regexp.MustCompile("^newcategory:\\s*(\\w+)")
	funcIgnoreRE          = regexp.MustCompile("^[a-z\\-]+:.*")
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
	functions := make(FunctionCategories)
	finfo := &FunctionsInfo{make([]Version, 0, 8), make([]Version, 0, 8), make([]string, 0, 8), ""}

	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}
		line = strings.TrimRight(line, "\n")

		if funcEmptyLineRE.MatchString(line) {
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
			array := false
			if param[4] == "array" {
				array = true
			} // TODO: reference as value?
			currentFunction.Parameters = append(currentFunction.Parameters, Parameter{param[1], param[2], out, array})
		} else if category := funcCategoryRE.FindStringSubmatch(line); category != nil {
			currentFunction.Category = category[1]
			if functions[category[1]] == nil {
				functions[category[1]] = make([]*Function, 0, 4)
			}
			functions[category[1]] = append(functions[category[1]], currentFunction)
		} else if subcategory := funcSubCategoryRE.FindStringSubmatch(line); subcategory != nil {
			currentFunction.Category = subcategory[1]
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
			currentFunction.Offset = offset[2]
		} else if glxropcode := funcGlxRopCodeRE.FindStringSubmatch(line); glxropcode != nil {
			currentFunction.GlxRopCode = glxropcode[2]
		} else if glxsingle := funcGlxSingleRE.FindStringSubmatch(line); glxsingle != nil {
			currentFunction.GlxSingle = glxsingle[1]
		} else if wglflags := funcWglFlagsRE.FindStringSubmatch(line); wglflags != nil {
			currentFunction.WglFlags = wglflags[2]
		} else if glxflags := funcGlxFlagsRE.FindStringSubmatch(line); glxflags != nil {
			currentFunction.GlxFlags = glxflags[2]
		} else if dlflags := funcDlFlagsRE.FindStringSubmatch(line); dlflags != nil {
			currentFunction.DlFlags = dlflags[2]
		} else if glfflags := funcGlfFlagsRE.FindStringSubmatch(line); glfflags != nil {
			currentFunction.GlfFlags = glfflags[2]
		} else if glxvendorpriv := funcGlxVendorPrivRE.FindStringSubmatch(line); glxvendorpriv != nil {
			currentFunction.GlxVendorPriv = glxvendorpriv[1]
		} else if glextmask := funcGlextMaskRE.FindStringSubmatch(line); glextmask != nil {
			currentFunction.GlextMask = glextmask[2]
		} else if extension := funcExtensionRE.FindStringSubmatch(line); extension != nil {
			currentFunction.Extension = extension[2]
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
		} else if passthru := funcPassthruRE.FindStringSubmatch(line); passthru != nil{
			finfo.Passthru += passthru[1] + "\n"
		} else if newcategory := funcNewCategoryRE.FindStringSubmatch(line); newcategory != nil {
			functions[newcategory[1]] = make([]*Function, 0)
		} else if funcIgnoreRE.MatchString(line) {
			// ignore
		} else {
			return functions, finfo, fmt.Errorf("Unable to parse line: '%s'", line)
		}
	}

	return functions, finfo, nil
}
