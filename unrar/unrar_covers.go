package unrar

import (
	"encoding/json"
	"fmt"
	"github.com/asentientbanana/uz/models"
	rar "github.com/gen2brain/go-unarr"
	"github.com/google/uuid"
	"os"
	p "path/filepath"
	"sort"
	"strings"
)

func Urar_covers_from_dir(jsonString string, output string) string {

	var results []models.Cover
	var files []string

	err := json.Unmarshal([]byte(jsonString), &files)
	if err != nil {
		//errString := fmt.Sprintf("{ error:  %s}", err)
		return "[]"
	}

	for _file := range files {

		nameID := uuid.New().String()
		archive, err := rar.NewArchive(files[_file])
		if err != nil {
			fmt.Println(err)
			return "[]"
		}
		defer archive.Close()
		list, _ := archive.List()
		//fmt.Println(list)
		sort.Strings(list)
		entryErr := archive.EntryFor(list[0])
		//fmt.Println("Got:: ", archive.Name())
		if entryErr != nil {
			fmt.Println(entryErr)
			continue
		}

		splitPath := strings.Split(archive.Name(), "/")
		archiveName := splitPath[len(splitPath)-1]
		fmt.Println("GOT:: ", archiveName)
		dstPath := p.Join(output, nameID+p.Ext(archiveName))
		data, readErr := archive.ReadAll()

		if readErr != nil {
			continue
		}

		cf, createErr := os.Create(dstPath)
		defer cf.Close()
		if createErr != nil {
			continue
		}

		_, writeErr := cf.Write(data)
		if writeErr != nil {
			continue
		}

		results = append(results, models.Cover{
			ArchiveName:     archiveName,
			DestinationPath: dstPath,
			DirectoryFile:   files[_file],
		})
	}

	jsonResponse, err := json.Marshal(results)
	if err != nil {
		return `[]`
	}

	return string(jsonResponse)
}
