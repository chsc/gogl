package main

import (
	"os"
	"io"
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
	// TODO: Implement type map reader
	return nil, nil
}

