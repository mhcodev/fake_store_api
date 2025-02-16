package container

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisServer struct {
	Client *redis.Client
}

var (
	RServer *RedisServer
	mu      sync.Mutex
)

func NewRedisServer(client *redis.Client) {
	mu.Lock()
	defer mu.Unlock()
	if RServer == nil {
		RServer = &RedisServer{
			Client: client,
		}
	}
}

func StartRedisServer() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	// Create a context with a timeout for the Ping operation.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to redis server")

	NewRedisServer(client)
}

func GetRedisClient() *RedisServer {
	mu.Lock()
	defer mu.Unlock()

	if RServer == nil || RServer.Client == nil {
		log.Fatal("Redis server is not initialized.")
	}
	return RServer
}

func (rs *RedisServer) Set(ctx context.Context, key string, values map[string]interface{}) error {
	result, err := rs.Client.HSet(ctx, key, values).Result()
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %v", err)
	}

	fmt.Println("Set result:", result)
	return nil
}

func (rs *RedisServer) GetOne(ctx context.Context, key string, field string) (string, error) {
	// Check the type of the key
	dataType, err := rs.Client.Type(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to check key type: %v", err)
	}

	if dataType != "hash" {
		return "", fmt.Errorf("unexpected key type: %s, expected hash", dataType)
	}

	result, err := rs.Client.HGet(ctx, key, field).Result()

	if err != nil {
		return "", fmt.Errorf("failed to deserialize JSON: %v", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(result), &data)

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to get value: %v", err)
	}

	value, ok := data[key].(string)
	if !ok {
		return "", fmt.Errorf("value for key %s is not a string", key)
	}

	return value, nil
}

func (rs *RedisServer) GetAll(ctx context.Context, key string) (map[string]string, error) {
	result, err := rs.Client.HGetAll(ctx, key).Result()

	if err == redis.Nil || len(result) == 0 {
		return nil, nil // Return nil if key doesn't exist
	} else if err != nil {
		return nil, fmt.Errorf("failed to get value: %v", err)
	}

	fmt.Println("GetAll", result)

	return result, nil
}
