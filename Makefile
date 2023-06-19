linuxBuild= libunzip.so 
macBuild= libunzip.dylib 
winBuild= libunzip.dll 
headers= libunzip.h

l_outputs= $(linuxBuild) $(headers)
m_outputs= $(macBuild) $(headers)
w_outputs= $(winBuild) $(headers)
output_target= ../lib/ffi/

clean: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	rm  $(output_target)$(linuxBuild)
	rm  $(output_target)$(macBuild)  
	rm  $(output_target)$(winBuild)  
	cp ./$(headers) $(output_target)$(headers)
	cp $(linuxBuild)  $(output_target)$(linuxBuild)
	cp $(macBuild)  $(output_target)$(macBuild)  
	cp $(winBuild)  $(output_target)$(winBuild)  



build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	
