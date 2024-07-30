package service

import (
	"LibraryManagement/model"
	"errors"
)

type LibraryManager interface {
    AddBook(book model.Book)
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []model.Book
    ListBorrowedBooks(memberID int) []model.Book

    AddMember(members model.Member)

}

type Library struct {
    books   map[int]model.Book
    members map[int]model.Member
    
    
}

func NewLibrary() *Library {
    return &Library{
        books:   make(map[int]model.Book),
        members: make(map[int]model.Member),
    }
}

func (l *Library) AddBook(book model.Book) {
    l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
    delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
    book, exists := l.books[bookID]
    if !exists || book.Status == "Borrowed" {
        return errors.New("book not available")
    }
    
    member, exists := l.members[memberID]
    if !exists {
        newMember := model.Member{ID: memberID}
        l.members[memberID] = newMember
    }
    
    book.Status = "Borrowed"
    l.books[bookID] = book
    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.members[memberID] = member
    
    return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
    member, exists := l.members[memberID]
    if !exists {
        return errors.New("member not found")
    }

    book, exists := l.books[bookID]
    if !exists || book.Status == "Available" {
        return errors.New("book not borrowed")
    }
    
    for i, borrowedBook := range member.BorrowedBooks {
        if borrowedBook.ID == bookID {
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            book.Status = "Available"
            l.books[bookID] = book
            l.members[memberID] = member
            return nil
        }
    }

    return errors.New("book not found in member's borrowed books")
}

func (l *Library) ListAvailableBooks() []model.Book {
    availableBooks := []model.Book{}
    for _, book := range l.books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []model.Book {
    member, exists := l.members[memberID]
    if !exists {
        return nil
    }
    return member.BorrowedBooks
}


func (l *Library) AddMember(member model.Member) {
    l.members[member.ID] = member
}

func (l *Library) GetMember(memberID int) (model.Member, error) {
    member, exists := l.members[memberID]
    if !exists {
        return model.Member{}, errors.New("member not found")
    }
    return member, nil
}