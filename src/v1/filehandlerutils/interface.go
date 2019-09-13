package filehandlerutils

// FileHandler : Interface of file handler
type FileHandler interface {
	PrintToJSONFile() (string, int, error)
	PrintToFile() (string, int, error)
	ReadFromFile() (string, int, error)
}
