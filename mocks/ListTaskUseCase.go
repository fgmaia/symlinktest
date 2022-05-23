// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/fgmaia/task/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// ListTaskUseCase is an autogenerated mock type for the ListTaskUseCase type
type ListTaskUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, userID, found
func (_m *ListTaskUseCase) Execute(ctx context.Context, userID string, found func(*entities.Task) error) error {
	ret := _m.Called(ctx, userID, found)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, func(*entities.Task) error) error); ok {
		r0 = rf(ctx, userID, found)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}