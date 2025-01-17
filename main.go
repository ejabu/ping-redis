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
	// Read Redis URL from environment variable
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL environment variable is not set")
	}

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	client := redis.NewClient(options)

	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}()

	// Infinite loop to ping Redis every 5 seconds
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		pong, err := client.Ping(ctx).Result()
		if err != nil {
			log.Printf("Failed to ping Redis: %v", err)
		} else {
			fmt.Printf("Redis PING Response: %s\n", pong)
		}

		// Wait for 5 seconds before the next ping
		time.Sleep(2 * time.Second)
	}
}
