# Golang-API

This repository contains a simple golang-API to manage CRUD operations on users and their friends.

To use the API, follow these instructions:
Clone the repository into your desired local directory
```
git clone https://github.com/tommywiffler/golang-api.git
```
Navigate to the project directory and run the following commands to clean up any dependencies
```
go mod tidy
go mod vendor
```
Run the API server
```
go run main.go
```

# API's

The API endpoints are exposed at the following URI
```
http://127.0.0.1:8080/api/
```

## To create a user
```
POST /users
```
```json
{
    "id": 3333,
    "name": "Todd",
    "email": "todd@email.com",
    "age": 45,
    "friends": [
        1234, 1212
    ] 
}
```

## To get all users
```
GET /users
```

## To get a specific user
```
GET /users/{id}
```

## To update a user
```
PATCH /users/{id}
```
```json
{
    "name": "Dodd",
    "email": "dodd@email.com",
    "age": 46
}
```

## To delete a user
```
DELETE /users/{id}
```

## To get friends of a user
```
GET /users/{id}/friends
```

## To add a friend to a user
```
POST /users/{id}/friends
```
```json
{
    "id": 5678
}
```

## To delete a friend of a user
```
DELETE /users/{id}/friends/{id}
```

# Testing

There is a Postman Collection of example requests included in the repository that can be imported to a Postman environment and used for testing.