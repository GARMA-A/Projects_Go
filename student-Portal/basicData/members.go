package basicdata

import (
	"time"
)
type BasicData struct {
	Name        string    `json:"name,omitempty"`
	Id          string    `json:"id,omitempty"`
	Age         int       `json:"age,omitempty"`
	Gender      string    `json:"gender,omitempty"`      // Optional field
	Address     string    `json:"address,omitempty"`     // Optional field
	DateOfBirth time.Time `json:"dateOfBirth,omitempty"`
	Phone       string    `json:"phone,omitempty"`       // Optional field
}

type Doctor struct {
	BasicData
}


type OptionalArguments func(*BasicData)


// basic data that is on doctor or student with must! parameters and optional ones
func NewBasicData(name string, dateOfBirht time.Time, id string, stuentOptionalData ...OptionalArguments) *BasicData {
	BasicData := &BasicData{Name: name, DateOfBirth: dateOfBirht, Id: id}
	BasicData.Age = int(time.Now().Year() - dateOfBirht.Year())
	for _, option := range stuentOptionalData {
		option(BasicData)
	}
	return BasicData
}

// basic data (Phone) optional  func to pass for student of doctor
func WithPhoneNumber(phone string) OptionalArguments {
	return func(b *BasicData) {
		b.Phone = phone
	}
}

// basic data (Gender) optional func to pass for student docotr
func WithGender(gender string) OptionalArguments {
	return func(b *BasicData) {
		b.Gender = gender
	}
}


// basic data (Address) optional func to pass for doctor
func WithAddress(address string) OptionalArguments {
	return func(b *BasicData) {
		b.Address = address
	}
}
