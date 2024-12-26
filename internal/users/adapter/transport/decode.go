package transport

import (
	"context"

	usersv1 "github.com/OzkrOssa/idktd/api/proto/gen/users/v1"
	"github.com/bufbuild/protovalidate-go"
	"go.opentelemetry.io/otel"
	otelCodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func decodeRegisterRequest(ctx context.Context, request interface{}) (interface{}, error) {
	ctx, span := otel.Tracer("TransportLayer").Start(ctx, "transport.decodeRegisterRequest")
	defer span.End()

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
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	return req, nil
}

func decodeGetUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	ctx, span := otel.Tracer("TransportLayer").Start(ctx, "transport.decodeGetUserRequest")
	defer span.End()

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
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	return req, nil
}

func decodeListUsersRequest(ctx context.Context, request interface{}) (interface{}, error) {
	ctx, span := otel.Tracer("TransportLayer").Start(ctx, "transport.decodeListUsersRequest")
	defer span.End()

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
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	return req, nil
}

func decodeUpdateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	ctx, span := otel.Tracer("TransportLayer").Start(ctx, "transport.decodeUpdateUserRequest")
	defer span.End()

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
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	return req, nil
}

func decodeDeleteUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	ctx, span := otel.Tracer("TransportLayer").Start(ctx, "transport.decodeDeleteUserRequest")
	defer span.End()

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
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	if err := validator.Validate(req); err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	return req, nil
}
