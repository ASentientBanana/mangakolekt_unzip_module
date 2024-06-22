package tests

import (
	"github.com/asentientbanana/uz/unrar"
	"testing"
)

func TestUnrarSingle(t *testing.T) {
	mockInput := "/home/petar/bigboy/Manga/SMUT/Puppy Lovers v01 c01-14.cbr"
	testOut := "/home/petar/Projects/mangakolekt/unzip/.tmp"
	unrar.Unrar_Single_Book(
		mockInput,
		testOut,
	)
}
