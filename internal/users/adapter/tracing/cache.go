package tracing

import (
	"context"
	"errors"
	"time"

	"github.com/OzkrOssa/idktd/pkg/storage/cache"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	otelCodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type CacheLayerTracing struct {
	tracer trace.Tracer
	cache  cache.CacheRepository
}

func NewCacheLayerTracing(cache cache.CacheRepository) *CacheLayerTracing {
	return &CacheLayerTracing{
		otel.Tracer("CacheLayer"),
		cache,
	}
}

func (cl *CacheLayerTracing) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	ctx, span := cl.tracer.Start(ctx, "cache.set")
	defer span.End()

	err := cl.cache.Set(ctx, key, value, ttl)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return err
	}

	return nil
}
func (cl *CacheLayerTracing) Get(ctx context.Context, key string) ([]byte, error) {
	ctx, span := cl.tracer.Start(ctx, "cache.get")
	defer span.End()

	b, err := cl.cache.Get(ctx, key)

	if err != nil {
		if err == redis.Nil {
			span.AddEvent("key not found")
			return nil, errors.New("key not fount")
		}
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return nil, err
	}

	return b, nil
}
func (cl *CacheLayerTracing) Delete(ctx context.Context, key string) error {
	ctx, span := cl.tracer.Start(ctx, "cache.delete")
	defer span.End()

	err := cl.cache.Delete(ctx, key)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return err
	}

	return nil
}
func (cl *CacheLayerTracing) DeleteByPrefix(ctx context.Context, prefix string) error {
	ctx, span := cl.tracer.Start(ctx, "cache.deleteByPrefix")
	defer span.End()

	err := cl.cache.DeleteByPrefix(ctx, prefix)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())

		return err
	}

	return nil
}
func (cl *CacheLayerTracing) Close() error {
	return cl.cache.Close()
}
