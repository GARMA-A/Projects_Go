package main

func main(){

	todos := Todos{}
	storage := newStorage[Todos]("todos.json")
	storage.load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Excute(&todos)
	storage.save(todos)


}