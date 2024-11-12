package doctor

import (
	"encoding/json"
	"errors"
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

// Not complete yet finish it you have a location on json called students to append the newStudent to it
func AddStudent(name string, dateOfBirht time.Time,
	id string, GPA float32,
	CurrentSemester int, HoursCompleted int,
	LateCourses []student.SubjectName,
	stuentOptionalData ...basicdata.OptionalArguments) error {

	myBasicData := basicdata.NewBasicData(name,
		dateOfBirht, id, stuentOptionalData...)

	myStudent := student.Student{
		BasicData:       *myBasicData,
		Gpa:             GPA,
		CurrentSemester: CurrentSemester,
		HoursCompleted:  HoursCompleted,
		LateCourses:     LateCourses}

	fileData, err2 := os.ReadFile("../json/students.json")
	if err2 != nil {
		return errors.New("connot read ../json/student.json file")
	}

	// Unmarshal existing JSON data into a slice of Person objects
	var existingData []student.Student
	if len(fileData) > 0 { // Check if file is not empty
		err := json.Unmarshal(fileData, &existingData)
		if err != nil {
			return errors.New("cannot unmarshal file data from : addstudent func")
		}
	}
	// Append new data to the existing data
	existingData = append(existingData, myStudent)

	updatedData, err := json.MarshalIndent(existingData, "", "   ")
	if err != nil {
		return errors.New("cannot marshal addStudent func")
	}
	err3 := os.WriteFile("../json/students.json", updatedData, 0644)

	if err3 != nil {
		return errors.New("cannot write the new student to a json file")
	}

	return nil

}
