package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	p "path/filepath"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/maruel/natural"
)

func FindCoverFromZip(files []*zip.File) (string, int) {
	fNames := []string{}
	var position int = 0
	for _, f := range files {
		if !f.FileInfo().IsDir() {
			fNames = append(fNames, f.Name)
			fmt.Println(f.Name)
		}
	}
	sort.Sort(natural.StringSlice(fNames))
	for index, f := range files {
		if f.Name == fNames[0] {
			position = index
		}
	}
	return fNames[0], position
}

func Split_files(files []string) [][]string {
	split_files := [][]string{}
	file_number := len(files)
	items_per_section := 4
	number_of_sections := math.Ceil(float64(file_number % items_per_section))
	fmt.Println("Number of sections: ", number_of_sections)
	if file_number <= items_per_section {
		split_files = append(split_files, files)
		return split_files
	}
	for i := 0; i < int(number_of_sections); i++ {
		if i == int(number_of_sections)-1 {
			split_files = append(split_files, files[i*4:file_number])
			break
		}
		curr := i * 4
		split_files = append(split_files, files[curr:curr+4])

	}
	return split_files
}

func Unzip_covers_from_dir(filesString string, output string) string {
	results := []string{}
	files := strings.Split(filesString, "&&")

	for _, dirFile := range files {
		nameID := uuid.New().String()
		archive, err := zip.OpenReader(dirFile)

		if err != nil {
			fmt.Println("Failed to open zip file:", err)
			fmt.Println("Error", dirFile)
			continue
		}
		defer archive.Close()
		res, pos := FindCoverFromZip(archive.File)

		dstPath := p.Join(output, nameID+filepath.Ext(res))

		cf, errC := os.Create(dstPath)
		if errC != nil {
			fmt.Println(errC)
			continue
		}
		defer cf.Close()
		archivedFile, err := archive.File[pos].Open()

		defer archivedFile.Close()
		_, cErr := io.Copy(cf, archivedFile)
		if cErr != nil {
			continue
		}

		splitPath := strings.Split(dirFile, "/")
		archiveName := splitPath[len(splitPath)-1]

		results = append(results, archiveName+";"+dstPath+";"+dirFile)
	}
	return strings.Join(results, "&?&")
}
