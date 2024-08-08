package repository

import (
	"clean_architecture/domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
    collection *mongo.Collection
}

func NewUserRepositoryImpl(coll *mongo.Collection) domain.UserRepository {
    return &UserRepositoryImpl{
        collection: coll,
    }
}

func (ur *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
    user.Role = "user"
    _, err := ur.collection.InsertOne(context.Background(), user)
    return user, err
}




// type userRepository struct {
// 	database   mongo.Database
// 	collection string
// }

// func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
// 	return &userRepository{
// 		database:   db,
// 		collection: collection,
// 	}
// }

// func (ur *userRepository) Create(c context.Context, user *domain.User) error {
// 	collection := ur.database.Collection(ur.collection)

// 	_, err := collection.InsertOne(c, user)

// 	return err
// }













// type UserRepositoryImpl struct {
// 	users 
// }

// func NewUserRepositoryImpl() domain.UserRepository {

// 	return &UserRepositoryImpl{users: make(map[string]domain.User)}
// }

// func (repo *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
// 	if _, exists := repo.users[user.Username]; exists {
// 		return domain.User{}, errors.New("user already exists")
// 	}
// 	user.Role = "user"
// 	repo.users[user.Username] = user
// 	return user, nil
// }
