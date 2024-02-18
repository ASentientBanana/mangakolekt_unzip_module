linuxBuild= manga_archive.so 
macBuild= manga_archive.dylib 
winBuild= manga_archive.dll 


linux: main.go
	echo "Started $(arch)"
	mkdir -p build/linux/$(arch)
	GOARCH=$(arch) go build -buildmode=c-shared -o  ./build/linux/$(arch)/$(linuxBuild) ./main.go 

windows: main.go
	go build -buildmode=c-archive main.go
	gcc -shared -pthread -o $(winBuild) main.c main.a -lWinMM -lntdll -lWS2_32


windows-ta:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o $(winBuild) main.go

build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	
	