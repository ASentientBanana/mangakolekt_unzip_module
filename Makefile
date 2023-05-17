
linuxBuild= libunzip.so 
outputs= $(linuxBuild) libunzip.h
output_target= ../lib/ffi/

build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	cp $(linuxBuild) libunzip.h $(output_target)

clean: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	rm ../lib/ffi/libunzip.so 
	cp $(outputs) $(output_target)
