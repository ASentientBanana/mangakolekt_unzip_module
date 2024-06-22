package tests

import (
	"fmt"
	"github.com/asentientbanana/uz/unrar"
	"testing"
)

func TestUnrar(t *testing.T) {
	mockInput := `["/home/petar/bigboy/Manga/SMUT/Puppy Lovers v01 c01-14.cbr"]`
	testOut := "/home/petar/Projects/mangakolekt/unzip/.tmp"

	result := unrar.Unrar_covers_from_dir(mockInput, testOut)
	fmt.Println("ReS:: ", result)
}
