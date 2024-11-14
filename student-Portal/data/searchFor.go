package data

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	basicdata "studentPortal/basicData"
	"studentPortal/student"
)

func WriteJSONToFile(filename string, data interface{}) error {
	file, err := os.Create(filename) // Creates or truncates the file
	if err != nil {
		return errors.New("cannot create file writeJSONToFile func")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

// Not competed need to do it
func SearchForTheId(SorD string, id string) (bool, error) {

	if strings.ToLower(SorD)[0] == 's' {
		var dataComeFromJsonFile []student.Student
		data, err := os.ReadFile("json/students.json")
		json.Unmarshal(data, &dataComeFromJsonFile)
		if err != nil {
			return false, errors.New("cannot read the file searchfortheid func")
		}
		for _, obj := range dataComeFromJsonFile {
			if obj.Id == id {
				return true, nil
			}
		}
	} else {

		var dataComeFromJsonFile []basicdata.Doctor
		data, err := os.ReadFile("json/doctors.json")
		json.Unmarshal(data, &dataComeFromJsonFile)
		if err != nil {
			return false, errors.New("cannot read the file searchfortheid func")
		}
		for _, obj := range dataComeFromJsonFile {
			if obj.Id == id {
				return true, nil
			}
		}
	}
	return false, nil
}
