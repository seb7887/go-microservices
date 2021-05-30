package rabbitmq

import (
	"log"
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/models"
)

type OrderMessage struct {
	OrderId string
	ProductName string
	TotalAmount int32
}

func PublishMessage(order *models.Order) error {
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

	payload := OrderMessage{
		OrderId: strconv.FormatUint(uint64(order.ID), 10),
		ProductName: order.ProductName,
		TotalAmount: order.TotalAmount,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Error formatting message")
	}
	message := amqp.Publishing{
		Body: body,
	}

	err = ch.Publish(topicName, "random-key", false, false, message)
	if err != nil {
		return fmt.Errorf("Failed publishing a message to the queue")
	}

	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Error declaring the queue")
	}

	err = ch.QueueBind(queueName, "#", topicName, false, nil)
	if err != nil {
		return fmt.Errorf("Error binding the queue")
	}

	log.Println("-> A new shipping has been published")

	return nil
}