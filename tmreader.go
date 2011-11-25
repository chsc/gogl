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
	tmEmptyOrCommentRE = regexp.MustCompile("^[ \\t]*(#.*)?$")
	tmTypePairRE       = regexp.MustCompile("^([A-Za-z0-9]+),\\*,\\*,[\\t ]*([A-Za-z0-9]+),\\*,\\*")
)

func ReadTypeMapFromFile(name string) (TypeMap, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadTypeMap(file)
}

func ReadTypeMap(r io.Reader) (TypeMap, error) {
	tm := make(TypeMap)
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\n")
		//fmt.Println(line)

		if tmEmptyOrCommentRE.MatchString(line) {
			// skip
		} else if typePair := tmTypePairRE.FindStringSubmatch(line); typePair != nil {
			tm[typePair[1]] = typePair[2]
		} else {
			fmt.Fprintf(os.Stderr, "Unable to parse line: %v", line)
		}
	}

	return tm, nil
}
