package logging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/OzkrOssa/idktd/internal/users/core/domain"
	"github.com/OzkrOssa/idktd/internal/users/core/port"
	kitlog "github.com/go-kit/log"
)

type LoggingService struct {
	logger kitlog.Logger
	svc    port.UserService
}

func NewLoggingService(logger kitlog.Logger, svc port.UserService) port.UserService {
	return &LoggingService{logger: logger, svc: svc}
}

func (ls *LoggingService) Register(ctx context.Context, user *domain.User) (u *domain.User, err error) {
	defer func(begin time.Time) {
		b, _ := json.Marshal(user)
		ls.logger.Log(
			"method", "register",
			"user", b,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.svc.Register(ctx, user)
}

func (ls *LoggingService) GetUser(ctx context.Context, id uint64) (u *domain.User, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "get_user",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.svc.GetUser(ctx, id)
}

func (ls *LoggingService) ListUsers(ctx context.Context, skip, limit uint64) (ulist []domain.User, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "list_users",
			"skip", skip,
			"limit", limit,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.svc.ListUsers(ctx, skip, limit)
}

func (ls *LoggingService) UpdateUser(ctx context.Context, user *domain.User) (u *domain.User, err error) {
	defer func(begin time.Time) {
		b, _ := json.Marshal(user)
		ls.logger.Log(
			"method", "update_user",
			"updated_user", b,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.svc.UpdateUser(ctx, user)
}

func (ls *LoggingService) DeleteUser(ctx context.Context, id uint64) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "delete_user",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.svc.DeleteUser(ctx, id)
}
