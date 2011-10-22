package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"regexp"
)

//TODO: better regexps
var (
	emptyRE        = regexp.MustCompile("^\n")
	commentRE      = regexp.MustCompile("^#.*")
	categoryRE     = regexp.MustCompile("^([_0-9A-Za-z]+) enum:")
	enumRE         = regexp.MustCompile("^[ \\t]+([_0-9A-Za-z]+)[ \\t]*=[ \\t]*([^ \\t#\n]*)")
	passthruRE     = regexp.MustCompile("^passthru:")
	useRE          = regexp.MustCompile("^[ \\t]+use ([_0-9A-Za-z]+)[ \\t]+([_0-9A-Za-z]+)")
	versionRE      = regexp.MustCompile("^VERSION_([0-9]+)_([0-9]+)$")
	versionDepRE   = regexp.MustCompile("^VERSION_([0-9]+)_([0-9]+)_DEPRECATED$")
)

func ReadEnumsFromFile(name string) (Enums, os.Error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadEnums(file)
}

func ReadEnums(r io.Reader) (Enums, os.Error) {
	enums := make(Enums)
	br := bufio.NewReader(r)
	currentCategory := ""
	for line, rerr := br.ReadString('\n'); rerr == nil; line, rerr = br.ReadString('\n') {
		//fmt.Printf("-%v-\n", line)
		if emptyRE.MatchString(line) {
			continue
		}
		if commentRE.FindStringSubmatch(line) != nil { // use matchstring?
			continue
		}
		if category := categoryRE.FindStringSubmatch(line); category != nil {
			//fmt.Printf("%v\n", category[1])
			currentCategory = category[1]
			enums[currentCategory] = make([]Enum, 0, 4)
			continue
		}
		if enum := enumRE.FindStringSubmatch(line); enum != nil {
			//fmt.Printf("%v %v\n", enum[1], enum[2])
			enums[currentCategory] = append(enums[currentCategory], Enum{enum[1], enum[2]})
			continue
		}
		if passthruRE.MatchString(line) {
			continue
		}
		if use := useRE.FindStringSubmatch(line); use != nil {
			//fmt.Printf("%v %v\n", use[1], use[2])
			enums[currentCategory] = append(enums[currentCategory], Enum{use[2], use[1]})
			continue
		}
		//return os.NewError("Unable to parse line: '" + line + "'")
		fmt.Fprintf(os.Stderr, "Unable to parse line: " + line)
	}
	// TODO: update reused enums
	// TODO: parse problematic spec parts
	return enums, nil
}

