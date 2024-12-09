package main

import (
	"log"

	"github.com/streadway/amqp"
)

func consumeQueue() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("image_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			// Process the image (e.g., download, compress, upload to S3)
			log.Printf("Processing image: %s", msg.Body)
		}
	}()

	log.Println("Waiting for messages...")
	<-forever
}

func main() {
	consumeQueue()
}
