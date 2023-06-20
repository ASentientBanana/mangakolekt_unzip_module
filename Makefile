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
	cp ./$(headers) $(output_target)linux/$(headers)
	cp $(linuxBuild)  $(output_target)linux/$(linuxBuild)
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	cp ./$(headers) $(output_target)mac/$(headers)
	cp $(macBuild)  $(output_target)mac/$(macBuild)  
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	cp ./$(headers) $(output_target)win/$(headers)
	cp $(winBuild)  $(output_target)win/$(winBuild)  

clean-all:
	rm  $(output_target)win/$(winBuild)  
	rm  $(output_target)linux/$(linuxBuild)
	rm  $(output_target)mac/$(macBuild)  


build: main.go
	go build -buildmode=c-shared -o  $(linuxBuild) main.go
	go build -buildmode=c-shared -o  $(macBuild)  main.go
	go build -buildmode=c-shared -o  $(winBuild)  main.go
	
