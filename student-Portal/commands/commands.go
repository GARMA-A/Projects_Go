package commands

import (
	"errors"
	"flag"
)

type Commands struct {
	student string
	doctor  string
}

func NewCmdFlag() *Commands {

	newCommand := Commands{}

	flag.StringVar(&newCommand.student, "s", "", "Define You are student and check the id")
	flag.StringVar(&newCommand.doctor, "d", "", "Define You are Doctor and check the id")

	flag.Parse()

	return &newCommand

}

func (c *Commands) StartAsStudentOrDoctor() (string , string , error) {
	switch {
	case c.student != "":
		return c.student,"s",nil
	case c.doctor != "":
		return c.doctor,"d",nil
	default:
		return "","nop!", errors.New("no id was passed")
	}
}
