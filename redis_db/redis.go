package redis_db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

// Initialize Redis client
func ConnectRedis(RedisURL string, RedisPass string, RedisDB int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     RedisURL,  // Redis address
		Password: RedisPass, // Password if any
		DB:       RedisDB,   // Default DB
		PoolSize: 100,
	})

	// Check connection
	// if pong, err := rdb.Ping(ctx).Result(); err != nil {
	// 	log.Fatalf("Redis connection error: %v", err)
	// } else {
	// 	fmt.Println("Redis connected :", pong)
	// }
}

// SetWithTTL sets a key with a TTL
func SetWithTTL(key string, value string, ttl time.Duration) error {
	err := rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to set key '%s': %w", key, err)
	}
	return nil
}

func ScanKey(keyPattern string, keyLimit int64) (redisKeys []string) {
	var cursor uint64
	var keys []string
	var err error
	for {
		var scannedKeys []string
		scannedKeys, cursor, err = rdb.Scan(ctx, cursor, keyPattern, keyLimit).Result()
		if err != nil {
			panic(err)
		}
		keys = append(keys, scannedKeys...)

		if cursor == 0 {
			break
		}
	}
	return keys
}

// GetValue retrieves a value by key
func GetValue(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key '%s' does not exist", key)
	} else if err != nil {
		return "", fmt.Errorf("error getting key '%s': %w", key, err)
	}
	return val, nil
}

// GetValue retrieves a value by key
func DelValue(key string) (int64, error) {
	// Delete the key
	deleted, err := rdb.Del(ctx, key).Result()
	if err != nil {
		fmt.Println("Error deleting key:", err)
		return 0, nil
	}

	if deleted > 0 {
		fmt.Printf("Key '%s' deleted successfully.\n", key)
	} else {
		fmt.Printf("Key '%s' does not exist.\n", key)
		return 0, nil
	}
	return deleted, nil
}
