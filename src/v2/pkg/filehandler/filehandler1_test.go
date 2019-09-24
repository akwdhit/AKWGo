package filehandler

import (
	"testing"
)

// TestMockPrintToFile : Testing for Mock of Print To JSON File
func TestPrintToFile(t *testing.T) {
	fileHandler := NewFileHandler1("PrintToFile.txt")
	fileHandler.PrintToFile("PrintToFileContent")
}

// TestMockPrintToJSONFile : Testing for Mock of Print To JSON File
func TestPrintToJSONFile(t *testing.T) {
	fileHandler := NewFileHandler1("PrintToJSONFile.txt")
	fileHandler.PrintToJSONFile("PrintToJSONFileContent")
}
