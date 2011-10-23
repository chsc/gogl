package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	funcEmptyLineRE    = regexp.MustCompile("^[ \\t]*(#.*)?$")
	funcRE             = regexp.MustCompile("^[ \\t]*([_0-9A-Za-z]+)\\(([^\\)]*)\\)[ \\t]*$")
	funcReturnRE       = regexp.MustCompile("^[ \\t]*return[ \\t]+([^ \\t#\n]+)[ \\t]*(#.*)?$")
	funcParamInValueRE = regexp.MustCompile("^[ \\t]*param[ \\t]+([^ \\t]+)[ \\t]+in[ \\t]+value[ \\t]*(#.*)?$")
	funcParamInArrayRE = regexp.MustCompile("^[ \\t]*param[ \\t]+([^ \\t]+)[ \\t]+in[ \\t]+array[ \\t]*\\[.*\\][ \\t]*(#.*)?$")
	funcCategoryRE     = regexp.MustCompile("^[ \\t]*category[ \\t]+([^ \\t#\n]+)[ \\t]*(#.*)?$")
	funcVersionRE      = regexp.MustCompile("^[ \\t]*version[ \\t]+([^ \\t#\n]+)[ \\t]*(#.*)?$")
	funcDeprecatedRE   = regexp.MustCompile("^[ \\t]*deprecated[ \\t]+([^ \\t#\n]+)[ \\t]*(#.*)?$")
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
	for line, rerr := br.ReadString('\n'); rerr == nil || rerr == os.EOF; line, rerr = br.ReadString('\n') {
		if rerr == os.EOF {
			line += "\n"
		}
		//fmt.Printf("-%v-\n", line)
		if !funcEmptyLineRE.MatchString(line) {
			//if {
			//} else {
			//return os.NewError("Unable to parse line: '" + line + "'")
			fmt.Fprintf(os.Stderr, "Unable to parse line: "+line)
			//}
		}
		if rerr == os.EOF {
			break
		}
	}

	return functions, nil
}
