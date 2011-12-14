// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	funcEmptyLineRE       = regexp.MustCompile("^[ \\t]*(#.*)?$")
	funcRE                = regexp.MustCompile("^[ \\t]*([_0-9A-Za-z]+)\\([0-9A-Za-z, _]*\\)")
	funcReturnRE          = regexp.MustCompile("^[ \\t]+return[ \\t]+([^ \\t#\\n]+)")
	funcParamInValueRE    = regexp.MustCompile("^[ \\t]+param[ \\t]+([^ \\t]+)[ \\t]+([^ \\t]+)[ \\t]+in[ \\t]+value")
	funcParamInArrayRE    = regexp.MustCompile("^[ \\t]+param[ \\t]+([^ \\t]+)[ \\t]+([^ \\t]+)[ \\t]+in[ \\t]+array")
	funcCategoryRE        = regexp.MustCompile("^[ \\t]+category[ \\t]+([^ \\t#\\n]+)")
	funcVersionRE         = regexp.MustCompile("^[ \\t]+version[ \\t]+([0-9]+)\\.([0-9]+)")
	funcDeprecatedRE      = regexp.MustCompile("^[ \\t]+deprecated[ \\t]+([0-9]+)\\.([0-9]+)")
	funcOffsetRE          = regexp.MustCompile("^[ \\t]+offset[ \\t]+([0-9\\*\\?]+)")
	funcGlxRopCodeRE      = regexp.MustCompile("^[ \\t]+glxropcode[ \\t]+([0-9\\*\\?]+)")
	funcGlxSingleRE       = regexp.MustCompile("^[ \\t]+glxsingle[ \\t]+([0-9\\?\\*]+)")
	funcWglFlagsRE        = regexp.MustCompile("^[ \\t]+wglflags[ \\t]+([0-9A-Za-z\\- ]+)")
	funcGlxFlagsRE        = regexp.MustCompile("^[ \\t]+glxflags[ \\t]+([0-9A-Za-z\\- ]+)")
	funcDlFlagsRE         = regexp.MustCompile("^[ \\t]+dlflags[ \\t]+([0-9A-Za-z\\- ]+)")
	funcGlfFlagsRE        = regexp.MustCompile("^[ \\t]+glfflags[ \\t]+([0-9A-Za-z\\*\\- ]+)")
	funcExtensionRE       = regexp.MustCompile("^[ \\t]+extension([ \\t]+([0-9A-Za-z\\- ]+))?")
	funcAllVersionsRE     = regexp.MustCompile("^version:[ \\t]*([0-9\\. ]+)")
	funcAllDeprVersionsRE = regexp.MustCompile("^deprecated:[ \\t]*([0-9\\. ]+)")
	funcPassthruRE        = regexp.MustCompile("^passthru:[ \\t]*(.*)")
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
	finfo := &FunctionsInfo{make([]Version, 0, 8), make([]Version, 0, 8), ""}

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
		} else if param := funcParamInValueRE.FindStringSubmatch(line); param != nil {
			currentFunction.Parameters = append(currentFunction.Parameters, Parameter{param[1], param[2], false})
		} else if param := funcParamInArrayRE.FindStringSubmatch(line); param != nil {
			currentFunction.Parameters = append(currentFunction.Parameters, Parameter{param[1], param[2], true})
		} else if category := funcCategoryRE.FindStringSubmatch(line); category != nil {
			currentFunction.Category = category[1]
			if functions[category[1]] == nil {
				functions[category[1]] = make([]*Function, 0, 4)
			}
			functions[category[1]] = append(functions[category[1]], currentFunction)
		} else if version := funcVersionRE.FindStringSubmatch(line); version != nil {
			v, err := MakeVersionFromMinorMajorString(version[1], version[2])
			if err != nil {
				return functions, finfo, errors.New("Unable to parse version: " + line)
			}
			currentFunction.Version = v
		} else if deprecated := funcDeprecatedRE.FindStringSubmatch(line); deprecated != nil {
			v, err := MakeVersionFromMinorMajorString(deprecated[1], deprecated[2])
			if err != nil {
				return functions, finfo, errors.New("Unable to parse version: " + line)
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
		} else if extension := funcExtensionRE.FindStringSubmatch(line); extension != nil {
			currentFunction.Extension = extension[1]
		} else if allVersions := funcAllVersionsRE.FindStringSubmatch(line); allVersions != nil {
			split := strings.Split(allVersions[1], " ")
			for _, verString := range split {
				v, err := MakeVersionFromString(verString)
				if err != nil {
					return functions, finfo, errors.New("Unable to parse version: " + line)
				}
				finfo.Versions = append(finfo.Versions, v)
			}
		} else if allDeprVersions := funcAllDeprVersionsRE.FindStringSubmatch(line); allDeprVersions != nil {
			split := strings.Split(allDeprVersions[1], " ")
			for _, verString := range split {
				v, err := MakeVersionFromString(verString)
				if err != nil {
					return functions, finfo, errors.New("Unable to parse version: " + line)
				}
				finfo.DeprecatedVersions = append(finfo.DeprecatedVersions, v)
			}
		} else if passthru := funcPassthruRE.FindStringSubmatch(line); passthru != nil{
			finfo.Passthru += passthru[1] + "\n"
		} else {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to parse line: %v\n", line)
		}
	}
	return functions, finfo, nil
}
