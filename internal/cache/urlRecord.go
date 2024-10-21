package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ribeirosaimon/shortify-read/config/db"
	"github.com/ribeirosaimon/shortify-read/internal/dto"
	"github.com/ribeirosaimon/tooltip/storage/redis"
	"github.com/ribeirosaimon/tooltip/tlog"
)

type UrlRecord interface {
	FindById(ctx context.Context, id string) (*dto.UrlRecord, error)
}

type urlRecordCache struct {
	urlCache redis.RConnInterface
}

func NewUrlRecord() *urlRecordCache {
	return &urlRecordCache{
		urlCache: db.NewRedisConnection(),
	}
}

func (u *urlRecordCache) FindById(ctx context.Context, id string) (*dto.UrlRecord, error) {
	tlog.Debug("NewUrlRecord.FindById", "Find Url in Redis")

	var wait sync.WaitGroup
	wait.Add(1)

	expirationTime := make(chan time.Duration, 1)
	go func() {
		ttl, err := u.urlCache.GetConnection().TTL(ctx, id).Result()
		if err != nil {
			tlog.Warn("NewUrlRecord.FindById.Gorotine1", err.Error())
		}
		expirationTime <- ttl
		wait.Done()
	}()

	urlInfo, err := u.urlCache.GetConnection().Get(ctx, id).Result()
	if err != nil {
		tlog.Warn("NewUrlRecord.FindById", err.Error())
		return nil, errors.New("url not found")
	}

	wait.Wait()
	duration := <-expirationTime

	return &dto.UrlRecord{
		Url:        urlInfo,
		Expiration: time.Now().Add(duration),
	}, nil
}
