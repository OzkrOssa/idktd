package endpoint

import (
	"context"

	usersv1 "github.com/OzkrOssa/idktd/api/proto/gen/users/v1"
	"github.com/OzkrOssa/idktd/internal/users/core/domain"
	"github.com/OzkrOssa/idktd/internal/users/core/port"
	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	RegisterEndpoint   endpoint.Endpoint
	GetUserEndpoint    endpoint.Endpoint
	ListUserEndpoint   endpoint.Endpoint
	UpdateUserEndpoint endpoint.Endpoint
	DeleteUserEndpoint endpoint.Endpoint
}

func NewEndpoint(us port.UserService) Endpoint {
	return Endpoint{
		RegisterEndpoint:   MakeRegisterEndpoint(us),
		GetUserEndpoint:    MakeGetUserEndpoint(us),
		ListUserEndpoint:   MakeListUserEndpoint(us),
		UpdateUserEndpoint: MakeUpdateUserEndpoint(us),
		DeleteUserEndpoint: MakeDeleteUserEndpoint(us),
	}
}

func MakeRegisterEndpoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		userReq, ok := request.(*usersv1.RegisterRequest)
		if !ok {
			return nil, err
		}

		user := &domain.User{
			Name:     userReq.Name,
			Email:    userReq.Email,
			Password: userReq.Password,
		}

		userRes, err := us.Register(ctx, user)
		if err != nil {
			return nil, err
		}

		return userRes, nil
	}
}

func MakeGetUserEndpoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(*usersv1.GetUserRequest)
		if !ok {
			return nil, err
		}

		userRes, err := us.GetUser(ctx, user.Id)
		if err != nil {
			return nil, err
		}

		return userRes, nil
	}
}

func MakeListUserEndpoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		listReq, ok := request.(*usersv1.ListUsersRequest)
		if !ok {
			return nil, err
		}

		userRes, err := us.ListUsers(ctx, listReq.Skip, listReq.Limit)
		if err != nil {
			return nil, err
		}

		return userRes, nil
	}
}

func MakeUpdateUserEndpoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*usersv1.UpdateUserRequest)
		if !ok {
			return nil, err
		}

		//TODO: fix it

		updateUser := &domain.User{
			ID:       req.Id,
			Name:     *req.Name,
			Email:    *req.Email,
			Password: *req.Password,
			Role:     domain.UserRole(*req.Role),
		}

		user, err := us.UpdateUser(ctx, updateUser)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeDeleteUserEndpoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*usersv1.DeleteUserRequest)
		if !ok {
			return nil, err
		}

		return nil, us.DeleteUser(ctx, req.Id)
	}
}
