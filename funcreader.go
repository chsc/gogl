package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var (
	funcEmptyLineRE    = regexp.MustCompile("^[ \\t]*(#.*)?\\n")
	funcRE             = regexp.MustCompile("^[ \\t]*([_0-9A-Za-z]+)\\(.*\\)")
	funcReturnRE       = regexp.MustCompile("^[ \\t]*return[ \\t]+([^ \\t#\\n]+)")
	funcParamInValueRE = regexp.MustCompile("^[ \\t]*param[ \\t]+([^ \\t]+)[ \\t]+([^ \\t]+)[ \\t]+in[ \\t]+value")
	funcParamInArrayRE = regexp.MustCompile("^[ \\t]*param[ \\t]+([^ \\t]+)[ \\t]+([^ \\t]+)[ \\t]+in[ \\t]+array")
	funcCategoryRE     = regexp.MustCompile("^[ \\t]*category[ \\t]+([^ \\t#\\n]+)[ \\t]*(#.*)?$")
	funcVersionRE      = regexp.MustCompile("^[ \\t]*version[ \\t]+([0-9]+)\\.([0-9]+)")
	funcDeprecatedRE   = regexp.MustCompile("^[ \\t]*deprecated[ \\t]+([0-9]+)\\.([0-9]+)")
)

func ReadFunctionsFromFile(name string) (Functions, os.Error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadFunctions(file)
}

func ReadFunctions(r io.Reader) (Functions, os.Error) {
	functions := make(Functions)
	br := bufio.NewReader(r)
	var currentFunction *Function = nil
	for line, rerr := br.ReadString('\n'); rerr == nil || rerr == os.EOF; line, rerr = br.ReadString('\n') {
		if rerr == os.EOF {
			line += "\n"
		}
		//fmt.Printf(line)
		if !funcEmptyLineRE.MatchString(line) {
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
				major, _ := strconv.Atoi(version[1])
				minor, _ := strconv.Atoi(version[2])
				currentFunction.Version = Version{major, minor}
			} else if deprecated := funcDeprecatedRE.FindStringSubmatch(line); deprecated != nil {
				major, _ := strconv.Atoi(deprecated[1])
				minor, _ := strconv.Atoi(deprecated[2])
				currentFunction.DeprecatedVersion = Version{major, minor}
			} else {
				//return os.NewError("Unable to parse line: '" + line + "'")
				fmt.Fprintf(os.Stderr, "Unable to parse line: "+line)
			}
		}
		if rerr == os.EOF {
			break
		}
	}

	return functions, nil
}
