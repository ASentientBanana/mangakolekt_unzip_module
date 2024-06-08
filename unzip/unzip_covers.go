package unzip

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	p "path/filepath"
	"strings"

	"github.com/asentientbanana/uz/models"
	"github.com/asentientbanana/uz/util"
	"github.com/google/uuid"
)

func Unzip_covers_from_dir(jsonString string, output string) string {

	results := []string{}
	files := []models.Cover{}
	// json.Unmarshal()
	// files = strings.Split(filesString, "&&")

	//TODO: Steps
	// read json string and extract file data

	err := json.Unmarshal([]byte(jsonString), &files)
	if err != nil {
		errString := fmt.Sprintf("{ error:  %s}", err)
		return errString
	}

	for _, dirFile := range files {
		nameID := uuid.New().String()
		archive, err := zip.OpenReader(dirFile.DirectoryFile)

		if err != nil {

			fmt.Println("Failed to open zip file:", err)
			fmt.Println("Error", dirFile)
			continue
		}
		defer archive.Close()

		for _, f := range archive.File {
			//This is to ignore if its a dir
			if f.FileInfo().IsDir() {
				continue
			}

			if !util.IsImageFile(f.Name) {
				continue
			}

			_, dir := p.Split(f.Name)
			fmt.Println(dir)

			// dstPath := p.Join(output, dir)
			dstPath := p.Join(output, nameID+filepath.Ext(f.Name))

			cf, errC := os.Create(dstPath)
			if errC != nil {
				fmt.Println(errC)
				continue
			}
			archivedFile, err := f.Open()
			defer archivedFile.Close()

			if err != nil {
				continue
			}
			io.Copy(cf, archivedFile)
			if err != nil {
				continue
			}

			splitPath := strings.Split(dirFile.DirectoryFile, "/")
			archiveName := splitPath[len(splitPath)-1]

			//template
			//cbz-name/cover-path/cbz-path
			results = append(results, archiveName+";"+dstPath+";"+dirFile)
			cf.Close()
			break
		}
	}
	return strings.Join(results, "&?&")
}
