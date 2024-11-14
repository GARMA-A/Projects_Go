package commands

import (
	"errors"
	"flag"
	"os"
	"os/exec"
)

type Commands struct {
	student string
	doctor  string
}

//  Note that only work for linux 
func ClearConsole() {
	var cmd  *exec.Cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
func (c *Commands) StartAsStudentOrDoctor() (id string,sOrD string,err error) {
	switch {
	case c.student != "":
		return c.student, "s", nil
	case c.doctor != "":
		return c.doctor, "d", nil
	default:
		return "", "nop!", errors.New("no id was passed")
	}
}
