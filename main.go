package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"time"
)

type Transaction struct {
	Date     time.Time `json:"date"`
	Title    string    `json:"title"`
	Receiver string    `json:"receiver"`
	Value    uint16    `json:"value"`
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

func getRandomTransaction() Transaction {
	return Transaction{
		Date:     time.Now(),
		Title:    randString(5),
		Receiver: randString(10),
		Value:    uint16(rand.Intn(100)),
	}
}

var ctx = context.Background()

func sendTransactions(rdb *redis.Client, interval time.Duration) {
	for {
		transaction := getRandomTransaction()
		data, err := json.Marshal(transaction)
		if err != nil {
			log.Fatalf("Error while decoding json %v", err)
		}

		error := rdb.Publish(ctx, "transactions-channel", data).Err()
		if error != nil {
			log.Fatalf("Something went wrong, %v", err)
		}
		time.Sleep(interval)
	}
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	sendTransactions(rdb, 10*time.Second)
}
