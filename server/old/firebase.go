package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func maina() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	topic := "testTopic"

	message := &messaging.Message{
		Data: map[string]string{
			"san":  "120",
			"time": "2:30",
		},
		Topic: topic,
	}

	res, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully sent message:", res)
}
