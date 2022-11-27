package go_redis_redis_rate

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"testing"
	"time"
)

func TestAlpha(t *testing.T) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "175.178.7.180:9004",
	})

	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow(ctx, "project:123", redis_rate.PerSecond(10))
	if err != nil {
		panic(err)
	}

	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
	time.Sleep(1)
	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
	time.Sleep(1)
	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
}
