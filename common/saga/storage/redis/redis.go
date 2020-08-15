// nolint: gochecknoinits
package redis

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mapserver2007/golang-example-app/common/saga"
	"github.com/mapserver2007/golang-example-app/common/saga/storage"
)

var storageInstance storage.Storage
var redisInit sync.Once

func init() {
	saga.StorageProvider = func(cfg storage.StorageConfig) storage.Storage {
		redisInit.Do(func() {
			storageInstance = newRedisStorage(
				cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password,
			)
		})
		return storageInstance
	}
}

type redisStorage struct {
	pool *redis.Pool
}

func newRedisStorage(host, port, password string) storage.Storage {
	addr := fmt.Sprintf("%s:%s", host, port)
	pool := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialPassword(password))
		},
	}

	return &redisStorage{
		pool: pool,
	}
}

func (s *redisStorage) AppendLog(logId, data string) error {
	conn := s.pool.Get()
	defer conn.Close()
	_, err := conn.Do("RPUSH", logId, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *redisStorage) Lookup(logId string) ([]string, error) {
	conn := s.pool.Get()
	defer conn.Close()
	results, err := redis.Strings(conn.Do("LRANGE", logId, 0, -1))
	return results, err
}

func (s *redisStorage) Close() error {
	s.pool.Close()
	return nil
}

func (s *redisStorage) Cleanup(logId string) error {
	conn := s.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", logId)
	return err
}
