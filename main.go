package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

__attribute__((weak))
 void my_memcpy_wrapper(void* dest, const void* src, size_t size) {
     memcpy(dest, src, size);
 }

*/
import "C"

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	p "path/filepath"
	"strconv"
	"strings"
	"unsafe"

	"github.com/google/uuid"
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

func removeAllContents(dirPath string) error {
	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	// Read the directory entries
	entries, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	// Remove each file or directory within the directory
	for _, entry := range entries {
		entryPath := p.Join(dirPath, entry.Name())
		err = os.RemoveAll(entryPath)
		if err != nil {
			return err
		}
	}
	return nil
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

			dstPath := p.Join(output, nameID+"."+filepath.Ext(f.Name))
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

// Function to check if a file has an image extension
func isImageFile(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

//export Unzip_Single_book
func Unzip_Single_book(_filePath *C.char, _dest *C.char) {

	// Convert C string to Go string
	zipPath := C.GoString(_filePath)
	dest := C.GoString(_dest)

	removeAllContents(dest)

	// Open the zip file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		fmt.Println("Failed to open zip file:", err)

	}
	defer r.Close()

	// Iterate over each file in the zip archive
	fmt.Println("Number of files in a book: ", len(r.File))
	for i, file := range r.File {

		if file.FileInfo().IsDir() {
			continue
		}
		// Check if the file is an image (you can modify this condition as needed)
		if !isImageFile(file.Name) {
			continue
		}
		fileTargetPath := path.Join(dest, strconv.Itoa(i)+filepath.Ext(file.Name))
		createdFile, creationErr := os.Create(fileTargetPath)
		if creationErr != nil {
			fmt.Println(creationErr)
			continue
		}
		archivedFile, openErr := file.Open()

		if openErr != nil {
			continue
		}

		defer archivedFile.Close()

		_, copyErr := io.Copy(createdFile, archivedFile)

		if copyErr != nil {
			fmt.Println("Failed to copy the file ", createdFile.Name())
			continue
		}
	}
}

//export FreeStrings
func FreeStrings(str *C.char, count C.int) {
	C.free(unsafe.Pointer(str))
}

func main() {}
