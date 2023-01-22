# Go RESTful API Example
This is a Go program that demonstrates basic RESTful API operations using the chi router library. It allows for creating, reading, updating and deleting of user data in a in-memory data store.

## Dependencies
+ Go
+ chi router library [GitHub Pages](github.com/go-chi/chi/v5)
Usage
Run go get github.com/go-chi/chi/v5 to install the chi router library
Run go run main.go to run the program
Use a tool like Postman to test the different routes and verify that the CRUD operations are working correctly
Routes
GET /users: Retrieves all users
POST /users: Creates a new user with the following json format:
Copy code
    {
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password",
        "address": "123 Main St.",
        "zipcode": "12345",
        "state": "NY"
    }
PUT /users/{id}: Updates a user with the specified ID using the same json format as above
DELETE /users/{id}: Deletes a user with the specified ID
File Structure
Copy code
main.go
This file contains the main function, and all the functions used to perform the CRUD operations, and the configuration of the router
