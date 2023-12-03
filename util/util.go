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

func MarkFile(file *os.File, dest string) {

	_, err := file.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	image, _, err := image.DecodeConfig(file)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	isDouble := image.Width > image.Height

	fmt.Println("is Double ", isDouble)
	if !isDouble {
		return
	}
	_, f := path.Split(file.Name())
	newPath := file.Name()
	newPath = strings.Replace(newPath, f, "_"+f, 1)
	fmt.Println("Looking at: ", newPath, " from:  ", f, "to: ", "_"+f)
	os.Rename(file.Name(), newPath)
	if err != nil {
		fmt.Println("Err::")
		fmt.Println(err)
		return
	}
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
