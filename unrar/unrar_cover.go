package unrar

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mholt/archiver/v4"
)

func Unrar_covers_from_dir(filesString string, output string) string {

	files := strings.Split(filesString, "&&")

	for _, dirFile := range files {
		format := archiver.Rar{}
		ctx := context.Background()
		handler := func(ctx context.Context, f archiver.File) error {

			fmt.Println("Name: ")
			fmt.Println(f.Name())
			fmt.Println("Name in archive: ")
			fmt.Println(f.NameInArchive)

			return nil
		}
		f, err := os.Open(dirFile)
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
