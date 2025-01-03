// Code generated by mockery v2.50.0. DO NOT EDIT.

package mock

import (
	context "context"

	domain "github.com/OzkrOssa/idktd/internal/users/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

type UserService_Expecter struct {
	mock *mock.Mock
}

func (_m *UserService) EXPECT() *UserService_Expecter {
	return &UserService_Expecter{mock: &_m.Mock}
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserService) DeleteUser(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserService_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type UserService_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *UserService_Expecter) DeleteUser(ctx interface{}, id interface{}) *UserService_DeleteUser_Call {
	return &UserService_DeleteUser_Call{Call: _e.mock.On("DeleteUser", ctx, id)}
}

func (_c *UserService_DeleteUser_Call) Run(run func(ctx context.Context, id uint64)) *UserService_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *UserService_DeleteUser_Call) Return(_a0 error) *UserService_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserService_DeleteUser_Call) RunAndReturn(run func(context.Context, uint64) error) *UserService_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *UserService) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*domain.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type UserService_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *UserService_Expecter) GetUser(ctx interface{}, id interface{}) *UserService_GetUser_Call {
	return &UserService_GetUser_Call{Call: _e.mock.On("GetUser", ctx, id)}
}

func (_c *UserService_GetUser_Call) Run(run func(ctx context.Context, id uint64)) *UserService_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *UserService_GetUser_Call) Return(_a0 *domain.User, _a1 error) *UserService_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetUser_Call) RunAndReturn(run func(context.Context, uint64) (*domain.User, error)) *UserService_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// ListUsers provides a mock function with given fields: ctx, skip, limit
func (_m *UserService) ListUsers(ctx context.Context, skip uint64, limit uint64) ([]domain.User, error) {
	ret := _m.Called(ctx, skip, limit)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) ([]domain.User, error)); ok {
		return rf(ctx, skip, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) []domain.User); ok {
		r0 = rf(ctx, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64) error); ok {
		r1 = rf(ctx, skip, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_ListUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUsers'
type UserService_ListUsers_Call struct {
	*mock.Call
}

// ListUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - skip uint64
//   - limit uint64
func (_e *UserService_Expecter) ListUsers(ctx interface{}, skip interface{}, limit interface{}) *UserService_ListUsers_Call {
	return &UserService_ListUsers_Call{Call: _e.mock.On("ListUsers", ctx, skip, limit)}
}

func (_c *UserService_ListUsers_Call) Run(run func(ctx context.Context, skip uint64, limit uint64)) *UserService_ListUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64))
	})
	return _c
}

func (_c *UserService_ListUsers_Call) Return(_a0 []domain.User, _a1 error) *UserService_ListUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_ListUsers_Call) RunAndReturn(run func(context.Context, uint64, uint64) ([]domain.User, error)) *UserService_ListUsers_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: ctx, user
func (_m *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) (*domain.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type UserService_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - ctx context.Context
//   - user *domain.User
func (_e *UserService_Expecter) Register(ctx interface{}, user interface{}) *UserService_Register_Call {
	return &UserService_Register_Call{Call: _e.mock.On("Register", ctx, user)}
}

func (_c *UserService_Register_Call) Run(run func(ctx context.Context, user *domain.User)) *UserService_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.User))
	})
	return _c
}

func (_c *UserService_Register_Call) Return(_a0 *domain.User, _a1 error) *UserService_Register_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_Register_Call) RunAndReturn(run func(context.Context, *domain.User) (*domain.User, error)) *UserService_Register_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) (*domain.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserService_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *domain.User
func (_e *UserService_Expecter) UpdateUser(ctx interface{}, user interface{}) *UserService_UpdateUser_Call {
	return &UserService_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, user)}
}

func (_c *UserService_UpdateUser_Call) Run(run func(ctx context.Context, user *domain.User)) *UserService_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.User))
	})
	return _c
}

func (_c *UserService_UpdateUser_Call) Return(_a0 *domain.User, _a1 error) *UserService_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_UpdateUser_Call) RunAndReturn(run func(context.Context, *domain.User) (*domain.User, error)) *UserService_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
