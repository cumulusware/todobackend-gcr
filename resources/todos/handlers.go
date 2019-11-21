package todos

import (
	"encoding/json"
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
func ReadAll(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := ds.GetAll()
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, todos)
	}
}

// Create handles the POST method to create a new todo.
func Create(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Todo
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&t); err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		_, err := ds.Create(&t)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		helpers.RespondWithJSON(w, http.StatusCreated, t)
	}
}

// DeleteAll handles the DELETE method to delete all todos.
func DeleteAll(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := ds.DeleteAll(); err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		helpers.RespondWithJSON(w, http.StatusNoContent, nil)
	}
}
