


build: main.go
	go build -buildmode=c-shared -o libunzip.so main.go
