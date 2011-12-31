include $(GOROOT)/src/Make.inc
 
TARG=gogl
GOFILES=\
	main.go\
	download.go\
	structs.go\
	enumreader.go\
	funcreader.go\
	tmreader.go\
	group.go\
	generator.go\
	util.go
 
include $(GOROOT)/src/Make.cmd

format:
	gofmt -w *.go

download_khronos:
	./gogl -download

# TODO: Get rid of revision number. How to get the newest revision from bitbucket? 
download_alfonse:
	./gogl -download -url=https://bitbucket.org/alfonse/gl-xml-specs/raw/5f8ea57530c2/glspecs -dir=alfonse_specs 

gen_khronos:
	./gogl

gen_alfonse: 
	./gogl -dir=alfonse_specs
