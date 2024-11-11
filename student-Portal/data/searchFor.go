package data

import (
	"os"
)

// Not competed need to do it
func SearchForTheId() bool {
	fileLocation := "../json/students.json"
	_,err :=os.ReadFile(fileLocation)
	if err != nil {
		return false
	}
	return false
}