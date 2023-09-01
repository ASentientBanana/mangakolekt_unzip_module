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

clean:
	rm  $(output_target)win/$(winBuild)  
	rm  $(output_target)linux/$(linuxBuild)
	rm  $(output_target)mac/$(macBuild)  


build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	
	