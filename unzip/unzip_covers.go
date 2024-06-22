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

	results := []models.Cover{}
	files := []string{}
	err := json.Unmarshal([]byte(jsonString), &files)
	if err != nil {
		//errString := fmt.Sprintf("{ error:  %s}", err)
		return "[]"
	}

	for _, dirFile := range files {
		nameID := uuid.New().String()
		archive, err := zip.OpenReader(dirFile)

		if err != nil {
			continue
		}

		for _, f := range archive.File {
			//This is to ignore if it's a dir
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

			if err != nil {
				continue
			}
			_, copyErr := io.Copy(cf, archivedFile)
			if copyErr != nil {
				continue
			}

			splitPath := strings.Split(dirFile, "/")
			archiveName := splitPath[len(splitPath)-1]

			//template
			//cbz-name/cover-path/cbz-path
			results = append(results, models.Cover{
				ArchiveName:     archiveName,
				DestinationPath: dstPath,
				DirectoryFile:   dirFile,
			})

			archivedFile.Close()
			cf.Close()
			break
		}

		archive.Close()
	}
	jsonResponse, err := json.Marshal(results)
	if err != nil {
		return `[]`
	}
	return string(jsonResponse)
}
