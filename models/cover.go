package models

type Cover struct {
	ArchiveName     string `json:"archiveName"`
	DestinationPath string `json:"destinationPath"`
	DirectoryFile   string `json:"directoryFile"`
}
type FFICoverResponse struct {
	Error  string  `json:"error"`
	Covers []Cover `json:"covers"`
}

type TargetFiles struct {
	Files []string `json:"files"`
}
