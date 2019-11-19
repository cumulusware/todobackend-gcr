package todos

import (
	"net/http"
	"todobackend-gcr/helpers"
)

// DescribeAll handles the OPTIONS method for the todos/ endpoint.
func DescribeAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithOptions(w, "GET,POST,DELETE,OPTIONS")
	}
}

// ReadAll handles the GET method to list all todos.
func ReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusOK, "Hello")
	}
}

// Create handles the POST method to create a new todo.
func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := Todo{
			Title: "a todo",
		}
		helpers.RespondWithJSON(w, http.StatusOK, todo)
	}
}

// DeleteAll handles the DELETE method to delete all todos.
func DeleteAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusNoContent, "")
	}
}
