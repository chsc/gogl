include $(GOROOT)/src/Make.inc
 
TARG=gogl
GOFILES=main.go download.go structs.go enumreader.go funcreader.go tmreader.go sort.go generator.go
 
include $(GOROOT)/src/Make.cmd

format:
	gofmt -w *.go

# TODO: add pregenerated OpenGL packages ...

gen21:
	gogl -version=2.1 -out=gl21/gl.go

gen31:
	gogl -version=3.1 -deprecation -out=gl31/gl.go

gen31c:
	gogl -version=3.1 -out=gl31c/gl.go
