# 📋 Task Management REST API

## 🌟 Overview

This is a Task Management REST API built with Go and the Gin framework. The API includes JWT-based user authentication, allowing users to create accounts, log in, create tasks, and manage them. Only admin users can delete all tasks.

## ✨ Features

- 📝 User Registration and Login with JWT authentication
- 📌 Task creation, updating, and deletion
- 🔒 Role-based access control (admin and user roles)

## 🔧 Environment Variables

Create a `.env` file in the root directory of your project with the following variables:

```env
LOCAL_SERVER_PORT=:8080
MONGODB_URL=mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
JWT_SECRET=your_jwt_secret
```

## 🚀 Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/task-management-api.git
   ```
2. Navigate to the project directory:
   ```sh
   cd task-management-api
   ```
3. Install the required Go modules:
   ```sh
   go mod tidy
   ```

## 🏃 Running the Application

1. Ensure MongoDB is running and accessible with the provided connection string in `.env`.
2. Run the application:
   ```sh
   go run main.go
   ```

## 📚 API Endpoints

### 🔑 User Authentication

- **POST** `/auth/register` - Register a new user

  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "message": "User registered successfully"
    }
    ```

- **POST** `/auth/login` - Log in a user
  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "token": "jwt_token"
    }
    ```

### 🗂️ Task Management

- **GET** `/tasks` - Get all tasks (admin only)

  - Response:
    ```json
    [
      {
        "id": "string",
        "title": "string",
        "description": "string",
        "status": "string"
      }
    ]
    ```

- **POST** `/tasks` - Create a new task (authenticated users)

  - Request Body:
    ```json
    {
      "title": "string",
      "description": "string",
      "status": "string"
    }
    ```
  - Response:
    ```json
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "status": "string"
    }
    ```

- **PUT** `/tasks/:id` - Update a task (authenticated users)

  - Request Body:
    ```json
    {
      "title": "string",
      "description": "string",
      "status": "string"
    }
    ```
  - Response:
    ```json
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "status": "string"
    }
    ```

- **DELETE** `/tasks/:id` - Delete a task (authenticated users)

  - Response:
    ```json
    {
      "message": "Task deleted successfully"
    }
    ```

- **DELETE** `/tasks` - Delete all tasks (admin only)
  - Response:
    ```json
    {
      "message": "All tasks deleted successfully"
    }
    ```

## 🛡️ Middleware

- **JWT Middleware**: Protects routes by verifying JWT tokens.

## 🗂️ Project Structure

```
.
├── auth
│   ├── auth.go
│   └── jwt.go
├── controllers
│   ├── taskController.go
│   └── userController.go
├── db
│   ├── connectionDB.go
│   └── dbOperations.go
├── middleware
│   └── authMiddleware.go
├── models
│   ├── task.go
│   └── user.go
├── routes
│   ├── taskRoutes.go
│   └── userRoutes.go
├── utils
│   └── hash.go
├── .env
├── go.mod
├── go.sum
└── main.go
```
