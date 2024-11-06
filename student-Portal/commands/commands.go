package commands

import (
	"errors"
	"flag"
)

type Commands struct {
	student string
	doctor  string
}

// take the -s or -d and the id for entering student or doctor mode
func NewCmdFlag() *Commands {

	newCommand := Commands{}

	flag.StringVar(&newCommand.student, "s", "", "Define You are student and check the id")
	flag.StringVar(&newCommand.doctor, "d", "", "Define You are Doctor and check the id")

	flag.Parse()

	return &newCommand

}
// method from commands struct to start the student or doctor mode
func (c *Commands) StartAsStudentOrDoctor() (string, string, error) {
	switch {
	case c.student != "":
		return c.student, "s", nil
	case c.doctor != "":
		return c.doctor, "d", nil
	default:
		return "", "nop!", errors.New("no id was passed")
	}
}
