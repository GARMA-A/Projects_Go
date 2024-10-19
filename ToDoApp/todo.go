package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/aquasecurity/table"
)


type Todo  struct{
	Title  string
	completed bool
	createdAt time.Time 
	completedAt *time.Time
}

type Todos []Todo



func (todos *Todos) add(title string){
	todo := Todo{
		Title: title,
		completed: false,
		createdAt:  time.Now(),
		completedAt: nil,
	}
	*todos = append(*todos, todo)
}


func (todos *Todos) validatIendex(	index int) error {

	if index > 0   ||  index >=len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
  return nil
}


func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validatIendex(index); err != nil {
		return err
	}

	*todos = append(t[:index] , t[index+1:]...)

	return nil
}


func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validatIendex(index); err != nil {
		return err
	}


	isCompleted  := t[index].completed

	if !isCompleted{
		completionTime := time.Now()
		t[index].completedAt = &completionTime 
	}

	 t[index].completed = !isCompleted

	return nil
} 


func (todos *Todos) edit(index int, title string) error {

	t := *todos

	if err := t.validatIendex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}


func (todos *Todos) print()  {

	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#" ,"Title","Check" , "CreatedAt" , "CompletedAt")

	for index , t := range  *todos {
		completed := "NotDone"
		completeAt := ""

		if t.completed{
			completed = "Done"
			if t.completedAt != nil {
				completeAt = t.completedAt.Format(time.RFC1123)

				
			}
		}

           table.AddRow(strconv.Itoa(index) , t.Title , completed, t.createdAt.Format(time.RFC1123), completeAt)



	}

	table.Render()


}