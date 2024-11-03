package main

import (
	"bufio"
	"fmt"
	"os"
)

// -----------------------------------------------
// Write to a new text file or overwrite an existing one
func writeToFile(filename, content string) error {
	file, err := os.Create(filename) // Creates or truncates the file
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------------------
// Append text to an existing file (or create if not exists)
func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------------------
// Read entire file content as a single string
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// -----------------------------------------------
// Read file line by line and print to console
func readFileLineByLine(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// -----------------------------------------------
// Write lines to file (replaces existing content)
func writeLinesToFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// -----------------------------------------------
// Append lines to an existing file (or create if not exists)
func appendLinesToFile(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// -----------------------------------------------
// Delete a text file
func deleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("error deleting file: %v", err)
	}
	fmt.Println("File deleted successfully!")
	return nil
}



