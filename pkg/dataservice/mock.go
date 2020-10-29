package dataservice

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/CodingSquire/mai-monolit/pkg/api"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) SaveThesis(ctx context.Context,data api.Thesis)(result int, err error){
	args := m.Called(data)
	if a, ok := args.Get(0).(int); ok {
		return a, args.Error(1)
	}
	return result, args.Error(1)
}

func (m *MockClient) GetThesis(ctx context.Context,id int)(result api.Thesis, err error){
	args := m.Called(id)
	if a, ok := args.Get(0).(api.Thesis); ok {
		return a, args.Error(1)
	}
	return result, args.Error(1)
}
