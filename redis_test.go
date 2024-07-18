package belajargoredis

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)
	// err := client.Close()
	// assert.Nil(t, err)
}

func TestPing(t *testing.T) {
	res, err := client.Ping(ctx).Result()

	assert.Nil(t, err)
	assert.Equal(t, "PONG", res)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "code", "ABRAM", 3*time.Second)
	res, err := client.Get(ctx, "code").Result()
	assert.Nil(t, err)
	assert.Equal(t, "ABRAM", res)

	time.Sleep(5 * time.Second)
	_, err = client.Get(ctx, "code").Result()
	assert.NotNil(t, err)
}
