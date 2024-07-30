package controllers

import (
	"LibraryManagement/model"
	"LibraryManagement/service"
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
)

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls") 
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Run() {
	library := service.NewLibrary()
	library.AddBook(model.Book{ID: 1, Title: "1984", Author: "George Orwell", Status: "Available"})
	library.AddBook(model.Book{ID: 2, Title: "The Hobbit", Author: "J.R.R. Tolkien", Status: "Available"})
	
	library.AddMember(model.Member{ID: 1, Name: "Alice"})

	for {
		fmt.Println("\n*********************************************")
		fmt.Println("*  Console-Based Library Management System   *")
		fmt.Println("*********************************************")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			clearConsole()
			var id int
			var title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter book author: ")
			fmt.Scan(&author)
			library.AddBook(model.Book{ID: id, Title: title, Author: author, Status: "Available"})
			fmt.Println("\nBook added successfully!")
		case 2:
			clearConsole()
			var id int
			fmt.Print("Enter book ID to remove: ")
			fmt.Scan(&id)
			library.RemoveBook(id)
			fmt.Println("\nBook removed successfully!")
		case 3:
			clearConsole()
			var bookID, memberID int
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			if err := library.BorrowBook(bookID, memberID); err != nil {
				fmt.Println("\nError:", err)
			} else {
				fmt.Println("\nBook borrowed successfully!")
			}
		case 4:
			clearConsole()
			var bookID, memberID int
			fmt.Print("Enter book ID to return: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("\nError:", err)
			} else {
				fmt.Println("\nBook returned successfully!")
			}
		case 5:
			clearConsole()
			books := library.ListAvailableBooks()
			fmt.Println("\nAvailable Books:")
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
			fmt.Fprintln(w, "ID\tTitle\tAuthor\t")
			for _, book := range books {
				fmt.Fprintf(w, "%d\t%s\t%s\t\n", book.ID, book.Title, book.Author)
			}
			w.Flush()
		case 6:
			clearConsole()
			var memberID int
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("\nBorrowed Books:")
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
			fmt.Fprintln(w, "ID\tTitle\tAuthor\t")
			for _, book := range books {
				fmt.Fprintf(w, "%d\t%s\t%s\t\n", book.ID, book.Title, book.Author)
			}
			w.Flush()
		case 7:
			clearConsole()
			fmt.Println("\nExiting...")
			return
		default:
			fmt.Println("\nInvalid choice, please try again.")
		}
	}
}
