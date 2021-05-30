package rabbitmq

import (
	"log"
	"fmt"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/db"
	"github.com/seb7887/go-microservices/models"
)

func AMQPListen() error {
	conn, err := amqp.Dial(config.GetConfig().AMQPUrl)
	if err != nil {
		return fmt.Errorf("Error to connecting to the message broker")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel")
	}
	defer ch.Close()

	topicName := "new_shipping"
	queueName := "shippings"
	err = ch.ExchangeDeclare(topicName, "topic", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Failed creating the exchange")
	}

	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Error declaring the queue")
	}

	err = ch.QueueBind(queueName, "#", topicName, false, nil)
	if err != nil {
		return fmt.Errorf("Error binding the queue")
	}

	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Failer to register as a consumer")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Println("Received shipping order")
			var shipping models.Shipping
			json.Unmarshal(d.Body, &shipping)
			db.DB.Create(&shipping)
			d.Ack(false)
		}
	}()

	fmt.Println("Shipping service listening for events...")

	<-forever
	return nil
}