package basicdata

import (
	"time"
)
type BasicData struct {
	Name        string
	Id          string
	Age         int
	Gender      string // optional
	Address     string // optional
	DateOfBirth time.Time
	Phone       string // optional
}
type Doctor struct {
	BasicData
}


type OptionalArguments func(*BasicData)

func NewBasicData(name string, dateOfBirht time.Time, id string, stuentOptionalData ...OptionalArguments) *BasicData {
	BasicData := &BasicData{Name: name, DateOfBirth: dateOfBirht, Id: id}
	BasicData.Age = int(time.Now().Year() - dateOfBirht.Year())
	for _, option := range stuentOptionalData {
		option(BasicData)
	}
	return BasicData
}
func WithPhoneNumber(phone string) OptionalArguments {
	return func(b *BasicData) {
		b.Phone = phone
	}

}

func WithGender(gender string) OptionalArguments {
	return func(b *BasicData) {
		b.Gender = gender
	}
}

func WithAddress(address string) OptionalArguments {
	return func(b *BasicData) {
		b.Address = address
	}
}

// Doctor feature
// func WithLateCourses(subs ...subjectName) optionalArguments {
// 	return func(s *Student) {
// 		s.LateCourses = append(s.LateCourses, subs...)
// 	}
// }

// Doctor feature
// func WithGrades(grades map[subjectName]grade) optionalArguments {
// 	return func(s *Student) {
// 		if s.Grades == nil {
// 			s.Grades = make(map[subjectName]grade)
// 		}
// 		for k, v := range grades {

// 			s.Grades[k] = v
// 		}

// 	}
// }
