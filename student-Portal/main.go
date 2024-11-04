package main

import (
	"fmt"
	"studentPortal/commands"
	// "studentPortal/data"
)




func main(){

	cmd := commands.NewCmdFlag()
	id ,ch , err := cmd.StartAsStudentOrDoctor()
	fmt.Println(err)
	fmt.Println(id ,ch)

}