package main

import (
	"fmt"
	"studentPortal/data"
)




func main(){

	subjects ,err := data.SubjectsFromJsonToSlice()
	fmt.Println(err)
	fmt.Printf("%#v" , subjects)

}