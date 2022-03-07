package rmqp

// this package is the initiation of AMQP
// and handle of any create connection of Rabbit MQ Framework

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Initiate a struct of a connection
type (
	AMQP struct {
		MsgBrokerURL string
		Connection   *amqp.Connection
		ErrConn      error
		Channel      *amqp.Channel
		ErrChannel   error
		Queue        *amqp.Queue
		ErrQueue     error
		ExchangeName string
	}
)

// SetAmqpURL : Build an AMQP URL via config
// param : nil
// return :
// 1. @URL ( Connection string of AMQP for dailing ) -> string
func (rabbit *AMQP) SetAmqpURL(host string, port int, uname string, pwd string) {

	rabbit.MsgBrokerURL = amqp.URI{
		Scheme:   "amqp",
		Host:     host,
		Port:     port,
		Username: uname,
		Password: pwd,
		Vhost:    "/",
	}.String()
}

// SetUpConnectionAmqp : Setup and initiation of dailing of the connection
// this will handle AMQP connection and channel of themself
// param :
// 1. @rabbitURL ( Rabbit URL of the connection string ) -> string
// returns :
// 1. @Conn ( Connection struct compiler consist of Connection & Channel of AMQP ) -> struct interface
// 2. @error ( Related error ) -> error
func (rabbit *AMQP) SetUpConnectionAmqp() {

	// This function should handle the
	// connection init of AMQP
	amqpConn, errConn := amqp.Dial(rabbit.MsgBrokerURL)

	// Checking of the AMQP connection

	if errConn != nil {

		// Write the log if the connection is not connected
		fmt.Printf("[x] Failed Initializing Broker Connection : %#v", errConn)

		// Defer the AMQP related connection then close
		// although the connection is not established
		defer amqpConn.Close()

		// Return to method requestor when error occured
		rabbit.ErrConn = errConn

	} else {

		// Just write the related connection
		// when established
		fmt.Println("[v] Success Initializing Broker Connection")

		// Set into global interface connection of amqp
		rabbit.Connection = amqpConn
	}
}

// SetUpOnceChannel : Setup the once channel of related connection
// Params :
// 1. @c ( Connection struct compiler consist of Connection & Channel of this interface for this AMQP ) -> struct interface
// 2. @logName ( Writing of the related worker ) -> string
// 3. @exchaneName ( Exchange AMQP Name ) -> string
// 4. @queueName ( Queue AMQP Name ) -> string
// Return :
// 1. @Chn ( Struct ) -> struct interface
func SetUpOnceChannel(rabbit *amqp.Connection, exchType string, exchDurable bool, exchName string, queueDurable bool, queueName string) (*amqp.Channel, *amqp.Queue, error) {

	// This is a channel after the connection
	// is established, then initiate the AMQP Channel
	ch, errCh := rabbit.Channel()

	// Checking of the AMQP connection channel

	if errCh != nil {

		// Write the related connection channel when failed
		fmt.Printf("[x] Failed open channel : %#v", errCh)

		// this is an exception only when channel is not established
		// then the connection AMQP should be closed also
		defer ch.Close()

	} else {

		// Just write the related connection
		// when established
		fmt.Println("[v] Success open channel")
	}

	// Declaring of AMQP Exchange method
	errExch := ch.ExchangeDeclare(
		exchName,    // name
		exchType,    // type
		exchDurable, // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)

	// Check the error of declaring an exchange

	if errExch != nil {

		// Write the log if exchange is error
		fmt.Printf("[x] Failed to declare a exchange : %#v", errExch)

	} else {

		// Write the log if the exchange declare is successful
		fmt.Printf("[v] Success declaring an exchange : %s", exchName)
	}

	// Declaring of AMQP queue method

	q, errQueue := ch.QueueDeclare(
		queueName,    // name
		queueDurable, // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)

	// Check if Queue is error

	if errQueue != nil {

		// Write the log if queue is error
		fmt.Printf("[x] Failed to declare a queue : %#v", errQueue)

	} else {

		// Write the log if the queue declare is successful
		fmt.Printf("[v] Success declaring a queue : %s", q.Name)

	}

	// Return the AMQP queue compiler
	// when all is ok
	return ch, &q, nil
}

// SetUpChannel : Setup the once channel of related connection
// Params :
// 1. @c ( Connection struct compiler consist of Connection & Channel of this interface for this AMQP ) -> struct interface
// 2. @logName ( Writing of the related worker ) -> string
// 3. @exchaneName ( Exchange AMQP Name ) -> string
// 4. @queueName ( Queue AMQP Name ) -> string
// Return :
// 1. @Chn ( Struct ) -> struct interface
func (rabbit *AMQP) SetUpChannel(exchType string, exchDurable bool, exchName string, queueDurable bool, queueName string) {

	// This is a channel after the connection
	// is established, then initiate the AMQP Channel
	ch, errCh := rabbit.Connection.Channel()

	// Checking of the AMQP connection channel

	if errCh != nil {

		// Write the related connection channel when failed
		fmt.Printf("[x] Failed open channel : %#v\n", errCh)

		// this is an exception only when channel is not established
		// then the connection AMQP should be closed also
		defer ch.Close()

		// Assign error channel
		rabbit.ErrChannel = errCh

	} else {

		// Just write the related connection
		// when established
		fmt.Println("[v] Success open channel")

		// Assign successful channel
		rabbit.Channel = ch
	}

	// Declaring of AMQP Exchange method
	errExch := rabbit.Channel.ExchangeDeclare(
		exchName,    // name
		exchType,    // type
		exchDurable, // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)

	// Check the error of declaring an exchange

	if errExch != nil {

		// Write the log if exchange is error
		fmt.Printf("[x] Failed to declare a exchange : %#v\n", errExch)

	} else {

		// Write the log if the exchange declare is successful

		rabbit.ExchangeName = exchName

		fmt.Printf("[v] Success declaring an exchange : %s\n", exchName)
	}

	// Declaring of AMQP queue method
	var (
		q        amqp.Queue
		errQueue error
	)

	// Dead Letter Exchange (DLX) handler
	if exchName == "DLX" || queueName == "DLQ" {

		q, errQueue = ch.QueueDeclare(
			queueName,    // name
			queueDurable, // durable
			false,        // delete when unused
			false,        // exclusive
			false,        // no-wait
			amqp.Table{"x-dead-letter-exchange": exchName}, // arguments
		)

	} else {

		q, errQueue = ch.QueueDeclare(
			queueName,    // name
			queueDurable, // durable
			false,        // delete when unused
			false,        // exclusive
			false,        // no-wait
			nil,          // arguments
		)

	}

	// Check if Queue is error

	if errQueue != nil {

		// Write the log if queue is error
		fmt.Printf("[x] Failed to declare a queue worker : %#v\n", errQueue)

		// Closing channel
		defer rabbit.Channel.Close()

		// Closing connection
		defer rabbit.Connection.Close()

		// Assign failed channel
		rabbit.ErrQueue = errQueue

	} else {

		// Write the log if the queue declare is successful
		fmt.Printf("[v] Success declaring a queue worker : %s\n", q.Name)

		rabbit.Queue = &q
	}

}
