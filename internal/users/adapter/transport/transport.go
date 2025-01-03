package transport

import (
	"context"

	usersv1 "github.com/OzkrOssa/idktd/api/proto/gen/users/v1"
	"github.com/OzkrOssa/idktd/internal/users/adapter/endpoint"
	"github.com/OzkrOssa/idktd/internal/users/core/domain"
	"github.com/go-kit/kit/transport"
	gt "github.com/go-kit/kit/transport/grpc"
	kitlog "github.com/go-kit/log"
	"go.opentelemetry.io/otel"
	otelCodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcTransport struct {
	RegisterHandler   gt.Handler
	GetUserHandler    gt.Handler
	ListUsersHandler  gt.Handler
	UpdateUserHandler gt.Handler
	DeleteUserHandler gt.Handler
	usersv1.UnimplementedUserServiceServer
	tracer trace.Tracer
}

func NewGrpcTransport(endpoint endpoint.Endpoint, logger kitlog.Logger) usersv1.UserServiceServer {
	opts := gt.ServerErrorHandler(transport.NewLogErrorHandler(logger))
	return &grpcTransport{
		RegisterHandler:   gt.NewServer(endpoint.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, opts),
		GetUserHandler:    gt.NewServer(endpoint.GetUserEndpoint, decodeGetUserRequest, encodeGetUserResponse, opts),
		ListUsersHandler:  gt.NewServer(endpoint.ListUserEndpoint, decodeListUsersRequest, encodeListUsersResponse, opts),
		UpdateUserHandler: gt.NewServer(endpoint.UpdateUserEndpoint, decodeUpdateUserRequest, encodeUpdateUserResponse, opts),
		DeleteUserHandler: gt.NewServer(endpoint.DeleteUserEndpoint, decodeDeleteUserRequest, encodeDeleteUserResponse, opts),
		tracer:            otel.Tracer("TransportLayer"),
	}
}

func (g *grpcTransport) Register(ctx context.Context, request *usersv1.RegisterRequest) (*usersv1.RegisterResponse, error) {
	ctx, span := g.tracer.Start(ctx, "transport.register")
	defer span.End()

	_, resp, err := g.RegisterHandler.ServeGRPC(ctx, request)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		switch err {
		case domain.ErrConflictingData:
			return nil, status.Errorf(codes.AlreadyExists, "%s", err.Error())
		case domain.ErrInternal:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.RegisterResponse), nil
}
func (g *grpcTransport) GetUser(ctx context.Context, request *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	ctx, span := g.tracer.Start(ctx, "transport.getUser")
	defer span.End()

	_, resp, err := g.GetUserHandler.ServeGRPC(ctx, request)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		switch err {
		case domain.ErrDataNotFound:
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		case domain.ErrInternal:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.GetUserResponse), nil

}
func (g *grpcTransport) ListUsers(ctx context.Context, request *usersv1.ListUsersRequest) (*usersv1.ListUsersResponse, error) {
	ctx, span := g.tracer.Start(ctx, "transport.listUsers")
	defer span.End()

	_, resp, err := g.ListUsersHandler.ServeGRPC(ctx, request)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		switch err {
		case domain.ErrInternal:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.ListUsersResponse), nil
}
func (g *grpcTransport) UpdateUser(ctx context.Context, request *usersv1.UpdateUserRequest) (*usersv1.UpdateUserResponse, error) {
	ctx, span := g.tracer.Start(ctx, "transport.updateUsers")
	defer span.End()

	_, resp, err := g.UpdateUserHandler.ServeGRPC(ctx, request)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		switch err {
		case domain.ErrDataNotFound:
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		case domain.ErrNoUpdatedData:
			return nil, status.Errorf(codes.FailedPrecondition, "%s", err.Error())
		case domain.ErrInternal:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.UpdateUserResponse), nil
}
func (g *grpcTransport) DeleteUser(ctx context.Context, request *usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error) {
	ctx, span := g.tracer.Start(ctx, "transport.deleteUsers")
	defer span.End()

	_, resp, err := g.DeleteUserHandler.ServeGRPC(ctx, request)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		switch err {
		case domain.ErrDataNotFound:
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		case domain.ErrInternal:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
		}
	}

	return resp.(*usersv1.DeleteUserResponse), nil
}
