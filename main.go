package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/wiliehidayat87/rmqp"
)

var rabbit rmqp.AMQP

func main() {

	if len(os.Args) < 2 {

		fmt.Println("No args")

	} else {

		rabbit.SetAmqpURL(
			"localhost",
			5675,
			"admin",
			"admin",
		)

		rabbit.SetUpConnectionAmqp()

		_func := os.Args[1]

		if _func == "publish" {

			msg := os.Args[2]

			publishData(msg)

		} else if _func == "consume" {

			consumeData()
		}
	}

}

func publishData(msg string) {

	fmt.Println("msgUrl = " + rabbit.MsgBrokerURL)

	rabbit.SetUpChannel("direct", true, "eTest", true, "qTest")

	rabbit.IntegratePublish("eTest", "qTest", "plain/text", "123", msg)
}

func consumeData() {

	rabbit.SetUpChannel("direct", true, "eTest", true, "qTest")

	messagesData := rabbit.Subscribe(1, false, "qTest", "eTest", "qTest")

	//messagesData := rmqp.
	// Initial sync waiting group
	var wg sync.WaitGroup

	// Loop forever listening incoming data
	forever := make(chan bool)

	// Set into goroutine this listener
	go func() {

		// Loop every incoming data
		for d := range messagesData {

			wg.Add(1)

			mainProcessor(&wg, d.Body)

			wg.Wait()

			// Manual consume queue
			d.Ack(false)

		}

	}()

	fmt.Println("[*] Waiting for data...")

	<-forever

}

func mainProcessor(wg *sync.WaitGroup, message []byte) {

	fmt.Printf("Data = %#v\n", string(message))

	wg.Done()

}
