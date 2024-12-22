package test

import (
	"context"
	"testing"
	"time"

	usersv1 "github.com/OzkrOssa/idktd/api/proto/gen/users/v1"
	"github.com/OzkrOssa/idktd/internal/users/adapter/endpoint"
	"github.com/OzkrOssa/idktd/internal/users/core/domain"
	"github.com/OzkrOssa/idktd/internal/users/core/port/mock"
	"github.com/OzkrOssa/idktd/internal/users/core/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestUserEndpoint_Register(t *testing.T) {
	ctx := context.Background()

	name := gofakeit.Name()
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 15)

	request := &usersv1.RegisterRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	expectedServiceInput := &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	hashedPassword, _ := util.HashPassword(password)

	expectedResponse := &domain.User{
		ID:        gofakeit.Uint64(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      domain.RoleReader,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testCases := []struct {
		desc     string
		mock     func(svc *mock.UserService)
		input    *usersv1.RegisterRequest
		expected struct {
			response interface{}
			err      error
		}
	}{
		{
			desc: "Success",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().
					Register(ctx, expectedServiceInput).
					Return(expectedResponse, nil)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: expectedResponse,
				err:      nil,
			},
		},
		{
			desc: "Conflict data",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().
					Register(ctx, expectedServiceInput).
					Return(nil, domain.ErrConflictingData)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrConflictingData,
			},
		},
		{
			desc: "Internal error",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().
					Register(ctx, expectedServiceInput).
					Return(nil, domain.ErrInternal)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockSvc := mock.NewUserService(t)
			tc.mock(mockSvc)

			ep := endpoint.NewEndpoint(mockSvc)

			resp, err := ep.RegisterEndpoint(ctx, tc.input)

			assert.Equal(t, tc.expected.response, resp)
			assert.Equal(t, tc.expected.err, err)
		})
	}
}

func TestUserEndpoint_GetUser(t *testing.T) {
	ctx := context.Background()

	userID := gofakeit.Uint64()
	password := gofakeit.Password(true, true, true, true, false, 15)

	request := &usersv1.GetUserRequest{
		Id: userID,
	}

	hashedPassword, _ := util.HashPassword(password)
	expectedResponse := &domain.User{
		ID:        userID,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Password:  hashedPassword,
		Role:      domain.RoleReader,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	testCases := []struct {
		desc     string
		mock     func(svc *mock.UserService)
		input    *usersv1.GetUserRequest
		expected struct {
			response interface{}
			err      error
		}
	}{
		{
			desc: "Success",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().
					GetUser(ctx, userID).
					Return(expectedResponse, nil)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: expectedResponse,
				err:      nil,
			},
		},
		{
			desc: "Not found",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().
					GetUser(ctx, userID).
					Return(nil, domain.ErrDataNotFound)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrDataNotFound,
			},
		},
		{
			desc: "Internal error",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().
					GetUser(ctx, userID).
					Return(nil, domain.ErrInternal)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockSvc := mock.NewUserService(t)
			tc.mock(mockSvc)

			ep := endpoint.NewEndpoint(mockSvc)

			resp, err := ep.GetUserEndpoint(ctx, tc.input)

			assert.Equal(t, tc.expected.response, resp)
			assert.Equal(t, tc.expected.err, err)
		})
	}
}

