package transport

import (
	"context"

	usersv1 "github.com/OzkrOssa/idktd/api/proto/gen/users/v1"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func decodeRegisterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*usersv1.RegisterRequest)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payload from client")
	}

	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&usersv1.RegisterRequest{},
		),
	)

	if err != nil {
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*usersv1.GetUserRequest)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payload from client")
	}

	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&usersv1.GetUserRequest{},
		),
	)

	if err != nil {
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeListUsersRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*usersv1.ListUsersRequest)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payload from client")
	}

	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&usersv1.ListUsersRequest{},
		),
	)

	if err != nil {
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*usersv1.UpdateUserRequest)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payload from client")
	}

	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&usersv1.UpdateUserRequest{},
		),
	)

	if err != nil {
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*usersv1.DeleteUserRequest)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payload from client")
	}

	validator, err := protovalidate.New(
		protovalidate.WithMessages(
			&usersv1.DeleteUserRequest{},
		),
	)

	if err != nil {
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}
