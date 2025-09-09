package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"myback/internal/app"
	"myback/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go server port")
	flag.Parse()

	myapp, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	myapp.Logger.Println("Application started on port", port)
	router := routes.SetupRoutes(myapp)
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}
	err = server.ListenAndServe()
	if err != nil {
		myapp.Logger.Fatal(err)
	}
}
