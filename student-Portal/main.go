package main

import (
	"encoding/json"
	"os"
	basicdata "studentPortal/basicData"
	"studentPortal/commands"
	"studentPortal/doctor"
	"studentPortal/student"
)

// import "studentPortal/data"

// "fmt"
// "studentPortal/commands"
// "studentPortal/data"

/*****************************/
// some notes you create the basic data struct also insert it into
// student and doctor and create new functions that will
// take all basic data and  return struct of it your next metion
// is to implement it on the system to create newStudent newDoctor
// Good luck!
/*****************************/

func init() {
	// doctor.AddDoctor()
}


func main() {
	flag := commands.NewCmdFlag()
	id, sOrD, err := flag.StartAsStudentOrDoctor()
	var found bool = false
	if err != nil {
		print("error on read the flags of main", err)
		return
	}
	if sOrD == "s" {

		fileData, err := os.ReadFile("json/students.json")
		if err != nil {
			print("err on read file from the main read student", err)
			return
		}
		var jsonData []student.Student
		json.Unmarshal(fileData, &jsonData)
		for _, currentStudent := range jsonData {
			if currentStudent.Id == id {
				commands.Pause("Found succesfully loading your page....")
				student.GlobalCurrentStudent = currentStudent
				found = true
				break
			}
		}
		if found {
			student.StudentStartScreen()
		} else {
			println("The student not found on the file")
		}
	} else if sOrD == "d" {

		fileData, err := os.ReadFile("json/doctors.json")
		if err != nil {
			print("err on read file from the on main read docotor", err)
			return
		}
		var jsonData []basicdata.Doctor
		json.Unmarshal(fileData, &jsonData)
		for _, curDoctor := range jsonData {
			if curDoctor.BasicData.Id == id {
				commands.Pause("Found succesfully loading your page....")
				doctor.GlobalCurrentDoctor = curDoctor
				found = true
				break
			}
		}

		if found {
			doctor.DocotrStartScreen()
		} else {
			println("The doctor not found on the file")
		}
	} else {
		println("Invalid sOrD value from main")
	}

}
