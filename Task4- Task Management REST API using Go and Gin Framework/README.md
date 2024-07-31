# Task Management REST API

## Objective

The objective of this task is to create a simple Task Management REST API using Go programming language and Gin Framework. This API supports basic CRUD operations for managing tasks.

## Requirements

- Implement a REST API with the following endpoints:
  - `GET /tasks`: Get a list of all tasks.
  - `GET /tasks/:id`: Get the details of a specific task.
  - `PUT /tasks/:id`: Update a specific task. This endpoint accepts a JSON body with the new details of the task.
  - `DELETE /tasks/:id`: Delete a specific task.
  - `POST /tasks`: Create a new task. This endpoint accepts a JSON body with the task's title, description, due date, and status.
- Use an in-memory database to store tasks.
- Ensure proper error handling and response codes for different scenarios.
- Provide clear and concise documentation for each endpoint, including expected request payloads and response formats.
- Utilize Postman to test each endpoint of the Task Management API.

## Folder Structure

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

## Implementation

### main.go

Entry point of the application.

### controllers/controller.go

Handles incoming HTTP requests and invokes the appropriate service methods.

### model/task.go

Defines the data structures used in the application.

### db/storage.go

Contains the in-memory database.

### router/router.go

Sets up the routes and initializes the Gin router.

### services/service.go

Contains business logic and data manipulation functions.

### docs/api_documentation.md

Contains API documentation and other related documentation.

## API Endpoints

### GET /tasks

**Description:** Get a list of all tasks.

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

### GET /tasks/:id

**Description:** Get the details of a specific task.

**Response:**

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "This is Task 1",
  "status": "Pending"
}
```

### PUT /tasks/:id

**Description:** Update a specific task.

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
  "message": "Successfully Updated"
}
```

### DELETE /tasks/:id

**Description:** Delete a specific task.

**Response:**

```json
{
  "message": "Successfully Deleted"
}
```

### POST /tasks

**Description:** Create a new task.

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

## Error Handling

- **404 Not Found:** When the specified resource is not found.
- **400 Bad Request:** When the request payload is invalid.
- **500 Internal Server Error:** For unexpected server errors.

## Testing

- Utilize Postman to test each endpoint of the Task Management API.
- Ensure that the API returns the correct responses for various scenarios.

## Instructions to Run the API

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go mod tidy` to install the dependencies.
4. Start the server using `go run main.go`.
5. Use Postman or curl to interact with the API.
