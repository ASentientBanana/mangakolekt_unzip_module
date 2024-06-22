package tests

import (
	"archive/zip"
	"fmt"
	"testing"
)

func TestAnalyze(t *testing.T) {
	zipPath := "/home/petar/bigboy/Manga/OnePiece/Vol. 95.cbz"
	// zipPath := "/home/petar/bigboy/Manga/Vagabond/# 181.cbz"
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range r.File {
		fmt.Println(i)
		if v.FileInfo().IsDir() {
			fmt.Println(v.FileInfo().Name())
		} else {
			fmt.Println("  |")
			fmt.Println("  -->", v.FileInfo().Name())
		}
	}

}
