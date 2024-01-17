package tests

import (
	"archive/zip"
	"fmt"
	"testing"

	"github.com/asentientbanana/uz/unzip"
)

var testCases = []string{
	"/home/petar/bigboy/Manga/OnePiece/Vol. 95.cbz",
}

func TestFindCover(t *testing.T) {
	target := testCases[0]
	archive, err := zip.OpenReader(target)

	if err != nil {
		fmt.Println("Failed to open zip file:", err)
		fmt.Println("Error")
	}
	defer archive.Close()

	e, i := unzip.FindCoverFromZip(archive.File)
	fmt.Println("Found item: ", e)
	fmt.Println("Found index: ", i)
}
