package sandbox

import (
	"archive/zip"
	"fmt"
	"image"
	"io"
	"os"
	"path"
	"path/filepath"
	p "path/filepath"
	"strconv"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ReadDir(path string) ([]string, error) {
	dir, err := os.ReadDir(path)
	files := []string{}
	if err != nil {
		return nil, err
	}
	for _, e := range dir {
		files = append(files, path+"/"+e.Name())
	}
	return files, nil
}

func removeAllContents(dirPath string) error {
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
		entryPath := p.Join(dirPath, entry.Name())
		err = os.RemoveAll(entryPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// Function to check if a file has an image extension
func isImageFile(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

func markFile(file *os.File, dest string) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return
	}
	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return
	}
	isDouble := image.Width > image.Height

	if isDouble {
		_, f := path.Split(file.Name())
		newPath := filepath.Join(dest, "_"+f)
		fmt.Println(f)
		fmt.Println(newPath)
		err := os.Rename(file.Name(), newPath)
		if err != nil {
			fmt.Println("Err::")
			fmt.Println(err)
			return
		}

	}
}

//export Unzip_Single_book
func unzipSingle(zipPath, dest string) {

	removeAllContents(dest)

	// Open the zip file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		fmt.Println("Failed to open zip file:", err)

	}
	defer r.Close()

	// Iterate over each file in the zip archive
	for i, file := range r.File {

		if file.FileInfo().IsDir() {
			continue
		}
		// Check if the file is an image (you can modify this condition as needed)
		if !isImageFile(file.Name) {
			continue
		}

		archivedFile, openErr := file.Open()

		name := strconv.Itoa(i) + filepath.Ext(file.Name)

		fileTargetPath := path.Join(dest, name)
		createdFile, creationErr := os.Create(fileTargetPath)
		if creationErr != nil {
			fmt.Println(creationErr)
			continue
		}

		if openErr != nil {
			continue
		}

		defer archivedFile.Close()

		_, copyErr := io.Copy(createdFile, archivedFile)

		if copyErr != nil {
			fmt.Println("Failed to copy the file ", createdFile.Name())
			continue
		}
		markFile(createdFile, dest)
	}
}

func main() {
	unzipSingle("/home/petar/bigboy/Manga/OnePiece/", "/home/petar/Documents/mangakolekt/current")
	// Unzip_Single_book()
}
