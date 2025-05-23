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

<table>
  <thead>
    <tr>
      <th>Engi</th>
      <th>Gorilla Mux</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>
        <pre><code class="language-go">
func (api *NotesAPI) Routers() engi.Routes {
  return engi.Routes{
    // ...
    engi.GET("{id}"): engi.Handle(
      api.Get,
      path.Integer("id", validate.Greater(0)),
      middlewares.Description("get note by id"),
    ),
    // ...
  }
}
        </code></pre>
      </td>
      <td>
        <pre><code class="language-go">
func (api *TasksAPI) RegisterRoutes(r *mux.Router) {
  r.HandleFunc("/tasks/{id:[0-9]+}", api.Get).Methods("GET")
}
        </code></pre>
      </td>
    </tr>
  </tbody>
</table>

As shown above, Engi provides more functionality on the router level allowing us to define parameters with validation, generate documentation, etc.

- **Second Example** - parameter parsing:

<table>
  <thead>
    <tr>
      <th>Engi</th>
      <th>Gorilla Mux</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>
        <pre><code class="language-go">
// All parameters that didn't pass validation, that can't be parsed to integer 
//	or didn't found in path will cause BadRequest error to be returned.

var id = request.Integer("id", placing.InPath)
        </code></pre>
      </td>
      <td>
        <pre><code class="language-go">
var ctx = r.Context()

idStr := mux.Vars(r)["id"]

id, err := strconv.ParseInt(idStr, 10, 64)
if err != nil {
  http.Error(w, "invalid id", http.StatusBadRequest)
  return
}
        </code></pre>
      </td>
    </tr>
  </tbody>
</table>

As shown above, Engi's way of defining parameters is more concise and easier to understand even if it's less flexible.

- **Third Example** - response creation:

<table>
  <thead>
    <tr>
      <th>Engi</th>
      <th>Gorilla Mux</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>
        <pre><code class="language-go">
// Object and marshaler already set in engine or in service 
// and all handlers of this service will return responses with same formatting

return response.OK(note)
        </code></pre>
      </td>
      <td>
        <pre><code class="language-go">
w.WriteHeader(http.StatusCreated) // Optional for 200, mandatory for others
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(task)
        </code></pre>
      </td>
    </tr>
  </tbody>
</table>

As this example shows, Engi's way of creating responses is more concise and easier to understand and maintain consistent formatting of the service’s responses. 
