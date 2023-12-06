package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/asentientbanana/uz/util"
)

func Unzip_Single_Book(zipPath string, dest string) []string {

	util.RemoveAllContents(dest)

	var orderedContent []string

	// Open the zip file
	r, err := zip.OpenReader(zipPath)

	if err != nil {
		fmt.Println("Failed to open zip file:", err)
		return []string{}
	}
	defer r.Close()

	// Iterate over each file in the zip archive
	for _, file := range r.File {

		if file.FileInfo().IsDir() {
			os.Mkdir(path.Join(dest, file.Name), os.ModePerm)
			continue
		}

		// Check if the file is an image (you can modify this condition as needed)
		if !util.IsImageFile(file.Name) {
			continue
		}

		archivedFile, openErr := file.Open()

		defer archivedFile.Close()

		_, f := path.Split(file.Name)

		// newPath = strings.Replace(newPath, f, "_"+f, 1)
		nameSlices := strings.Split(file.FileInfo().Name(), "-")
		name := strings.Replace(file.Name, f, nameSlices[0], 1)
		fileTargetPath := path.Join(dest, name)

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

		filePath, err := util.MarkFile(createdFile)
		if err != nil {
			return []string{}
		}

		orderedContent = append(orderedContent, filePath)
	}
	return orderedContent
}
