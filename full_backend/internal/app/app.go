package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"myback/internal/api"
	"myback/internal/store"
	"myback/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, fmt.Errorf("app: NewApplication: %w", err)
	}

	err = store.MigrateFs(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// our stores will go here
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// our handlers will go here

	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}
	defer app.DB.Close()
	return app, nil
}

func (app *Application) HealthCheack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
