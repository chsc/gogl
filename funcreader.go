package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	fEmptyRE    = regexp.MustCompile("^\n")
	fCommentRE  = regexp.MustCompile("^#.*")
	fCategoryRE = regexp.MustCompile("^[\t ]+category[ \t]+([_0-9A-Za-z]+)")
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
	currentCategory := ""
	for line, rerr := br.ReadString('\n'); rerr == nil; line, rerr = br.ReadString('\n') {
		//fmt.Printf("-%v-\n", line)
		if fEmptyRE.MatchString(line) {
			continue
		}
		if fCommentRE.FindStringSubmatch(line) != nil { // use matchstring?
			continue
		}
		if category := fCategoryRE.FindStringSubmatch(line); category != nil {
			//fmt.Printf("%v\n", category[1])
			currentCategory = category[1]
			functions[currentCategory] = make([]Function, 0, 4)
			continue
		}
		// TODO: implement regexp for return, param, version, category, ...
		fmt.Fprintf(os.Stderr, "Unable to parse line: " + line)
	}

	return functions, nil
}

