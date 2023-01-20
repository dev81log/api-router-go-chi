package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Create a request to pass to our handler
	req, _ := http.NewRequest("GET", "/users", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and pass in the request and response recorder
	r := chi.NewRouter()
	r.Get("/users", getUsers)
	r.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect
	var users []User
	json.Unmarshal(rr.Body.Bytes(), &users)
	assert.Equal(t, 0, len(users))
}

func TestCreateUser(t *testing.T) {
	user := User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "mysecretpassword",
		Address:  "123 Main St",
		ZipCode:  "12345",
		State:    "NY",
	}

	body, _ := json.Marshal(user)

	// Create a request to pass to our handler
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and pass in the request and response recorder
	r := chi.NewRouter()
	r.Post("/users", createUser)
	r.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect
	var createdUser User
	json.Unmarshal(rr.Body.Bytes(), &createdUser)
	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.Address, createdUser.Address)
	assert.Equal(t, user.ZipCode, createdUser.ZipCode)
	assert.Equal(t, user.State, createdUser.State)
}

func TestUpdateUser(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "Jane Doe",
		Email:    "janedoe@example.com",
		Password: "mysecretpassword",
		Address:  "456 Park Ave",
		ZipCode:  "67890",
		State:    "CA",
	}
	users = append(users, user)

	body, _ := json.Marshal(user)

	// Create a request to pass to our handler
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(body))

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and pass in the request and response recorder
	r := chi.NewRouter()
	r.Put("/users/{id}", updateUser)
	r.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect
	var updatedUser User
	json.Unmarshal(rr.Body.Bytes(), &updatedUser)
	assert.Equal(t, user.Name, updatedUser.Name)
	assert.Equal(t, user.Email, updatedUser.Email)
	assert.Equal(t, user.Password, updatedUser.Password)
	assert.Equal(t, user.Address, updatedUser.Address)
	assert.Equal(t, user.ZipCode, updatedUser.ZipCode)
	assert.Equal(t, user.State, updatedUser.State)
}

func TestDeleteUser(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "mysecretpassword",
		Address:  "123 Main St",
		ZipCode:  "12345",
		State:    "NY",
	}
	users = append(users, user)

	// Create a request to pass to our handler
	req, _ := http.NewRequest("DELETE", "/users/1", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and pass in the request and response recorder
	r := chi.NewRouter()
	r.Delete("/users/{id}", deleteUser)
	r.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect
	var deletedUser User
	json.Unmarshal(rr.Body.Bytes(), &deletedUser)
	assert.Equal(t, user.ID, deletedUser.ID)
	assert.Equal(t, user.Name, deletedUser.Name)
	assert.Equal(t, user.Email, deletedUser.Email)
	assert.Equal(t, user.Password, deletedUser.Password)
	assert.Equal(t, user.Address, deletedUser.Address)
	assert.Equal(t, user.ZipCode, deletedUser.ZipCode)
	assert.Equal(t, user.State, deletedUser.State)
}
