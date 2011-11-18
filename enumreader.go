package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"regexp"
	"strings"
)

// TODO: better regexps
var (
	enumEmptyLineRE  = regexp.MustCompile("^[ \\t]*(#.*)?$")
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
	deferredUseEnums := make(map[string]map[string]string)
	currentCategory := ""
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err == os.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\n")

		if enumEmptyLineRE.MatchString(line) || enumPassthruRE.MatchString(line) {
			// skip
		} else if category := enumCategoryRE.FindStringSubmatch(line); category != nil {
			//fmt.Printf("%v\n", category[1])
			currentCategory = category[1]
			categories[currentCategory] = make(Enums)
		} else if enum := enumRE.FindStringSubmatch(line); enum != nil {
			//fmt.Printf("%v %v\n", enum[1], enum[2])
			categories[currentCategory][enum[1]] = Enum{enum[1], enum[2]}
		} else if use := enumUseRE.FindStringSubmatch(line); use != nil {
			//fmt.Printf("%v %v\n", use[1], use[2])
			if deferredUseEnums[currentCategory] == nil {
				deferredUseEnums[currentCategory] = make(map[string]string)
			}
			deferredUseEnums[currentCategory][use[2]] = use[1]
		} else {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to parse line: "+line)
		}
	}

	for category, enums := range deferredUseEnums {
		for name, referencedCategory := range enums {
			if dereference, ok := categories[referencedCategory][name]; ok {
				categories[category][name] = dereference
			} else if dereference, ok := categories[referencedCategory+"_DEPRECATED"][name]; ok {
				categories[category][name] = dereference
			} else {
				fmt.Fprintf(os.Stderr, "WARNING: Failed to dereference %v: \"use %v %v\"\n", category, referencedCategory, name)
			}
		}
	}

	return categories, nil
}
