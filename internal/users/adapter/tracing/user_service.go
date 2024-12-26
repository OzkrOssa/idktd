package tracing

import (
	"context"

	"github.com/OzkrOssa/idktd/internal/users/core/domain"
	"github.com/OzkrOssa/idktd/internal/users/core/port"
	"go.opentelemetry.io/otel"
	otelCodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type ServiceLayerTracing struct {
	tracer trace.Tracer
	svc    port.UserService
}

func NewServiceLayerTracing(svc port.UserService) *ServiceLayerTracing {
	return &ServiceLayerTracing{otel.Tracer("ServiceLayer"), svc}
}

func (t *ServiceLayerTracing) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	ctx, span := t.tracer.Start(ctx, "service.register")
	defer span.End()

	rUser, err := t.svc.Register(ctx, user)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return rUser, nil
}

func (t *ServiceLayerTracing) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	ctx, span := t.tracer.Start(ctx, "service.getUsers")
	defer span.End()

	user, err := t.svc.GetUser(ctx, id)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return user, nil
}

func (t *ServiceLayerTracing) ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	ctx, span := t.tracer.Start(ctx, "service.listUsers")
	defer span.End()

	users, err := t.svc.ListUsers(ctx, skip, limit)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return users, nil
}

func (t *ServiceLayerTracing) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	ctx, span := t.tracer.Start(ctx, "service.updateUser")
	defer span.End()

	userUpdated, err := t.svc.UpdateUser(ctx, user)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return userUpdated, nil
}

func (t *ServiceLayerTracing) DeleteUser(ctx context.Context, id uint64) error {
	ctx, span := t.tracer.Start(ctx, "service.delete")
	defer span.End()

	err := t.svc.DeleteUser(ctx, id)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return err
	}

	return nil
}
