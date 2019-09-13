package filehandlerutils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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

// ReadFromJSON : Implementing interface of FileHandler, to read from JSON. Output: [string format of the file content, length of the file byte, err if any]
func (f *FileHandler1) ReadFromJSON(s *interface{}) (string, int, error) {
	// aVon https://www.golangprograms.com/golang-read-json-file-into-struct.html
	file, err := ioutil.ReadFile("files/result/republish_order.json")
	if err != nil {
		return "", 0, err
	}

	err = json.Unmarshal([]byte(file), &s)

	return string(file), len(file), err
}
