package unrar

import (
	"github.com/asentientbanana/uz/util"
	rar "github.com/gen2brain/go-unarr"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Unrar_Single_Book(zipPath string, dest string) {

	util.RemoveAllContents(dest)

	archive, err := rar.NewArchive(zipPath)
	if err != nil {
		return
	}
	_, extractErr := archive.Extract(dest)

	if extractErr != nil {
		return
	}
}
