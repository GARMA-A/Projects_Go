package student

import (
	"fmt"
	"os"
	basicdata "studentPortal/basicData"

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

// I need to show something whenever the program start with (-s id)
func StudentStartScreen(currentStudent Student) {
	fmt.Printf(`         Welcome %s , you are already stored in our memory
	 and your id is %s welcome back! .  
	 ---------------------------------------------------------
	 1) see your schedule
	 2) calculate GPA
	 3) see semester subjects
	 4) see any tasks asssign to you
	 ---------------------------------------------------------`+"\n",currentStudent.Name , currentStudent.Id)
}

// I need to print the current shedule for each student based on the
// current semester
func (s *Student) SeeYourShedule() {
	// TODO COMPLETE THIS FUNC
	t := table.New(os.Stdout)
	t.SetHeaders("1", "2", "3", "4", "5", "6", "7", "8", "9", "11", "12")

}

// This function should use by admin (doctor) only
// there is another func named WithLateCourses this is
// the method version of it for simplicity
func (s *Student) AddLateCourses(course ...SubjectName) {
	s.LateCourses = append(s.LateCourses, course...)
}
