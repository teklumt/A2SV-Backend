# 📋 Task Management REST API

This Task Management REST API, developed using Go and the Gin framework, supports basic CRUD operations for managing tasks using mongoDB Database. This documentation provides an overview of the API, its endpoints, and how to use it.

🔗 **[API Documentation](https://documenter.getpostman.com/view/32898780/2sA3kd9cJG)**

## 📁 Folder Structure

```
task_manager/
├── main.go
├── controllers/
│   └── controller.go
├── model/
│   └── task.go
├── db/
│   └── storage.go
├── router/
│   └── router.go
├── services/
│   └── service.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

## 💡 Implementation Details

### `main.go`

Entry point of the application. It sets up the router and starts the server.

### `controllers/controller.go`

Handles incoming HTTP requests and invokes the appropriate service methods.

### `model/task.go`

Defines the data structures used in the application, specifically the `Task` struct.

### `db/storage.go`

Contains the database connection and operations for MongoDB.

### `router/router.go`

Sets up the routes and initializes the Gin router.

### `services/service.go`

Contains business logic and data manipulation functions.

### `docs/api_documentation.md`

Contains API documentation and other related documentation.

## 🌐 API Endpoints

### 📜 GET /tasks

Fetches a list of all tasks.

**Response:**

```json
{
  "tasks": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "This is Task 1",
      "status": "Pending"
    },
    {
      "id": "2",
      "title": "Task 2",
      "description": "This is Task 2",
      "status": "Completed"
    },
    {
      "id": "3",
      "title": "Task 3",
      "description": "This is Task 3",
      "status": "Pending"
    },
    {
      "id": "4",
      "title": "Task 4",
      "description": "This is Task 4",
      "status": "Completed"
    }
  ]
}
```

### 🔍 GET /tasks/:id

Fetches the details of a specific task by its ID.

**Response:**

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "This is Task 1",
  "status": "Pending"
}
```

### ✏️ PUT /tasks/:id

Updates a specific task.

**Request Body:**

```json
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "status": "Completed"
}
```

**Response:**

```json
{
  "message": "Successfully Updated",
  "task": {
    "id": "1",
    "title": "Updated Task Title",
    "description": "Updated Task Description",
    "status": "Completed"
  }
}
```

### 🗑️ DELETE /tasks/:id

Deletes a specific task.

**Response:**

```json
{
  "message": "Successfully Deleted"
}
```

### ➕ POST /tasks

Creates a new task.

**Request Body:**

```json
{
  "title": "New Task Title",
  "description": "New Task Description",
  "status": "Pending"
}
```

**Response:**

```json
{
  "message": "Successfully Added",
  "task": {
    "id": "5",
    "title": "New Task Title",
    "description": "New Task Description",
    "status": "Pending"
  }
}
```

## ⚠️ Error Handling

- **404 Not Found:** When the specified resource is not found.
- **400 Bad Request:** When the request payload is invalid.
- **500 Internal Server Error:** For unexpected server errors.

## 🧪 Testing

Utilize Postman to test each endpoint of the Task Management API. Ensure that the API returns the correct responses for various scenarios.

## 🛠️ Instructions to Run the API

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go mod tidy` to install the dependencies.
4. Set up MongoDB Atlas and configure the connection string in the `db/storage.go` file.
5. Start the server using `go run main.go`.
6. Use Postman or curl to interact with the API.

## 🏁 Conclusion

This Task Management REST API, built using Go and the Gin framework, provides a robust and easy-to-use interface for managing tasks. It includes endpoints for creating, retrieving, updating, and deleting tasks, with MongoDB Atlas as the storage solution and detailed error handling. Use the provided instructions to set up and test the API, and refer to the example responses to understand the expected outputs for each endpoint.
