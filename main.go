package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
)

// func createTopic(ctx context.Context) {
// 	conn, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Close()
// }

func produce(ctx context.Context) {
	i := 0
	l := log.New(os.Stdout, "\u001b[34mKafka Writer: \u001B[0m", log.LstdFlags|log.Lshortfile)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Balancer:  &kafka.Hash{},
		BatchSize: 1,
		Logger:    l,
	})
	for {
		m := "this is the message number " + strconv.Itoa(i)
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(m),
		})
		if err != nil {
			panic("could not write message: " + err.Error())
		}
		l.Println("produced: ", m)
		i++
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	l := log.New(os.Stdout, "\u001b[36mKafka Reader: \u001B[0m", log.LstdFlags|log.Lshortfile)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger:  l,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message: " + err.Error())
		}
		l.Println("consumed: ", string(msg.Value))
	}
}

func main() {
	ctx := context.Background()
	// createTopic(ctx)
	go produce(ctx)
	consume(ctx)
}
