package doctor

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	basicdata "studentPortal/basicData"
	"studentPortal/commands"
	"studentPortal/data"
	"studentPortal/student"
	"time"
)

var GlobalCurrentDoctor basicdata.Doctor

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
func AddStudent(newstudent student.Student, fileRelativePath string) error {

	newstudent.Age = time.Now().Year() - newstudent.DateOfBirth.Year()

	fileData, err2 := os.ReadFile(fileRelativePath)
	if err2 != nil {
		fmt.Println(err2)
		return errors.New("cannot read " + fileRelativePath)
	}

	// Unmarshal existing JSON data into a slice of Person objects
	var existingData []student.Student
	if len(fileData) > 0 { // Check if file is not empty
		err := json.Unmarshal(fileData, &existingData)
		if err != nil {
			fmt.Println(err)
			return errors.New("cannot unmarshal file data in AddStudent function")
		}
	}
	// Append new data to the existing data
	existingData = append(existingData, newstudent)

	updatedData, err := json.MarshalIndent(existingData, "", "   ")
	if err != nil {
		return errors.New("cannot marshal updated student data in AddStudent function")
	}
	err3 := os.WriteFile(fileRelativePath, updatedData, 0644)

	if err3 != nil {
		fmt.Println(err3)
		return errors.New("cannot write the updated student data to the JSON file")
	}

	return nil

}

func DeleteStudent(index int) {

	fileByteSlice, err := os.ReadFile("json/students.json")
	if err != nil {
		commands.Pause(err)
	}
	var jsonFileData []student.Student
	err = json.Unmarshal(fileByteSlice, &jsonFileData)
	if err != nil {
		commands.Pause(err)
	}
	if index < 0 || index >= len(jsonFileData) {
		commands.Pause(fmt.Errorf("index out of range"))
		return
	}
	jsonFileData = append(jsonFileData[:index], jsonFileData[index+1:]...)
	fileByteSlice, err = json.MarshalIndent(jsonFileData, "", "   ")
	if err != nil {
		commands.Pause(err)
	}
	err = os.WriteFile("json/students.json", fileByteSlice, 0644)
	if err != nil {
		commands.Pause(err)
	}

}

// so here is the problem when you hit 'r' after press 1 to add student
// it return succesful but when you enter 1 again do not go again to
// option one function please fix this problem
func DocotrStartScreen() {
	commands.ClearConsole()
	fmt.Printf(`Welcome %s , you are already stored in our memory
	 and your id is %s welcome back! .  
	 ---------------------------------------------------------
	 1) add student  
	 2) delete student  
	 3) add attendance
	 4) add grades
	 5) Exit
	 ---------------------------------------------------------`+"\n", GlobalCurrentDoctor.Name, GlobalCurrentDoctor.Id)
	var ch string
	fmt.Print("Enter choice : ")
	fmt.Scanf("%s", &ch)
	OptionScreenForDoctor(ch)
}

func OptionScreenForDoctor(option string) {
	switch option {
	case "1":
		OptionOneOnDoctor()
	case "2":
		OptionTwoOnDoctor()
	case "5":
		os.Exit(0)
	default:
		commands.Pause("This is bad input not 1,2,3 or 4")
		DocotrStartScreen()
	}
}

func OptionOneOnDoctor() {
	commands.ClearConsole()

	var innerSwitchOption string

	fmt.Print("Press 'c' to complete or press 'r' to return :")

	fmt.Scanf("%s", &innerSwitchOption)

	switch innerSwitchOption {
	case "r":
		DocotrStartScreen()
	case "c":
		commands.ClearConsole()
		var newStudent student.Student
	takeTheName:
		fmt.Printf("So please enter the name of the student : ")
		fmt.Scanf("%s:", &newStudent.Name)
		if newStudent.Name == "" {
			commands.Pause("Not valid name !")
			commands.ClearConsole()
			goto takeTheName
		}
		var birthDateStr string
	takeTheBirthDate:
		fmt.Printf("So please enter student birth date (YYYY-MM-DD) : ")
		fmt.Scanf("%s:", &birthDateStr)
		realBirthDate, err := time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			commands.Pause("Not Valid birthDate the format is  (YYYY-MM-DD) \n like this 2004-01-12 or 2002-12-02 ", err)
			commands.ClearConsole()
			goto takeTheBirthDate
		}
		newStudent.DateOfBirth = realBirthDate
	enterTheId:
		fmt.Printf("Now enter the id : ")
		fmt.Scanf("%s", &newStudent.Id)
		found, _, err := data.SearchForTheId("s", newStudent.Id)
		if err != nil {
			commands.Pause(err)
			print("\nThe optionOneOnDoctor func")
		}
		if found {
			commands.Pause("This id is already used before please Enter another one ")
			commands.ClearConsole()
			goto enterTheId
		}
		fmt.Printf("\nVery good now the student data is \n Name : %s  \n Birth Date %s  \n  Id : %s ", newStudent.Name, newStudent.DateOfBirth.Format("2006-01-02"), newStudent.Id)
		commands.Pause("\nTo continue press enter ... ")

		err2 := AddStudent(newStudent, "json/students.json")
		if err2 != nil {
			commands.Pause(err2)
		} else {
			commands.Pause("student added succefully!")
		}
		commands.ClearConsole()
		DocotrStartScreen()

	default:
		commands.Pause("Invalid Option...")
		OptionOneOnDoctor()
	}
}

func OptionTwoOnDoctor() {
startOfTheFunc:
	var innerOption, id string
	commands.ClearConsole()
	fmt.Printf("To delete student please enter et's ID (to cancel enter r) : ")
	fmt.Scan(&id)
	if id[0] == 'r' {
		commands.ClearConsole()
		DocotrStartScreen()
	}
	found, index, err := data.SearchForTheId("s", id)
	if err != nil {
		commands.Pause(err)
	}
	if found {
		fmt.Print("The Id was found! are you sure you want to delete \n press 'd' for delete or 'r' for return : ")
		fmt.Scan(&innerOption)
		switch innerOption {
		case "d":
			DeleteStudent(index)
			commands.Pause("The student Deleted succesfully ")
			commands.ClearConsole()
			DocotrStartScreen()
		case "r":
			commands.Pause("Return to the start screen ")
			DocotrStartScreen()
		default:
			commands.Pause("Invalid Option...")
			goto startOfTheFunc
		}
	} else {
		commands.Pause("There is no such id on the file ! \n press 't' for try again \n press 'r' for retuen to main screen")
		fmt.Scan(&innerOption)
		switch innerOption {
		case "t":
			commands.ClearConsole()
			goto startOfTheFunc
              case "r":
			commands.ClearConsole()
			DocotrStartScreen()
		default:
			commands.Pause("Invalid option...")
		}
	}

}


func AddDoctor(newDoc basicdata.Doctor, filerelativePath string) {

	olddata, err2 := os.ReadFile(filerelativePath)
	if err2 != nil {
		print("from AddDoctor ", err2)
	}
	var dataJson []basicdata.Doctor
	json.Unmarshal(olddata, &dataJson)
	dataJson = append(dataJson, newDoc)

	jsonData, err := json.Marshal(dataJson)
	if err != nil {
		fmt.Println("Error marshaling doctor data:", err)
		return
	}
	fmt.Println("Doctor data:", string(jsonData))
	err = os.WriteFile(filerelativePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing doctor data to file:", err)
		return
	}
}
