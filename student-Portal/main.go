package main

import "studentPortal/commands"

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


func init(){
	flag := commands.NewCmdFlag()
	id,sOrD,err := flag.StartAsStudentOrDoctor()
	if err != nil {
		return 
	}
	if sOrD == "s"{

	}


}



func main() {

	// cmd := commands.NewCmdFlag()
	// id ,ch , err := cmd.StartAsStudentOrDoctor()
	// fmt.Println(err)
	// fmt.Println(id ,ch)
}
