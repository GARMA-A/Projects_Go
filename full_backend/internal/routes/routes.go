package routes

import (
	"github.com/go-chi/chi/v5"

	"myback/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/health", app.HealthCheack)
	router.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
	router.Post("/workouts",
		app.WorkoutHandler.HandleCreateWorkout)
	return router
}
