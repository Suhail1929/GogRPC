package data

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Data struct {
	Json string
}

// ReadFiles reads the content of a JSON file and returns it as a Data object
func ReadFiles(file *os.File) (Data, error) {
	// Read the JSON file
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return Data{}, fmt.Errorf("failed to read JSON file: %w", err)
	}

	return Data{Json: string(jsonData)}, nil
}

// gRPC server address

func (s Data) SendToServer() (string, error) {

	return "json", nil
}
