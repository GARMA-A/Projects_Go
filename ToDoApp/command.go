package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {

	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new Todo plz specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit Todo by index plz specify title")
	flag.IntVar(&cf.Del, "del", -1, "Specifiy Todo index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specifiy Todo index to delet")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Excute(todos *Todos) {

	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error invalid format for edit please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error : invalid index for edit ")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")

	}

}
