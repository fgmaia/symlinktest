// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	taskpb "github.com/fgmaia/task/pb/taskpb"
)

// TaskServiceClient is an autogenerated mock type for the TaskServiceClient type
type TaskServiceClient struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: ctx, in, opts
func (_m *TaskServiceClient) CreateTask(ctx context.Context, in *taskpb.CreateTaskRequest, opts ...grpc.CallOption) (*taskpb.CreateTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *taskpb.CreateTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *taskpb.CreateTaskRequest, ...grpc.CallOption) *taskpb.CreateTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taskpb.CreateTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taskpb.CreateTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTask provides a mock function with given fields: ctx, in, opts
func (_m *TaskServiceClient) FindTask(ctx context.Context, in *taskpb.FindTaskRequest, opts ...grpc.CallOption) (*taskpb.FindTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *taskpb.FindTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *taskpb.FindTaskRequest, ...grpc.CallOption) *taskpb.FindTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taskpb.FindTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taskpb.FindTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: ctx, in, opts
func (_m *TaskServiceClient) ListTasks(ctx context.Context, in *taskpb.ListTaskRequest, opts ...grpc.CallOption) (taskpb.TaskService_ListTasksClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 taskpb.TaskService_ListTasksClient
	if rf, ok := ret.Get(0).(func(context.Context, *taskpb.ListTaskRequest, ...grpc.CallOption) taskpb.TaskService_ListTasksClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(taskpb.TaskService_ListTasksClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taskpb.ListTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *TaskServiceClient) Login(ctx context.Context, in *taskpb.LoginRequest, opts ...grpc.CallOption) (*taskpb.LoginResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *taskpb.LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, *taskpb.LoginRequest, ...grpc.CallOption) *taskpb.LoginResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taskpb.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taskpb.LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadImage provides a mock function with given fields: ctx, opts
func (_m *TaskServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (taskpb.TaskService_UploadImageClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 taskpb.TaskService_UploadImageClient
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) taskpb.TaskService_UploadImageClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(taskpb.TaskService_UploadImageClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
