package unrar

import (
	"context"
	"fmt"
	"io"
	"os"
	p "path/filepath"
	"strings"

	"github.com/mholt/archiver/v4"
)

func Unrar_Single_Book(filesString string, output string) string {
	fmt.Println("Starting")
	files := strings.Split(filesString, "&&")

	for _, dirFile := range files {
		format := archiver.Rar{}
		ctx := context.Background()
		handler := func(ctx context.Context, f archiver.File) error {
			targetPath := p.Join(output, f.NameInArchive)

			new_file, err := os.Create(targetPath)
			defer new_file.Close()

			if err != nil {
				fmt.Println("Problem creating file from  archive")
				fmt.Println(err)

				return err
			}
			f_reader, err := f.Open()
			defer f_reader.Close()
			if err != nil {
				fmt.Println("Problem loading file from  archive")
				return err
			}
			_, cErr := io.Copy(new_file, f_reader)

			if cErr != nil {
				fmt.Println("Problem copy")
				return err
			}

			return nil
		}
		f, err := os.Open(dirFile)
		defer f.Close()
		// f, err := os.OpenFile(dirFile, os.O_RDONLY, 0755)
		if err != nil {
			continue
		}

		format.Extract(ctx, f, nil, handler)

		if err != nil {
			fmt.Println(err)
		}
	}

	// err := format.Extract(ctx)
	// return strings.Join(results, "&?&")
	return ""
}
