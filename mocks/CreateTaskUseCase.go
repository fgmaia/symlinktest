// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"

	input "github.com/fgmaia/task/internal/usecases/ports/input"

	mock "github.com/stretchr/testify/mock"

	output "github.com/fgmaia/task/internal/usecases/ports/output"
)

// CreateTaskUseCase is an autogenerated mock type for the CreateTaskUseCase type
type CreateTaskUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, createTask
func (_m *CreateTaskUseCase) Execute(ctx context.Context, createTask *input.CreateTaskInput) (*output.CreateTaskOutput, error) {
	ret := _m.Called(ctx, createTask)

	var r0 *output.CreateTaskOutput
	if rf, ok := ret.Get(0).(func(context.Context, *input.CreateTaskInput) *output.CreateTaskOutput); ok {
		r0 = rf(ctx, createTask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.CreateTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *input.CreateTaskInput) error); ok {
		r1 = rf(ctx, createTask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
