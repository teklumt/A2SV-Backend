package mongo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockCollection is a mock implementation of the Collection interface.
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(SingleResult)
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	args := m.Called(ctx, document)
	return args.Get(0), args.Error(1)
}

func (m *MockCollection) InsertMany(ctx context.Context, documents []interface{}) ([]interface{}, error) {
	args := m.Called(ctx, documents)
	return args.Get(0).([]interface{}), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(Cursor), args.Error(1)
}

func (m *MockCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCollection) Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error) {
	args := m.Called(ctx, pipeline)
	return args.Get(0).(Cursor), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}
