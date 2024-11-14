package main

import (
	"encoding/json"
	"os"
	basicdata "studentPortal/basicData"
	"studentPortal/commands"
	"studentPortal/doctor"
	"studentPortal/student"
	"time"
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
		print("error on read the flags of main" , err)
		return
	}
	if sOrD == "s" {

		fileData, err := os.ReadFile("json/students.json")
		if err != nil {
			print("err on read file from the main read student" , err)
			return
		}
		var jsonData []student.Student
		var foundedStudent student.Student
		json.Unmarshal(fileData, &jsonData)
		for _, obj := range jsonData {
			if obj.Id == id {
				println("Found succesfully loading your page....")
				time.Sleep(time.Second * 3)
				foundedStudent = obj
				student.StudentStartScreen(foundedStudent)
				found = true
				break
			}

		}
		if !found {
			println("The student not found on the file")
		}
	} else if sOrD == "d" {

		fileData, err := os.ReadFile("json/doctors.json")
		if err != nil {
			print("err on read file from the on main read docotor" , err)
			return
		}
		var jsonData []basicdata.Doctor
		var foundedDoctor basicdata.Doctor
		json.Unmarshal(fileData, &jsonData)
		for _, obj := range jsonData {
			if obj.BasicData.Id == id {
				println("Found succesfully loading your page....")
				time.Sleep(time.Second * 3)
				foundedDoctor = obj
				doctor.DocotrStartScreen(foundedDoctor)
				found = true
				break
			}
		}

		if !found {
			println("The doctor not found on the file")
		}
	} else {
		println("Invalid sOrD value from init on main")
	}

}
