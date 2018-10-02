package main

import (
	"context"
	"kafka-go"
	"strconv"
	"time"
)

func main() {
	topic := "new_topic"
	partition := 0

	producer, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	number := 1
	for number < 20 {
		producer.SetWriteDeadline(time.Now().Add(10 * time.Second))
		strNumber := strconv.FormatInt(int64(number), 10)
		producer.WriteMessages(
			kafka.Message{Value: []byte("Value " + strNumber), Key: []byte("K" + strNumber)},
		)
		number++
	}
	producer.Close()
}
