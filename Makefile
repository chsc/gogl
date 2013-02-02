format:
	go fmt *.go

download_khronos:
	./gogl -download

download_alfonse:
	./gogl -download -url=https://bitbucket.org/alfonse/gl-xml-specs/raw/tip/glspecs -dir=alfonse_specs 

gen_khronos:
	./gogl

gen_alfonse: 
	./gogl -dir=alfonse_specs

# Just test go install
install_all:
	go install ./gl21
#	go install ./gl30
	go install ./gl31
#	go install ./gl31c
#	go install ./gl32
#	go install ./gl32c
#	go install ./gl32
#	go install ./gl32c
	go install ./gl33
#	go install ./gl33c
#	go install ./gl40
#	go install ./gl41c
	go install ./gl42
#	go install ./gl42c
	go install ./gl43
	go install ./arb
	go install ./ext
	go install ./ati
	go install ./amd
	go install ./nv
#	go install ./glx
#	go install ./wgl
