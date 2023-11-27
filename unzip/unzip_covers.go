package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	p "path/filepath"
	"strings"

	"github.com/asentientbanana/uz/util"
	"github.com/google/uuid"
)

func Unzip_covers_from_dir(filesString string, output string) string {
	results := []string{}
	files := []string{}
	files = strings.Split(filesString, "&&")

	for _, dirFile := range files {
		nameID := uuid.New().String()
		archive, err := zip.OpenReader(dirFile)

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

			splitPath := strings.Split(dirFile, "/")
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
