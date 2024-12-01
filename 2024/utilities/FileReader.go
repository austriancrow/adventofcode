package utilities

import (
	"os"
)

func FileReader(filepath string) []byte {
	file_data, file_error := os.ReadFile(filepath)
	ErrorChecker(file_error)
	return file_data
}
