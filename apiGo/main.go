package main

import (
	"fmt"
	"net/http"

	"apigo/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("welcome to the go server")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}
