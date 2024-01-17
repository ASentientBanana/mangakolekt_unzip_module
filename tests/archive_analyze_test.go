package tests

import (
	"fmt"
	"testing"

	"github.com/asentientbanana/uz/unrar"
)

func TestAnalyze(t *testing.T) {
	filePath := "/home/petar/bigboy/Manga/SMUT/Puppy Lovers v01 c01-14.cbr"
	res := unrar.Unrar_Single_Book(filePath, "/home/petar/Documents/mangakolekt/current")
	fmt.Println("Res: ", res)
}
