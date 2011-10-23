package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"regexp"
)

// TODO: better regexps
var (
	enumEmptyLineRE  = regexp.MustCompile("^[ \\t]*(#.*)?\\n")
	enumCategoryRE   = regexp.MustCompile("^[ \\t]*([_0-9A-Za-z]+)[ \\t]+enum:")
	enumRE           = regexp.MustCompile("^[ \\t]*([_0-9A-Za-z]+)[ \\t]*=[ \\t]*([^ \\t#\\n]*)")
	enumPassthruRE   = regexp.MustCompile("^[ \\t]*passthru:")
	enumUseRE        = regexp.MustCompile("^[ \\t]*use[ \\t]+([_0-9A-Za-z]+)[ \\t]+([_0-9A-Za-z]+)")
	enumVersionRE    = regexp.MustCompile("^VERSION_([0-9]+)_([0-9]+)$")
	enumVersionDepRE = regexp.MustCompile("^VERSION_([0-9]+)_([0-9]+)_DEPRECATED$")
)

func ReadEnumsFromFile(name string) (EnumCategories, os.Error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadEnums(file)
}

func ReadEnums(r io.Reader) (EnumCategories, os.Error) {
	categories := make(EnumCategories)
	br := bufio.NewReader(r)
	currentCategory := ""
	for line, rerr := br.ReadString('\n'); rerr == nil || rerr == os.EOF; line, rerr = br.ReadString('\n') {
		if rerr == os.EOF {
			line += "\n"
		}
		//fmt.Printf(line)
		if !enumEmptyLineRE.MatchString(line) && !enumPassthruRE.MatchString(line) {
			if category := enumCategoryRE.FindStringSubmatch(line); category != nil {
				//fmt.Printf("%v\n", category[1])
				currentCategory = category[1]
				categories[currentCategory] = make(Enums)
			} else if enum := enumRE.FindStringSubmatch(line); enum != nil {
				//fmt.Printf("%v %v\n", enum[1], enum[2])
				categories[currentCategory][enum[1]] = Enum{enum[1], enum[2]}
			} else if use := enumUseRE.FindStringSubmatch(line); use != nil {
				//fmt.Printf("%v %v\n", use[1], use[2])
				categories[currentCategory][use[2]] = categories[use[1]][use[2]]
			} else {
				//return os.NewError("Unable to parse line: '" + line + "'")
				fmt.Fprintf(os.Stderr, "Unable to parse line: "+line)
			}
		}
		if rerr == os.EOF {
			break
		}
	}
	// TODO: update reused enums
	return categories, nil
}
