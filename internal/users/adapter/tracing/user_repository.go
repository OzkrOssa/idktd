package tracing

import (
	"context"

	"github.com/OzkrOssa/idktd/internal/users/core/domain"
	"github.com/OzkrOssa/idktd/internal/users/core/port"
	"go.opentelemetry.io/otel"
	otelCodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type RepositoryLayerTracing struct {
	tracer trace.Tracer
	repo   port.UserRepository
}

func NewRepositoryLayerTracing(repo port.UserRepository) *RepositoryLayerTracing {
	return &RepositoryLayerTracing{
		otel.Tracer("RepositoryLayer"),
		repo,
	}
}

func (rl *RepositoryLayerTracing) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	ctx, span := rl.tracer.Start(ctx, "repository.createUser")
	defer span.End()

	createdUser, err := rl.repo.CreateUser(ctx, user)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return createdUser, nil
}

func (rl *RepositoryLayerTracing) GetUserByID(ctx context.Context, id uint64) (*domain.User, error) {
	ctx, span := rl.tracer.Start(ctx, "repository.getUserById")
	defer span.End()

	user, err := rl.repo.GetUserByID(ctx, id)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return user, nil
}

func (rl *RepositoryLayerTracing) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, span := rl.tracer.Start(ctx, "repository.getUserByEmail")
	defer span.End()

	user, err := rl.repo.GetUserByEmail(ctx, email)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return user, nil
}

func (rl *RepositoryLayerTracing) ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	ctx, span := rl.tracer.Start(ctx, "repository.listUsers")
	defer span.End()

	users, err := rl.repo.ListUsers(ctx, skip, limit)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return users, nil
}

func (rl *RepositoryLayerTracing) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	ctx, span := rl.tracer.Start(ctx, "repository.updateUser")
	defer span.End()

	userUpdated, err := rl.repo.UpdateUser(ctx, user)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return userUpdated, nil
}

func (rl *RepositoryLayerTracing) DeleteUser(ctx context.Context, id uint64) error {
	ctx, span := rl.tracer.Start(ctx, "repository.deleteUser")
	defer span.End()

	err := rl.repo.DeleteUser(ctx, id)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return err
	}

	return nil
}
