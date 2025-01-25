package utils

import (
	"io/ioutil"
)

// ReadFile reads the contents of a file and returns it as a byte slice
func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
