package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Message struct {
	Time    int    `json:time`
	Content string `json:content`
}

func bodyFrom(args []string) (int, string) {
	var s string
	var i int
	fmt.Println(args)
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
		i = 1
	} else {
		s = args[1]
		i, _ = strconv.Atoi(args[2])
	}
	return i, s
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"qweqwe", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	time, content := bodyFrom(os.Args)
	message := Message{
		Time:    time,
		Content: content,
	}
	body, err := json.Marshal(message)
	failOnError(err, "Failed to declare a queue")

	// body := "Hello rabbitmq!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	var um Message
	json.Unmarshal(body, &um)
	log.Printf(" [x] Sent %v", um)
	failOnError(err, "Failed to publish a message")
}
