package doctor

import (
	basicdata "studentPortal/basicData"
	"studentPortal/student"
	"time"
)

type optionalStaffStudentByDoctor func(*student.Student)

// optional func ! the function version of the AddLateCourses method
// which add any number of late courses
func WithLateCourses(subs ...student.SubjectName) optionalStaffStudentByDoctor {
	return func(s *student.Student) {
		s.LateCourses = append(s.LateCourses, subs...)
	}
}

// optional func ! add some grades to the student specific course
func WithGrades(grades map[*student.SubjectName]student.Grade) optionalStaffStudentByDoctor {
	return func(s *student.Student) {
		for k, v := range grades {
			s.Grades[*k] = v
		}

	}
}



// Not complete yet finish it you have a location on json called students to append the newStudent to it  
func addStudent(name string, dateOfBirht time.Time, id string, GPA float32, CurrentSemester int, HoursCompleted int, LateCourses []student.SubjectName, stuentOptionalData ...basicdata.OptionalArguments) {

	myBasicData := basicdata.NewBasicData(name, dateOfBirht, id, stuentOptionalData...)
	myStudent := student.Student{BasicData: *myBasicData, Gpa: GPA, CurrentSemester: CurrentSemester, HoursCompleted: HoursCompleted, LateCourses: LateCourses}
	print(myStudent)
	// jsonData,err  := json.MarshalIndent(myBasicData , "" ,"   ")
	// if err != nil {
	// 	return
	// }

	// fileData, err2 := os.ReadFile("../json/students.json")
	// if err2 != nil {
	// 	return
	// }

	// // Unmarshal existing JSON data into a slice of Person objects
	// var existingData []student.Student
	// if len(fileData) > 0 { // Check if file is not empty
	// 	err = json.Unmarshal(fileData, &existingData)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// // Append new data to the existing data
	// existingData = append(existingData, newData)

}
