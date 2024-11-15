package student

import (
	"encoding/json"
	"fmt"
	"os"
	basicdata "studentPortal/basicData"
	"studentPortal/commands"

	"github.com/aquasecurity/table"
)

type SubjectName string
type Grade string

type Student struct {
	Gpa             float32               `json:"Gpa,omitempty"`             // doctor privilege
	CurrentSemester int                   `json:"CurrentSemester,omitempty"` // student not optional
	HoursCompleted  int                   `json:"HoursCompleted,omitempty"`  // student not optional
	Grades          map[SubjectName]Grade `json:"Grades,omitempty"`          // doctor privilege
	LateCourses     []SubjectName         `json:"LateCourses,omitempty"`     // doctor privilege
	basicdata.BasicData
}

var GlobalCurrentStudent Student

// I need to show something whenever the program start with (-s id)
func StudentStartScreen() {
	fmt.Printf(`         Welcome %s , you are already stored in our memory
	 and your id is %s welcome back! .  
	 ---------------------------------------------------------
	 1) calculate GPA
	 2) see semester subjects
	 3) see any assignments asssign to you
	 4) Exit
	 --------------------------------------------------------`+"\n", GlobalCurrentStudent.Name, GlobalCurrentStudent.Id)
	var option string
	fmt.Scan(&option)
	OptionScreenForStudent(option)
}
func OptionScreenForStudent(option string) {
	switch option {
	case "2":
		SeeStudentSmesters()
	case "5":
		os.Exit(0)
	default:
		commands.Pause("Invalid Option...")
		StudentStartScreen()
	}

}

// I need to print the current shedule for each student based on the
// current semester
func (s *Student) SeeYourShedule() {
	// TODO COMPLETE THIS FUNC
	t := table.New(os.Stdout)
	t.SetHeaders("1", "2", "3", "4", "5", "6", "7", "8", "9", "11", "12")

}

func SeeStudentSmesters() {
startOfTheFunc:
	commands.ClearConsole()
	var semesterNumber int
	fmt.Print("Plaease enter the semmester you want to see : ")
	fmt.Scan(&semesterNumber)
	var subjects [][]string
	byteFileData, err := os.ReadFile("json/subjects.json")
	if err != nil {
		commands.Pause(err)
	}
	err = json.Unmarshal(byteFileData, &subjects)
	if err != nil {
		commands.Pause(err)
	}
	t := table.New(os.Stdout)
	t.SetHeaders("Subject Name", "Credet Hours")
	switch semesterNumber {
	case 0:
		addTheRows(t, subjects[0])
	case 1:
		addTheRows(t, subjects[1])
	case 2:
		addTheRows(t, subjects[2])
	case 3:
		addTheRows(t, subjects[3])
	case 4:
		addTheRows(t, subjects[4])
	case 5:
		addTheRows(t, subjects[5])
	case 6:
		addTheRows(t, subjects[6])
	case 7:
		addTheRows(t, subjects[7])
	case 8:
		addTheRows(t, subjects[8])
	default:
		commands.Pause("Invalid input...")
		goto startOfTheFunc
	}
	t.Render()
	commands.Pause("Return to the main screen ")
	commands.ClearConsole()
	StudentStartScreen()

}

func addTheRows(t *table.Table, subjects []string) {

	for _, el := range subjects {
		t.AddRow(el, "3.0")
	}
}

// This function should use by admin (doctor) only
// there is another func named WithLateCourses this is
// the method version of it for simplicity
func (s *Student) AddLateCourses(course ...SubjectName) {
	s.LateCourses = append(s.LateCourses, course...)
}
