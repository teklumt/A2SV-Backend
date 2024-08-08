package db

// import (
// 	" clean_architecture/domain"
// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func CreateTask(task domain.Task) (*mongo.InsertOneResult, error) {
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     return TaskCollection.InsertOne(ctx, task)
// }

// func GetTaskByID(id string) (*domain.Task, error) {
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     var task domain.Task
//     TaskCollection.FindOne(ctx, bson.M{"id": id}).Decode(&task)
//     if task == (domain.Task{}) {
// 		return nil, mongo.ErrNoDocuments
// 	}

//     return &task, nil
// }

// func UpdateTask(id string, task domain.Task) (*mongo.UpdateResult, error) {
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()
//     update := bson.M{
//         "$set": task,
//     }

//     return TaskCollection.UpdateOne(ctx, bson.M{"id": id}, update)
// }

// func DeleteTask(id string) (*mongo.DeleteResult, error) {
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     return TaskCollection.DeleteOne(ctx, bson.M{"id": id})
// }

// func GetAllData() (*mongo.Cursor, error) {
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     cursor, err := TaskCollection.Find(ctx, bson.M{})
//     if err != nil {
//         return nil, err
//     }

//     return cursor, nil
// }
