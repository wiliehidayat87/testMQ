package rmqp

// this package is the initiation of AMQP
// and handle of any create connection of Rabbit MQ Framework

import (
	"fmt"

	"github.com/streadway/amqp"
)

// IntegratePublish : publish a message
func (rabbit *AMQP) IntegratePublish(exch string, queue string, contentType string, correlationId string, requestBody string) {

	err := rabbit.Channel.Publish(
		exch,  // exchange name
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:   contentType,
			DeliveryMode:  amqp.Persistent,
			CorrelationId: correlationId,
			Body:          []byte(requestBody),
		},
	)

	//input := Lib.ReduceWords(requestBody, 0, 30)

	if err != nil {
		fmt.Printf("[x] Failed published: %s, Data: %s ...\n", correlationId, requestBody)
	} else {
		fmt.Printf("[v] Published: %s, Data: %s ...", correlationId, requestBody)
	}

}
