package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileHandler struct {
	Type_of_file string
}

func (f_instance FileHandler) ReadFile(path string, dataPointer interface{}) error {
	// We can add more file types here, json for now
	if f_instance.Type_of_file == "json" {
		// Check if file exists, raising error if not
		file, err := os.Open(path)
		if err != nil {
			return errors.New("error reading file")
		} else {
			defer file.Close()
		}

		// Read the file
		decoder := json.NewDecoder(file)
		err = decoder.Decode(dataPointer)
		if err != nil {
			return errors.New("error decoding json file")
		}
		return nil
	} else {
		return errors.New("not supported file type")
	}
}

func (f_instance FileHandler) WriteFile(path string, data any) {
	// We can add more file types here, json for now
	if f_instance.Type_of_file == "json" {
		fmt.Println("Writing JSON file", path)
	} else {
		fmt.Println("Not supported File type:", f_instance.Type_of_file)
	}
}
