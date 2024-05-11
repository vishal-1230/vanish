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

const (
	green  = "\033[32m"
	reset  = "\033[0m"
	yellow = "\033[33m"
	red    = "\033[31m"
)

func (f_instance FileHandler) WriteFile(path string, data interface{}) error {
	// We can add more file types here, json for now
	if f_instance.Type_of_file == "json" {
		fmt.Println("[", yellow, "•", reset, "]", "Writing JSON file", path)
		file, err := os.Create(path)
		if err != nil {
			return errors.New("error creating file")
		} else {
			defer file.Close()
		}

		// Write the file
		encoder := json.NewEncoder(file)
		err = encoder.Encode(data)
		if err != nil {
			return errors.New("error encoding json file")
		}
		return nil
	} else {
		return errors.New("not supported file type: " + f_instance.Type_of_file)
	}
}
