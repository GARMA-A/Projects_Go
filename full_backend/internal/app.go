package internal

import (
	"log"
	"os"
)

type Application struct {
	logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &Application{
		logger: logger,
	}
	return app, nil
}
