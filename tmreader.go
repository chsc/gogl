package main

import (
	"os"
	"io"
	"bufio"
	"regexp"
	"fmt"
)

var (
	tmEmptyRE    = regexp.MustCompile("^\n")
	tmCommentRE  = regexp.MustCompile("^#.*")
	tmTypePairRE = regexp.MustCompile("^(.+),\\*,\\*,[\\t ]*(.+),\\*,\\*")
)

func ReadTypeMapFromFile(name string) (TypeMap, os.Error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadTypeMap(file)
}

func ReadTypeMap(r io.Reader) (TypeMap, os.Error) {
	tm := make(TypeMap)
	br := bufio.NewReader(r)
	for line, rerr := br.ReadString('\n'); rerr == nil; line, rerr = br.ReadString('\n') {
		if tmEmptyRE.MatchString(line) {
			continue
		}
		if tmCommentRE.FindStringSubmatch(line) != nil { // use matchstring?
			continue
		}
		if typePair := tmTypePairRE.FindStringSubmatch(line); typePair != nil {
			tm[typePair[1]] = typePair[2]
			continue
		}
		fmt.Fprintf(os.Stderr, "Unable to parse line: "+line)
	}
	return tm, nil
}
