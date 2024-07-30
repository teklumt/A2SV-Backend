`

# Console-Based Library Management System

This is a simple console-based library management system implemented in Go. The system allows you to add, remove, borrow, and return books, as well as list available and borrowed books. The console interface provides an easy way to manage a small library.

## Features

- Add new books to the library.
- Remove books from the library.
- Borrow books from the library.
- Return borrowed books to the library.
- List all available books in the library.
- List all borrowed books by a member.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/teklumt/A2SV-Backend-Tasks-2024/tree/main/Task3
   cd Task3
   ```

   `

2. Build the project:

   ```bash
   go build
   ```

3. Run the application:

   ```bash
   ./Task3
   ```

## Usage

Upon running the application, you will see the following menu:

`

---

- Console-Based Library Management System \*

---

1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books
7. Exit
   Enter your choice:

```

### Menu Options

1. **Add Book**: Adds a new book to the library. You will be prompted to enter the book ID, title, and author.
2. **Remove Book**: Removes a book from the library. You will be prompted to enter the book ID.
3. **Borrow Book**: Borrows a book from the library. You will be prompted to enter the book ID and member ID.
4. **Return Book**: Returns a borrowed book to the library. You will be prompted to enter the book ID and member ID.
5. **List Available Books**: Lists all available books in the library in a tabular format.
6. **List Borrowed Books**: Lists all borrowed books by a member in a tabular format. You will be prompted to enter the member ID.
7. **Exit**: Exits the application.

## Code Structure

- **controllers**: Contains the main application logic and the `Run` function which drives the menu.
- **model**: Contains the data models for `Book` and `Member`.
- **service**: Contains the `Library` service which manages the books and members.

`


```
