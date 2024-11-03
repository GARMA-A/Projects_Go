package student

import (
	"fmt"
	"os"

	"github.com/aquasecurity/table"
)



type Student struct {
	name string
	id string
	gpa float32
	currentSemester int

}



func StudentStartScreen(name , id string){
        fmt.Printf(`         Welcome %s , you are already stored in our memory
	 and your id is %s welcome back! .  
	 ---------------------------------------------------------
	 1) see your schedule
	 2) calculate GPA
	 3) see semester subjects
	 4) see any tasks asssign to you
	 ---------------------------------------------------------`+"\n" , name ,id)
}



func (s *Student) SeeYourShedule(){

   t := table.New(os.Stdout)
   t.SetHeaders("1","2","3","4","5","6","7","8","9","11","12")	

}



