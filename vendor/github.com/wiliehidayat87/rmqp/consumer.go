package rmqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Subscribe : consuming a message
func (rabbit *AMQP) Subscribe(qos int, autoAck bool, routingKey string, exchName string, queueName string) <-chan amqp.Delivery {

	// Set config for basic QOS
	// (Size transaction fetch per 1 minute)
	errQos := rabbit.Channel.Qos(
		qos,   // prefetch count
		0,     // prefetch size
		false, // global
	)

	if errQos != nil {

		fmt.Printf("[x] Failed set QOS Size: %d\n", qos)

		// Closing channel
		defer rabbit.Channel.Close()

		// Closing connection
		defer rabbit.Connection.Close()

		panic(errQos)

	}

	fmt.Printf("[v] Success Set QOS Size: %d\n", qos)

	// Bind concurrent queue
	errQueueBind := rabbit.Channel.QueueBind(
		queueName,  // queue name
		routingKey, // routing key
		exchName,   // exchange
		false,
		nil,
	)

	if errQueueBind != nil {

		fmt.Printf("[x] Failed Binding a Queue [ %s ], with Exchange Name [ %s ]\n", queueName, exchName)

		// Closing channel
		defer rabbit.Channel.Close()

		// Closing connection
		defer rabbit.Connection.Close()

		panic(errQueueBind)

	}

	fmt.Printf("[v] Success Binding a Queue [ %s ], with Exchange Name [ %s ]\n", queueName, exchName)

	messagesData, errConsumeChannel := rabbit.Channel.Consume(
		queueName, // queue
		"",        // consumer
		autoAck,   // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)

	if errConsumeChannel != nil {

		fmt.Printf("[x] Failed to register a consumer : %#v", errConsumeChannel)

		// Closing channel
		defer rabbit.Channel.Close()

		// Closing connection
		defer rabbit.Connection.Close()

		panic(errConsumeChannel)

	}

	fmt.Println("[v] Success to register a consumer")

	return messagesData
}
