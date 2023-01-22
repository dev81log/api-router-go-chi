![Asher_Duarte_rest_api_golang_code_thumb_7a0d8fd5-03e8-4ae6-bfb5-88eac4c9a15a (1)](https://user-images.githubusercontent.com/105469529/213918486-ab3741e9-212a-4700-ae54-72594bb3fefc.png)

# Go RESTful API (router go-chi)
This is a Go program that demonstrates basic RESTful API operations using the chi router library. It allows for creating, reading, updating and deleting of user data in a in-memory data store.

## Dependencies
+ Go
+ chi router library ([github.com/go-chi/chi/v5](url))

## Usage
1. Run `go get github.com/go-chi/chi/v5` to install the chi router library
2. Run `go run main.go` to run the program
3. Use a tool like [postman](https://www.postman.com/) to test the different routes and verify that the CRUD operations are working correctly

## Routes
`GET /users`: Retrieves all users

`POST /users`: Creates a new user with the following json format:
    
    `{
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password",
        "address": "123 Main St.",
        "zipcode": "12345",
        "state": "NY"
    }`
    
`PUT /users/{id}`: Updates a user with the specified ID using the same json format as above

`DELETE /users/{id}`: Deletes a user with the specified ID

## File Structure
`main.go`
This file contains the main function, and all the functions used to perform the CRUD operations, and the configuration of the router
