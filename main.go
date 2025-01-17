package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a Redis client
	redisURL := os.Getenv("REDIS_URL") // Use an environment variable for the Redis URL
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	client := redis.NewClient(options)

	// Ping Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	fmt.Printf("Redis PING response: %s\n", pong)

	// Close the client
	err = client.Close()
	if err != nil {
		log.Printf("Error while closing Redis connection: %v", err)
	}
}
