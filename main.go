package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>



*/
import "C"

import (
	"archive/zip"
	"fmt"
	"image"
	"io"
	"os"
	"path"
	"path/filepath"
	p "path/filepath"
	"strconv"
	"strings"
	"unsafe"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
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

func removeAllContents(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

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

			dstPath := p.Join(output, nameID+filepath.Ext(f.Name))
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

func markFile(file *os.File, dest string) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return
	}
	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return
	}
	isDouble := image.Width > image.Height

	if isDouble {
		_, f := path.Split(file.Name())
		newPath := filepath.Join(dest, "_"+f)
		fmt.Println(f)
		fmt.Println(newPath)
		err := os.Rename(file.Name(), newPath)
		if err != nil {
			fmt.Println("Err::")
			fmt.Println(err)
			return
		}

	}
}

//export Unzip_Single_book
func Unzip_Single_book(_filePath *C.char, _dest *C.char) C.int {

	// Convert C string to Go string
	zipPath := C.GoString(_filePath)
	dest := C.GoString(_dest)
	// zipPath := (_filePath)
	// dest := (_dest)

	removeAllContents(dest)

	// Open the zip file

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		fmt.Println("Failed to open zip file:", err)
		return 1
	}
	defer r.Close()

	// Iterate over each file in the zip archive
	for i, file := range r.File {

		if file.FileInfo().IsDir() {
			continue
		}
		// Check if the file is an image (you can modify this condition as needed)
		if !isImageFile(file.Name) {
			continue
		}

		archivedFile, openErr := file.Open()

		name := strconv.Itoa(i) + filepath.Ext(file.Name)

		fileTargetPath := path.Join(dest, name)
		createdFile, creationErr := os.Create(fileTargetPath)
		if creationErr != nil {
			fmt.Println(creationErr)
			continue
		}

		if openErr != nil {
			continue
		}

		defer archivedFile.Close()

		_, copyErr := io.Copy(createdFile, archivedFile)

		if copyErr != nil {
			fmt.Println("Failed to copy the file ", createdFile.Name())
			continue
		}
		markFile(createdFile, dest)
	}
	return 0
}

//export FreeStrings
func FreeStrings(str *C.char, count C.int) {
	C.free(unsafe.Pointer(str))
}

func main() {
	// _unzipSingle("/home/petar/bigboy/Manga/OnePiece/", "/home/petar/Documents/mangakolekt/current")
	// Unzip_Single_book()
}
