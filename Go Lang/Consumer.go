package main

import (
	"context"
	"fmt"
	"kafka-go"
	"strconv"
)

func main() {
	topic := "new_topic"
	partition := 0

	consumer, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	for {
		msg, err := consumer.ReadMessage(1e6)
		fmt.Println("Key : " + string(msg.Key) + "	Value : " + string(msg.Value) +
			"	Partition : " + strconv.FormatInt(int64(msg.Partition), 10) +
			"	Offset : " + strconv.FormatInt(msg.Offset, 10) +
			"	Time : " + string(msg.Time.Format("2006-01-02 15:04:05")))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	consumer.Close()
}
