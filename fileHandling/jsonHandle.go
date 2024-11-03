package main

import (
	"encoding/json"
	"fmt"
	"os"
)
// 0 - No permissions (---):
// No permission to read, write, or execute the file.
// 1 - Execute permission (--x):
// The user/group/other can execute the file (if it's a program or script) but cannot read or write to it.
// 2 - Write permission (-w-):
// The user/group/other can write (modify) the file but cannot read or execute it.
// 3 - Write and execute permissions (-wx):
// The user/group/other can write to and execute the file, but cannot read it.
// 4 - Read permission (r--):
// The user/group/other can read the file but cannot write to or execute it.
// 5 - Read and execute permissions (r-x):
// The user/group/other can read and execute the file, but cannot write to it.
// 6 - Read and write permissions (rw-):
// The user/group/other can read and write to the file but cannot execute it.
// 7 - Read, write, and execute permissions (rwx)
// -----------------------------------------------

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// -----------------------------------------------
// Write to a new JSON file or overwrite an existing one
func writeJSONToFile(filename string, data interface{}) error {
	file, err := os.Create(filename) // Creates or truncates the file
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------------------
// Append JSON data to an existing file
func appendJSONToFile(filename string, newData Person) error {
	// Read existing JSON data from the file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal existing JSON data into a slice of Person objects
	var existingData []Person
	if len(fileData) > 0 { // Check if file is not empty
		err = json.Unmarshal(fileData, &existingData)
		if err != nil {
			return err
		}
	}

	// Append new data to the existing data
	existingData = append(existingData, newData)

	// Write the updated data back to the file
	return writeJSONToFile(filename, existingData)
}

// -----------------------------------------------
// Write JSON data quickly to a new file (overwrites if exists)
func fastWritingJSON() {
	data := []byte(`{"name": "Bob", "age": 25}`)
	err := os.WriteFile("output.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

// -----------------------------------------------
// Write JSON data by appending at the end of file without overwriting
func writeToJson1(filename, name string, age int) error {
// os.O_RDWR: This constant specifies that the file 
// should be opened for both reading and writing.
//  It allows you to read from and write to the file.
//------------------------------------------------------------
// os.O_CREATE: This constant specifies that the file
// should be created if it does not already exist.
// If the file exists, it will be opened without truncating it.
//------------------------------------------------------------
// os.O_APPEND: This constant specifies that all writes to the file
// should append to the end of the file rather than overwriting
// existing content. It is usually used in conjunction with os.O_CREATE
// to ensure that writes are appended to the file.

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]interface{}{
		"name": name,
		"age":  age,
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------------------
// Append new JSON fields to an existing JSON file (single object)
func appendToJsonFile() {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the existing JSON data into a map
	existingData := map[string]interface{}{}
	err = json.Unmarshal(fileData, &existingData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Append new data to the existing data
	newData := map[string]interface{}{
		"newKey": "newValue",
	}

	for key, value := range newData {
		existingData[key] = value
	}

	// Write the combined data back to the file
	updatedData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("data.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Data appended successfully!")
}

// -----------------------------------------------
// Delete a JSON file
func deleteJSONFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("error deleting file: %v", err)
	}
	fmt.Println("File deleted successfully!")
	return nil
}