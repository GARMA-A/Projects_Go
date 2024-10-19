package main

import (
	"encoding/json"
	"os"
)




type Storage[T any] struct{
	FileName string
}




func newStorage[T any] (fileName string) *Storage[T] {

	return  &Storage[T]{FileName: fileName}

}

func (s *Storage[T]) save(data T) error {
	fileData , err := json.MarshalIndent(data,"","   ")

	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName,fileData,0644)

}


func (s *Storage[T]) load(data *T) error {
	fileData , err := os.ReadFile(s.FileName)
	if err != nil {
		return err 
	}

	return json.Unmarshal(fileData , data)
}








