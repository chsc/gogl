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
	tmEmptyOrCommentRE = regexp.MustCompile("^[ \\t]*(#.*)?$")
	tmTypePairRE       = regexp.MustCompile("^([_A-Za-z0-9]+),\\*,\\*,[\\t ]*([A-Za-z0-9\\*_ ]+),\\*,\\*")
)


func ReadTypeMapFromFiles(files []string) (TypeMap, error) {
	allTypes := make(TypeMap)
	for _, file := range files {
		types, err := ReadTypeMapFromFile(file)
		if err != nil {
			return nil, err
		}
		for k, v := range types {
			allTypes[k] = v
		}
	}
	return allTypes, nil
}

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

		if tmEmptyOrCommentRE.MatchString(line) {
			continue
		}

		if typePair := tmTypePairRE.FindStringSubmatch(line); typePair != nil {
			tm[typePair[1]] = typePair[2]
		} else {
			return tm, fmt.Errorf("Unable to parse line: '%v'", line)
		}
	}

	return tm, nil
}
