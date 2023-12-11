package util

import (
	"fmt"
	"image"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func IsImageFile(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

func RemoveAllContents(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := path.Join(dirPath, entry.Name())
		err = os.RemoveAll(entryPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func MarkFile(file *os.File) (string, error) {

	_, err := file.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	image, _, err := image.DecodeConfig(file)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	isDouble := image.Width > image.Height

	if !isDouble {
		return file.Name(), nil
	}
	// _, f := path.Split(file.Name())
	newPath := file.Name()
	ext := path.Ext(newPath)
	newPath = strings.Replace(newPath, ext, "__wide__"+ext, 1)
	os.Rename(file.Name(), newPath)
	if err != nil {
		fmt.Println("Err::")
		fmt.Println(err)
		return "", err
	}

	return newPath, nil

}

func ExtractNumericPart(s string) int {
	var numericPart string
	for _, char := range s {
		if char >= '0' && char <= '9' {
			numericPart += string(char)
		}
	}
	num, _ := strconv.Atoi(numericPart)
	return num
}
