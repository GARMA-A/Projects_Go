package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloHandler( w http.ResponseWriter , r *http.Request){

	if r.URL.Path != "/hello"{
		http.Error(w , "404 not found " , http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w , "method not found " , http.StatusNotFound)
		return
	}
   
	fmt.Fprint(w , "hello")
}


func formHandler(w http.ResponseWriter , r * http.Request ){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w , "Parseform() err %v" ,err)
		return 
	}
	fmt.Fprintf(w, "POST request sucesful")
	name := r.FormValue("name")
	address := r.FormValue("address")
      fmt.Fprintf(w , "Name = %s \n " , name)
      fmt.Fprintf(w , "Address = %s \n " , address)

}


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/" , fileServer)
	http.HandleFunc("/form" , formHandler)
	http.HandleFunc("/hello" , helloHandler)

	fmt.Printf("Starting our server at port 8000\n")
	if err := http.ListenAndServe(":8000",nil) ; err != nil {
		log.Fatal(err)

	}

}