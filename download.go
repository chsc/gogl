// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	KhronosRegistryBaseURL = "http://www.opengl.org/registry/api"
	OpenGLEnumSpecFile     = "enum.spec"
	OpenGLEnumExtSpecFile  = "enumext.spec"
	OpenGLSpecFile         = "gl.spec"
	OpenGLTypeMapFile      = "gl.tm"
)

func makeURL(base, file string) string {
	return fmt.Sprintf("%s/%s", base, file)
}

func DownloadFile(baseURL, fileName, outDir string) error {
	fullURL := makeURL(baseURL, fileName)
	fmt.Printf("Downloading %s ...\n", fullURL)
	r, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	absPath, err := filepath.Abs(outDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(absPath, 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(absPath, fileName), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func DownloadOpenGLSpecs(baseURL, outDir string) {
	DownloadFile(baseURL, OpenGLEnumExtSpecFile, outDir)
	DownloadFile(baseURL, OpenGLSpecFile, outDir)
	DownloadFile(baseURL, OpenGLTypeMapFile, outDir)
}

// TODO: download wgl, glx specs ...
