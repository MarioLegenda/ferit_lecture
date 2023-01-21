package memoryStore

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

type MemoryStore struct {
	redis *redis.Client
}

func NewMemoryStore(client *redis.Client) MemoryStore {
	return MemoryStore{redis: client}
}

func (m MemoryStore) Set(key string, value interface{}) error {
	if err := m.redis.Set(key, value, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (m MemoryStore) Get(key string, model interface{}) error {
	val, err := m.redis.Get(key).Result()

	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(val), model); err != nil {
		return err
	}

	return nil
}

func (m MemoryStore) Delete(key string) error {
	if err := m.redis.Del(key).Err(); err != nil {
		return err
	}

	return nil
}
