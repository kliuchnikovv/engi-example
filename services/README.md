# Services Package

This package contains examples of two similar services built using the same principles:
**CRUD APIs with database interaction.**

## Overview

- Each service demonstrates a typical CRUD (Create, Read, Update, Delete) API.
- Services interact with the database through dedicated store packages.

## Included Services

- **NotesAPI**: CRUD operations for notes written using Engi.
- **TasksAPI**: CRUD operations for tasks written using Gorrilla Mux. 

Both services follow the same structure and can serve as example of key advantages of Engi:
- **Structured Code**: Engi promotes a structured approach to code, making it easier to understand and maintain.
- **Automatic Error Handling**: Engi automatically handles errors, reducing the need for manual error checking.
- **Parameter Parsing**: Engi simplifies parameter parsing, making it easier to extract and validate data from requests.
- **Middleware Integration**: Engi integrates with middleware, allowing for easy addition of features like logging and authentication.

## Comparison
- **First Example** - routes definition based on simple `get record by id ` method:

| Engi                                           | Gorilla Mux                                                  |
|-----------------------------------------------|--------------------------------------------------------------|
| ```go                                         | ```go                                                        |
| func (api *NotesAPI) Routers() engi.Routes {  | func (api *TasksAPI) RegisterRoutes(r *mux.Router) {         |
|     return engi.Routes{                       |     r.HandleFunc("/tasks/{id:[0-9]+}", api.Get).Methods("GET") |
|         ...                                   | }                                                            |
|         engi.GET("{id}"): engi.Handle(        | ```                                                          |
|             api.Get,                          |                                                              |
|             path.Integer("id", validate.Greater(0)), |                                                              |
|             middlewares.Description("get note by id"), |                                                              |
|         ...                                   |                                                              |
|     }                                         |                                                              |
| }                                             |                                                              |
| ```                                           |                                                              |
