package filehandler

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// TestMockPrintToFile : Testing for Mock of Print To JSON File
func TestMockPrintToFile(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	fileHandler := NewMockFileHandler(controller)
	fileHandler.EXPECT().PrintToFile("PrintToFileContent").Return("", 0, nil)

	fileHandler.PrintToFile("PrintToFileContent")
}

// TestMockPrintToJSONFile : Testing for Mock of Print To JSON File
func TestMockPrintToJSONFile(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	fileHandler := NewMockFileHandler(controller)
	fileHandler.EXPECT().PrintToJSONFile("PrintToJSONFileContent").Return("", 0, nil)

	fileHandler.PrintToJSONFile("PrintToJSONFileContent")
}
