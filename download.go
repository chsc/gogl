package main

import (
	"fmt"
	"os"
	"http"
	"io/ioutil"
)

const (
	RegistryBaseURL       = "http://www.opengl.org/registry/api"
	OpenGLEnumSpecFile    = "enum.spec"
	OpenGLEnumExtSpecFile = "enumext.spec"
	OpenGLSpecFile        = "gl.spec"
	OpenGLTypeMapFile     = "gl.tm"
)

func makeURL(base, file string) string {
	return fmt.Sprintf("%s/%s", base, file)
}

func DownloadFile(baseURL, fileName string) os.Error {
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
	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func DownloadOpenGLSpecs() {
	// DownloadFile(RegistryBaseURL, OpenGLEnumSpecFile)
	DownloadFile(RegistryBaseURL, OpenGLEnumExtSpecFile)
	DownloadFile(RegistryBaseURL, OpenGLSpecFile)
	DownloadFile(RegistryBaseURL, OpenGLTypeMapFile)
}

// TODO: download wgl, glx specs ...

