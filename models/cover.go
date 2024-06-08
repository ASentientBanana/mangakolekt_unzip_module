package models

type Cover struct {
	ArchiveName     string `json:"archiveName"`
	DestinationPath string `json:"destinationPath"`
	DirectoryFile   string `json:"directoryFile"`
}

type TargetFiles struct {
	Files []string `json:"files"`
}
