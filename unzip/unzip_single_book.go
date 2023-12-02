package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/asentientbanana/uz/util"
)

func Unzip_Single_Book(zipPath string, dest string) int {

	util.RemoveAllContents(dest)

	// Open the zip file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		fmt.Println("Failed to open zip file:", err)
		return 1
	}
	defer r.Close()

	var test []string

	for _, file := range r.File {
		test = append(test, file.FileInfo().Name())
		fmt.Println(file.FileInfo().Name() + "\n")
	}

	// Iterate over each file in the zip archive
	for i, file := range r.File {

		if file.FileInfo().IsDir() {
			continue
		}

		// Check if the file is an image (you can modify this condition as needed)
		if !util.IsImageFile(file.Name) {
			continue
		}

		archivedFile, openErr := file.Open()
		defer archivedFile.Close()

		// name := strconv.Itoa(i) + filepath.Ext(file.Name)
		fileTargetPath := path.Join(dest, strconv.Itoa(i)+"_"+path.Ext(file.FileInfo().Name()))
		createdFile, creationErr := os.Create(fileTargetPath)
		if creationErr != nil {
			fmt.Println(creationErr)
			continue
		}

		if openErr != nil {
			continue
		}

		_, copyErr := io.Copy(createdFile, archivedFile)

		if copyErr != nil {
			fmt.Println("Failed to copy the file ", createdFile.Name())
			continue
		}
		util.MarkFile(createdFile, dest)
	}
	return 0
}
