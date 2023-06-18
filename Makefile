
linuxBuild= libunzip.so 
outputs= $(linuxBuild) libunzip.h
output_target= ../lib/ffi/
rm_target=../lib/ffi/libunzip.so

build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	cp $(linuxBuild) libunzip.h $(output_target)

clean: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	rm $(rm_target)
	cp $(outputs) $(output_target)


