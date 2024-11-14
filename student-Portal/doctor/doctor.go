package doctor

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	basicdata "studentPortal/basicData"
	"studentPortal/student"
	"time"
)

type optionalStaffStudentByDoctor func(*student.Student)

// optional func ! the function version of the AddLateCourses method
// which add any number of late courses
func WithLateCourses(subs ...student.SubjectName) optionalStaffStudentByDoctor {
	return func(s *student.Student) {
		s.LateCourses = append(s.LateCourses, subs...)
	}
}

// optional func ! add some grades to the student specific course
func WithGrades(grades map[*student.SubjectName]student.Grade) optionalStaffStudentByDoctor {
	return func(s *student.Student) {
		for k, v := range grades {
			s.Grades[*k] = v
		}

	}
}

// This function adds a new student to the students.json file.
// It reads the existing students from the file, appends the new student to the list,
// and then writes the updated list back to the file.
func AddStudent(name string, dateOfBirth time.Time,
	id string, GPA float32,
	CurrentSemester int, HoursCompleted int,
	LateCourses []student.SubjectName,
	studentOptionalData ...basicdata.OptionalArguments) error {

	myBasicData := basicdata.NewBasicData(name,
		dateOfBirth, id, studentOptionalData...)

	myStudent := student.Student{
		BasicData:       *myBasicData,
		Gpa:             GPA,
		CurrentSemester: CurrentSemester,
		HoursCompleted:  HoursCompleted,
		LateCourses:     LateCourses}

	fileData, err2 := os.ReadFile("../json/students.json")
	if err2 != nil {
		return errors.New("cannot read ../json/students.json file")
	}

	// Unmarshal existing JSON data into a slice of Person objects
	var existingData []student.Student
	if len(fileData) > 0 { // Check if file is not empty
		err := json.Unmarshal(fileData, &existingData)
		if err != nil {
			return errors.New("cannot unmarshal file data in AddStudent function")
		}
	}
	// Append new data to the existing data
	existingData = append(existingData, myStudent)

	updatedData, err := json.MarshalIndent(existingData, "", "   ")
	if err != nil {
		return errors.New("cannot marshal updated student data in AddStudent function")
	}
	err3 := os.WriteFile("../json/students.json", updatedData, 0644)

	if err3 != nil {
		return errors.New("cannot write the updated student data to the JSON file")
	}

	return nil

}

func DocotrStartScreen(currentDoctor basicdata.Doctor) {
	fmt.Printf(`         Welcome %s , you are already stored in our memory
	 and your id is %s welcome back! .  
	 ---------------------------------------------------------
	 1) add student  
	 2) delete student  
	 3) add attendance
	 4) add grades
	 ---------------------------------------------------------`+"\n", currentDoctor.Name, currentDoctor.Id)
}


func AddDoctor(){

	var newDoc basicdata.Doctor = basicdata.Doctor{BasicData: basicdata.BasicData{Name: "girgis" , Id: "1" , Age: 20 , Gender: "male" }}
	var sliceDoc []basicdata.Doctor = []basicdata.Doctor{newDoc}
	jsonData, err := json.Marshal(sliceDoc)
	if err != nil {
		fmt.Println("Error marshaling doctor data:", err)
		return
	}
	fmt.Println("Doctor data:", string(jsonData))
	err = os.WriteFile("json/doctors.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing doctor data to file:", err)
		return
	}
}