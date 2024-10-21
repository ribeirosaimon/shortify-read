package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ribeirosaimon/tooltip/testutils/tcontainer"
	"github.com/ribeirosaimon/tooltip/tserver"
	"github.com/stretchr/testify/assert"
)

const (
	urlId            = "12345"
	urlOriginalValue = "https://test.io"
)

func TestUrlRecordCache(t *testing.T) {
	ctx := context.Background()
	container := tcontainer.NewRedisTestContainer()
	err := container.Start()
	assert.Nil(t, err)
	tserver.NewMockEnvironment(tserver.MockEnvironment{RedisHost: container.GetHost()})

	err = inputUrlRecordInCache(ctx, tserver.GetRedisConfig().Host)
	assert.Nil(t, err)

	serviceCache := NewUrlRecord()
	for _, v := range []struct {
		testName string
		hasError bool
	}{
		{
			testName: "have to run",
			hasError: false,
		},
	} {
		t.Run(v.testName, func(t *testing.T) {
			urlRecord, err := serviceCache.FindById(ctx, urlId)
			if v.hasError {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, urlOriginalValue, urlRecord.Url)
		})
	}
}

// inputUrlRecordInCache adds a URL record to the Redis cache before executing tests.
func inputUrlRecordInCache(ctx context.Context, host string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: host,
		DB:   0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	if err = rdb.Set(ctx, urlId, urlOriginalValue, 1*time.Hour).Err(); err != nil {
		fmt.Println("Error setting key:", err)
		return err
	}

	return nil
}
