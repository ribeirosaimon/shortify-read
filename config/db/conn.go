package db

import (
	"github.com/ribeirosaimon/tooltip/storage/redis"
	"github.com/ribeirosaimon/tooltip/tserver"
)

func NewRedisConnection() redis.RConnInterface {
	config := tserver.GetRedisConfig()
	return redis.NewRedisConnection(
		redis.WithUrl(config.Host),
	)
}
