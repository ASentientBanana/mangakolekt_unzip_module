package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"os"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/asentientbanana/uz/unzip"
	"github.com/asentientbanana/uz/util"
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

// Function to check if a file has an image extension
//
//export Unzip_Single_book
func Unzip_Single_book(_filePath *C.char, _dest *C.char) *C.char {
	zipPath := C.GoString(_filePath)
	dest := C.GoString(_dest)
	content := unzip.Unzip_Single_Book(zipPath, dest)
	return C.CString(strings.Join(content, "?&?"))
}

//export Check_For_Lib_dir
func Check_For_Lib_dir(_path *C.char) {
	path := C.GoString(_path)
	util.CheckForLibDir(path)
}

//export Get_Files_From_Dir
func Get_Files_From_Dir(_path *C.char) *C.char {
	dirPath := C.GoString(_path)
	filesString := util.GetFilesFromDir(dirPath)
	return C.CString(filesString)
}

func main() {}
