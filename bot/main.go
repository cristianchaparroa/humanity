package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cristianchaparroa/humanity/bot/services"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	dial := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	fmt.Println(dial)

	conn, err := amqp.Dial(dial)

	if err != nil {
		fmt.Println(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"room", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)

	bot := services.NewBotService()

	go func() {
		for d := range msgs {
			body := d.Body
			log.Printf("--> Received a message: %s", body)

			m := bot.Process(string(body))
			log.Printf("<-- Received a message: %#v", m)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
