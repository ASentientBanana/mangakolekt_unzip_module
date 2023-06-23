linuxBuild= libunzip.so 
macBuild= libunzip.dylib 
winBuild= libunzip.dll 
headers= libunzip.h

l_outputs= $(linuxBuild) $(headers)
m_outputs= $(macBuild) $(headers)
w_outputs= $(winBuild) $(headers)
output_target= ../assets/

all: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	cp ./$(headers) ../lib/linux/$(headers)
	cp $(linuxBuild) ../lib/linux/$(linuxBuild)
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	cp ./$(headers) ../macos/$(headers)
	cp $(macBuild)  ../macos/$(macBuild)  
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	cp ./$(headers) ../windows/runner/$(headers)
	cp $(winBuild)  ../windows/runner/$(winBuild)  

clean-all:
	rm  $(output_target)win/$(winBuild)  
	rm  $(output_target)linux/$(linuxBuild)
	rm  $(output_target)mac/$(macBuild)  


build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	
