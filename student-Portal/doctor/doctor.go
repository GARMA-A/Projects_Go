package doctor

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	basicdata "studentPortal/basicData"
	"studentPortal/commands"
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
func AddStudent(newstudent student.Student) error {

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
	existingData = append(existingData, newstudent)

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
// so here is the problem when you hit 'r' after press 1 to add student 
// it return succesful but when you enter 1 again do not go again to 
// option one function please fix this problem 
func DocotrStartScreen(currentDoctor basicdata.Doctor) string {
	commands.ClearConsole()
	fmt.Printf(`         Welcome %s , you are already stored in our memory
	 and your id is %s welcome back! .  
	 ---------------------------------------------------------
	 1) add student  
	 2) delete student  
	 3) add attendance
	 4) add grades
	 ---------------------------------------------------------`+"\n", currentDoctor.Name, currentDoctor.Id)
	var ch string
	fmt.Print("Enter choice : ")
	fmt.Scanf("%s", &ch)
	return ch

}

func OptionScreen(option string, currentDoctor basicdata.Doctor) {
	switch option {
	case "1":
		OptionOneOnDoctor(currentDoctor)
	default:
		print("This is bad input not 1,2,3 or 4")
		fmt.Scanln()
		DocotrStartScreen(currentDoctor)
	}
}

func OptionOneOnDoctor(currentDoctor basicdata.Doctor) {
	var innerSwitchOption string
	commands.ClearConsole()
	fmt.Print("Ok to add a new student you will need to enter the \n (name ,date of birth, id)  and there is some optional staff like \n (phone number , gender , address) \n are you want to complete ?")
	fmt.Print("Press 'c' to complete or press 'r' to return :")
	fmt.Scanf("%s", &innerSwitchOption)
	if innerSwitchOption == "r" {
		DocotrStartScreen(currentDoctor)
	} else if innerSwitchOption == "c" {
		commands.ClearConsole()
		var newStudent student.Student
		fmt.Printf("So please enter the name of the student : ")
		fmt.Scanf("%s:", newStudent.Name)
		var birthDateStr string
	takeTheBirthDate:
	       fmt.Printf("So please enter student birth date (YYYY-MM-DD) : ")
		fmt.Scanf("%s:", &birthDateStr)
		realBirthDate, err := time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			print("there is error on this birth date format please enter the date like this (YYYY-MM-DD) ", err)
			fmt.Scan()
			commands.ClearConsole()
			goto takeTheBirthDate
		}
		newStudent.DateOfBirth = realBirthDate
		fmt.Printf("Now enter the id : ")
		fmt.Scanf("%s", &newStudent.Id)
		fmt.Printf("Very good now the student data is \n Name : %s  \n Birth Date %s  \n  Id : %s ", newStudent.Name, newStudent.DateOfBirth.Format("2006-01-02"), newStudent.Id)
		fmt.Printf("\nTo continue press enter ... ")
		fmt.Scan()
		err2 := AddStudent(newStudent)
		if err2 != nil {
			print(err2)
		} else {

			fmt.Println("student added succefully!")
		}
		fmt.Scan()
		commands.ClearConsole()

	} else {
		print("Invalid Option...")
		fmt.Scan()
		OptionOneOnDoctor(currentDoctor)
	}
}
func AddDoctor() {

	var newDoc basicdata.Doctor = basicdata.Doctor{BasicData: basicdata.BasicData{Name: "girgis", Id: "1", Age: 20, Gender: "male"}}
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
