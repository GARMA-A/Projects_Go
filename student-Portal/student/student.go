package student

import (
	"fmt"
	"os"
	basicdata "studentPortal/basicData"

	"github.com/aquasecurity/table"
)



type subjectName string
type grade string

type Student struct {
	Gpa             float32               // doctor // Done
	CurrentSemester int                   // not optional // Done
	HoursCompleted  string                // not optional
	Grades          map[subjectName]grade // doctor
	LateCourses     []subjectName         // doctor
	basicdata.BasicData
}

func StudentStartScreen(name, id string) {
	fmt.Printf(`         Welcome %s , you are already stored in our memory
	 and your id is %s welcome back! .  
	 ---------------------------------------------------------
	 1) see your schedule
	 2) calculate GPA
	 3) see semester subjects
	 4) see any tasks asssign to you
	 ---------------------------------------------------------`+"\n", name, id)
}


func (s *Student) SeeYourShedule() {
	// TODO COMPLETE THIS FUNC
	t := table.New(os.Stdout)
	t.SetHeaders("1", "2", "3", "4", "5", "6", "7", "8", "9", "11", "12")

}
// this is ok as it is not constructor pass by student
func (s *Student) AddLateCourses(course ...subjectName) {
	s.LateCourses = append(s.LateCourses, course...)
}
