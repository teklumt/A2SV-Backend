package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// type User struct {
// 	ID       primitive.ObjectID `bson:"_id,omitempity" json:"id" `
// 	Email    string             `json:"email"`
// 	Password string             `json:"password"`
// 	Role     string             `json:"role"`
// }
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string 		   `json:"username" bson:"username" validate:"required,min=3,max=50"`	
	Password string `json:"password" bson:"password" validate:"required,min=6,max=50"`
	Role string `json:"role" bson:"role"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	LoginUser(username string, password string) (User, error)
	// GetUserByEmail(email string) (User, error)
	// GetUserByID(id int) (User, error)
	// DeleteUser(id int) error
	// UpdateUser(user User) (User, error)
	// Login(email string, password string) (User, error)
}

// type PasswordService interface {
// 	HashPassword(password string) (string, error)
// }
