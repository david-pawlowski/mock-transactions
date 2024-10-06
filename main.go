package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"time"
)

type Transaction struct {
	date     time.Time `json:"date"`
	title    string    `json:"title"`
	receiver string    `json:"receiver"`
	value    uint16    `json:"value"`
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

func RandomTransaction() Transaction {
	return Transaction{
		date:     time.Now(),
		title:    RandString(5),
		receiver: RandString(10),
		value:    100,
	}
}

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	for {
		err := rdb.Publish(ctx, "transactions-channel", "Hello World!").Err()
		if err != nil {
			log.Fatalf("Something went wrong, %v", err)
		}
		time.Sleep(5 * time.Second)
	}
}