func TestUserEndpoint_ListUsers(t *testing.T) {
	ctx := context.Background()

	skip := gofakeit.Uint64()
	limit := gofakeit.Uint64()

	request := &usersv1.ListUsersRequest{
		Skip:  skip,
		Limit: limit,
	}

	var users []domain.User
	for i := 0; i < 10; i++ {
		password := gofakeit.Password(true, true, true, true, false, 15)
		hashedPassword, _ := util.HashPassword(password)

		users = append(users, domain.User{
			ID:       gofakeit.Uint64(),
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Role:     domain.RoleReader,
			Password: hashedPassword,
		})
	}

	testCases := []struct {
		desc     string
		mock     func(svc *mock.UserService)
		input    *usersv1.ListUsersRequest
		expected struct {
			response interface{}
			err      error
		}
	}{
		{
			desc: "Success",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().ListUsers(ctx, skip, limit).Return(users, nil)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: users,
				err:      nil,
			},
		},
		{
			desc: "Not users in list",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().ListUsers(ctx, skip, limit).Return([]domain.User{}, nil)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: []domain.User{},
				err:      nil,
			},
		},
		{
			desc: "Internal error",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().ListUsers(ctx, skip, limit).Return(nil, domain.ErrInternal)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockSvc := mock.NewUserService(t)
			tc.mock(mockSvc)

			ep := endpoint.NewEndpoint(mockSvc)

			resp, err := ep.ListUserEndpoint(ctx, tc.input)

			assert.Equal(t, tc.expected.response, resp)
			assert.Equal(t, tc.expected.err, err)
		})
	}
}

func TestUserEndpoint_UpdateUser(t *testing.T) {
	ctx := context.Background()

	userID := gofakeit.Uint64()
	name := gofakeit.Name()
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 15)
	hashedPassword, _ := util.HashPassword(password)

	role := usersv1.Role_ROLE_ADMIN

	request := &usersv1.UpdateUserRequest{
		Id:       userID,
		Name:     &name,
		Email:    &email,
		Password: &password,
		Role:     &role,
	}

	input := &domain.User{
		ID:       userID,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     domain.RoleAdmin,
	}

	expectedResponse := &domain.User{
		ID:        userID,
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      domain.RoleAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testCases := []struct {
		desc     string
		mock     func(svc *mock.UserService)
		input    *usersv1.UpdateUserRequest
		expected struct {
			response interface{}
			err      error
		}
	}{
		{
			desc: "Success",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().UpdateUser(ctx, input).Return(expectedResponse, nil)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: expectedResponse,
				err:      nil,
			},
		},
		{
			desc: "Not found",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().UpdateUser(ctx, input).Return(nil, domain.ErrDataNotFound)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrDataNotFound,
			},
		},
		{
			desc: "No update data",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().UpdateUser(ctx, input).Return(nil, domain.ErrNoUpdatedData)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrNoUpdatedData,
			},
		},
		{
			desc: "Internal error",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().UpdateUser(ctx, input).Return(nil, domain.ErrInternal)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockSvc := mock.NewUserService(t)
			tc.mock(mockSvc)

			ep := endpoint.NewEndpoint(mockSvc)

			resp, err := ep.UpdateUserEndpoint(ctx, tc.input)

			assert.Equal(t, tc.expected.response, resp)
			assert.Equal(t, tc.expected.err, err)
		})
	}
}

func TestUserEndpoint_Delete(t *testing.T) {
	ctx := context.Background()

	userID := gofakeit.Uint64()
	request := &usersv1.DeleteUserRequest{
		Id: userID,
	}

	testCases := []struct {
		desc     string
		mock     func(svc *mock.UserService)
		input    *usersv1.DeleteUserRequest
		expected struct {
			response interface{}
			err      error
		}
	}{
		{
			desc: "Success",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().DeleteUser(ctx, userID).Return(nil)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      nil,
			},
		},
		{
			desc: "Not found",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().DeleteUser(ctx, userID).Return(domain.ErrDataNotFound)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrDataNotFound,
			},
		},
		{
			desc: "Internal error",
			mock: func(svc *mock.UserService) {
				svc.EXPECT().DeleteUser(ctx, userID).Return(domain.ErrInternal)
			},
			input: request,
			expected: struct {
				response interface{}
				err      error
			}{
				response: nil,
				err:      domain.ErrInternal,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockSvc := mock.NewUserService(t)
			tc.mock(mockSvc)

			ep := endpoint.NewEndpoint(mockSvc)

			resp, err := ep.DeleteUserEndpoint(ctx, tc.input)

			assert.Equal(t, tc.expected.response, resp)
			assert.Equal(t, tc.expected.err, err)
		})
	}
}
