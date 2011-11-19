include $(GOROOT)/src/Make.inc
 
TARG=GoGL
GOFILES=main.go download.go structs.go enumreader.go funcreader.go tmreader.go sort.go generator.go
 
include $(GOROOT)/src/Make.cmd

format:
	gofmt -w *.go

generate:
	GoGL
