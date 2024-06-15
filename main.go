package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

*/
import "C"

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/asentientbanana/uz/unzip"
)

func ReadDir(path string) ([]string, error) {
	dir, err := os.ReadDir(path)
	files := []string{}
	if err != nil {
		return nil, err
	}
	for _, e := range dir {
		files = append(files, path+"/"+e.Name())
	}
	return files, nil
}

//export Unzip_Covers
func Unzip_Covers(_files *C.char, _path *C.char, _output *C.char) *C.char {

	output := C.GoString(_output)
	filesString := C.GoString(_files)

	results := unzip.Unzip_covers_from_dir(filesString, output)

	return C.CString(results)
}

//export Unzip_Single_book
func Unzip_Single_book(_filePath *C.char, _dest *C.char) *C.char {

	zipPath := C.GoString(_filePath)
	dest := C.GoString(_dest)

	content := unzip.Unzip_Single_Book(zipPath, dest)

	return C.CString(content)

}

func main() {
	// _unzipSingle("/home/petar/bigboy/Manga/OnePiece/", "/home/petar/Documents/mangakolekt/current")
	// Unzip_Single_book()
}
