package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramWorkoutID := chi.URLParam(r, "id")
	if paramWorkoutID == "" {
		http.Error(w, "Missing workout ID", http.StatusBadRequest)
		return
	}
	workoutID, err := strconv.ParseInt(paramWorkoutID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "this is the workoutID %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "created workout")
}
