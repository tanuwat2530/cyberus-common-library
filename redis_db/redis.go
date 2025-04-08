package redis_db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

// Initialize Redis client
func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis address
		Password: "",               // Password if any
		DB:       0,                // Default DB
	})

	// Check connection
	if pong, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Redis connection error: %v", err)
	} else {
		fmt.Println("Redis connected:", pong)
	}
}

// SetWithTTL sets a key with a TTL
func SetWithTTL(key string, value string, ttl time.Duration) error {
	err := rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to set key '%s': %w", key, err)
	}
	return nil
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

// func main() {
// 	initRedis()

// 	key := "mykey"
// 	value := "This is a value with TTL"
// 	ttl := 10 * time.Second // expires in 10 seconds

// 	// Set key with TTL
// 	if err := SetWithTTL(key, value, ttl); err != nil {
// 		log.Fatalf("SetWithTTL error: %v", err)
// 	}
// 	fmt.Println("Key set successfully with TTL")

// 	// Get the key
// 	val, err := GetValue(key)
// 	if err != nil {
// 		log.Printf("GetValue error: %v", err)
// 	} else {
// 		fmt.Printf("Retrieved value: %s\n", val)
// 	}
// }
