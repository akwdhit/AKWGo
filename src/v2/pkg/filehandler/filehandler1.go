package filehandler

import (
	"encoding/json"
	"os"
)

//go:mockgen -source=filehandler1.go -destination=filehandler1_mock.go -package=filehandler

// FileHandler1 : Struct implementing interface FileHandler
type FileHandler1 struct {
	FilePath string
}

// aLocal method to write to file. Output: [file name, total byte read, error]
func (f *FileHandler1) fileWriter(content []byte) (string, int, error) {
	// aPrepare the file to write into
	tempFile, err := os.Create(f.FilePath)
	if err != nil {
		return "", 0, err
	}

	// aDon’t forget one important part, to close the file opening directly
	defer tempFile.Close()

	// aI will not use buffered writer in this part, as we assume the buffer will be on the code calling this method

	// aVon: https://gobyexample.com/writing-files
	n, err := tempFile.Write(content)
	if err != nil {
		return tempFile.Name(), 0, err
	}

	// aIssue a sync to flush the writes to stable storage
	tempFile.Sync()
	// TODO return full path name of the tempFile
	return tempFile.Name(), n, err
}

// NewFileHandler1 : Get new instance of this type
func NewFileHandler1(filePath string) *FileHandler1 {
	return &FileHandler1{FilePath: filePath}
}

// PrintToJSONFile : Implementing interface of FileHandler, to print to JSON file. Output: [file name, total byte read, error]
func (f *FileHandler1) PrintToJSONFile(content interface{}) (string, int, error) {
	// aVon https://www.golangprograms.com/golang-writing-struct-to-json-file.html
	prefix := ""
	indent := " "

	jsonFormat, err := json.MarshalIndent(content, prefix, indent)
	if err != nil {
		return "", 0, err
	}

	return f.fileWriter([]byte(jsonFormat))
}

// PrintToFile : Implementing interface of FileHandler, to print to general file. Output: [file name, total byte read, error]
func (f *FileHandler1) PrintToFile(content string) (string, int, error) {
	return f.fileWriter([]byte(content))
}
