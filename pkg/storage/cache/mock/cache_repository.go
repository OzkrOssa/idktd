// Code generated by mockery v2.50.0. DO NOT EDIT.

package mock_cache

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// CacheRepository is an autogenerated mock type for the CacheRepository type
type CacheRepository struct {
	mock.Mock
}

type CacheRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *CacheRepository) EXPECT() *CacheRepository_Expecter {
	return &CacheRepository_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with no fields
func (_m *CacheRepository) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheRepository_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type CacheRepository_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *CacheRepository_Expecter) Close() *CacheRepository_Close_Call {
	return &CacheRepository_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *CacheRepository_Close_Call) Run(run func()) *CacheRepository_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CacheRepository_Close_Call) Return(_a0 error) *CacheRepository_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheRepository_Close_Call) RunAndReturn(run func() error) *CacheRepository_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, key
func (_m *CacheRepository) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CacheRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *CacheRepository_Expecter) Delete(ctx interface{}, key interface{}) *CacheRepository_Delete_Call {
	return &CacheRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, key)}
}

func (_c *CacheRepository_Delete_Call) Run(run func(ctx context.Context, key string)) *CacheRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CacheRepository_Delete_Call) Return(_a0 error) *CacheRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheRepository_Delete_Call) RunAndReturn(run func(context.Context, string) error) *CacheRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByPrefix provides a mock function with given fields: ctx, prefix
func (_m *CacheRepository) DeleteByPrefix(ctx context.Context, prefix string) error {
	ret := _m.Called(ctx, prefix)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByPrefix")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheRepository_DeleteByPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByPrefix'
type CacheRepository_DeleteByPrefix_Call struct {
	*mock.Call
}

// DeleteByPrefix is a helper method to define mock.On call
//   - ctx context.Context
//   - prefix string
func (_e *CacheRepository_Expecter) DeleteByPrefix(ctx interface{}, prefix interface{}) *CacheRepository_DeleteByPrefix_Call {
	return &CacheRepository_DeleteByPrefix_Call{Call: _e.mock.On("DeleteByPrefix", ctx, prefix)}
}

func (_c *CacheRepository_DeleteByPrefix_Call) Run(run func(ctx context.Context, prefix string)) *CacheRepository_DeleteByPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CacheRepository_DeleteByPrefix_Call) Return(_a0 error) *CacheRepository_DeleteByPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheRepository_DeleteByPrefix_Call) RunAndReturn(run func(context.Context, string) error) *CacheRepository_DeleteByPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key
func (_m *CacheRepository) Get(ctx context.Context, key string) ([]byte, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CacheRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *CacheRepository_Expecter) Get(ctx interface{}, key interface{}) *CacheRepository_Get_Call {
	return &CacheRepository_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *CacheRepository_Get_Call) Run(run func(ctx context.Context, key string)) *CacheRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CacheRepository_Get_Call) Return(_a0 []byte, _a1 error) *CacheRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CacheRepository_Get_Call) RunAndReturn(run func(context.Context, string) ([]byte, error)) *CacheRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value, ttl
func (_m *CacheRepository) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	ret := _m.Called(ctx, key, value, ttl)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte, time.Duration) error); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheRepository_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type CacheRepository_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value []byte
//   - ttl time.Duration
func (_e *CacheRepository_Expecter) Set(ctx interface{}, key interface{}, value interface{}, ttl interface{}) *CacheRepository_Set_Call {
	return &CacheRepository_Set_Call{Call: _e.mock.On("Set", ctx, key, value, ttl)}
}

func (_c *CacheRepository_Set_Call) Run(run func(ctx context.Context, key string, value []byte, ttl time.Duration)) *CacheRepository_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte), args[3].(time.Duration))
	})
	return _c
}

func (_c *CacheRepository_Set_Call) Return(_a0 error) *CacheRepository_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheRepository_Set_Call) RunAndReturn(run func(context.Context, string, []byte, time.Duration) error) *CacheRepository_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewCacheRepository creates a new instance of CacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCacheRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CacheRepository {
	mock := &CacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}