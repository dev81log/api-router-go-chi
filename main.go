package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	ZipCode  string `json:"zipcode"`
	State    string `json:"state"`
}

var users []User

func main() {
	r := chi.NewRouter()

	r.Get("/users", getUsers)
	r.Post("/users", createUser)
	r.Put("/users/{id}", updateUser)
	r.Delete("/users/{id}", deleteUser)

	http.ListenAndServe(":8000", r)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user.ID = len(users) + 1
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	updatedUser.ID = id

	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
