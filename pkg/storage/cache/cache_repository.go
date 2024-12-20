package cache

import (
	"context"
	"time"
)

//go:generate mockery --name=CacheRepository --case=underscore --output=mock --outpkg=mock_cache --with-expecter=true
type CacheRepository interface {
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	DeleteByPrefix(ctx context.Context, prefix string) error
	Close() error
}
