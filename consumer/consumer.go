package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart", 0)
	if err != nil {
		log.Fatal(err)
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 8))

	batch := conn.ReadBatch(1e3, 1e9) // 1e3 = 1000kb, 1e9 = 1000000kb
	bytes := make([]byte, 1e3)

	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		log.Print(string(bytes))
	}
}
