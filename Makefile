linuxBuild= libunzip.so 
macBuild= libunzip.dylib 
winBuild= libunzip.dll 
headers= libunzip.h


output_target= ../assets/



all: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	cp $(linuxBuild) ../lib/linux/$(linuxBuild)
	
	go build -buildmode=c-archive main.go
	gcc -shared -pthread -o libunzip.dll main.c main.a -lWinMM -lntdll -lWS2_32
	cp $(winBuild)  ../windows/runner/$(winBuild)  

linux: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	cp $(linuxBuild) ../lib/linux/$(linuxBuild)

windows: main.go
	go build -buildmode=c-archive main.go
	gcc -shared -pthread -o libunzip.dll main.c main.a -lWinMM -lntdll -lWS2_32

clean: 
	rm main.a
	rm main.h
	rm main.so

build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	
	