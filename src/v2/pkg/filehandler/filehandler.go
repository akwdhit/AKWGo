package filehandler

//go:mockgen -source=filehandler.go -destination=filehandler_mock.go -package=filehandler

// FileHandler : Interface of file handler
type FileHandler interface {
	PrintToFile(content string) (string, int, error)
	PrintToJSONFile(content interface{}) (string, int, error)
}
