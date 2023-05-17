package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"archive/zip"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	p "path/filepath"
	"strings"
	"unsafe"
)

// "time"

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

func getExt(name string) string {
	stringChunks := strings.Split(name, ".")
	return stringChunks[len(stringChunks)-1]
}

//export Unzip
func Unzip(_files *C.char, _path *C.char, _output *C.char) *C.char {
	results := []string{}
	files := []string{}

	output := C.GoString(_output)
	filesString := C.GoString(_files)
	// path := C.GoString(_path)

	files = strings.Split(filesString, "&&")

	for _, e := range files {
		nameID := uuid.New().String()
		archive, err := zip.OpenReader(e)
		if err != nil {
			break
		}
		defer archive.Close()
		for _, f := range archive.File {
			//This is to ignore if its a dir
			if f.FileInfo().IsDir() {
				continue
			}

			dstPath := p.Join(output, nameID+"."+getExt(f.Name))
			fmt.Println(dstPath)

			cf, errC := os.Create(dstPath)
			if errC != nil {
				fmt.Println(errC)
				continue
			}
			archivedFile, err := f.Open()
			if err != nil {
				continue
			}
			io.Copy(cf, archivedFile)
			if err != nil {
				continue
			}

			splitPath := strings.Split(e, "/")
			archiveName := splitPath[len(splitPath)-1]

			//template
			//cbz-name/cover-path/cbz-path
			results = append(results, archiveName+";"+dstPath+";"+e)
			cf.Close()
			archivedFile.Close()
			break
		}
	}
	return C.CString(strings.Join(results, "&?&"))
}

//export FreeStrings
func FreeStrings(str *C.char, count C.int) {
	C.free(unsafe.Pointer(str))
}

func main() {}
