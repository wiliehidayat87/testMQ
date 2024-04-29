package handler

import (
	"database/sql"
	"fmt"
	"sync"

	U "github.com/wiliehidayat87/mylib/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/wiliehidayat87/rmqp"
	"github.com/wiliehidayat87/testMQ/src/config"
	"github.com/wiliehidayat87/testMQ/src/datasource/redis"
)

type IncomingHandler struct {
	L   *U.Utils
	DB  *sql.DB
	R   *redis.Red
	Msg rmqp.AMQP
}

func NewIncomingHandler(obj IncomingHandler) *IncomingHandler {
	return &IncomingHandler{
		L:   obj.L,
		DB:  obj.DB,
		R:   obj.R,
		Msg: obj.Msg,
	}
}

func (h *IncomingHandler) PublishMessage(c *fiber.Ctx) error {

	h.L.SetUpLog(U.Utils{LogThread: h.L.GetUniqId(), LogName: "publisher"})

	corId := U.Concat("MOR", h.L.GetUniqId())
	message := c.Params("message")

	published := h.Msg.PublishMsg(rmqp.PublishItems{
		ExchangeName: config.RMQ_MOEXCHANGE,
		QueueName:    config.RMQ_MOQUEUE,
		ContentType:  "text/plain",
		CorId:        corId,
		Payload:      message,
		Priority:     0,
	})

	if !published {

		h.L.Write(h.L.LogName, "debug",
			fmt.Sprintf("[x] Failed published: %s, Data: %s ...", corId, message),
		)

	} else {

		h.L.Write(h.L.LogName, "debug",
			fmt.Sprintf("[v] Published: %s, Data: %s ...", corId, message),
		)
	}

	return c.Status(fiber.StatusOK).SendString("OK")
}

func (h *IncomingHandler) ConsumeMessage(c *fiber.Ctx) error {

	h.L.SetUpLog(U.Utils{LogThread: h.L.GetUniqId(), LogName: "consumer"})

	var m sync.Mutex

	messagesData := h.Msg.Subscribe(1, false, config.RMQ_MOQUEUE, config.RMQ_MOEXCHANGE, config.RMQ_MOQUEUE)

	// Loop forever listening incoming data
	forever := make(chan bool)

	// Set into goroutine this listener
	go func() {

		// Loop every incoming data
		for d := range messagesData {

			m.Lock()
			h.L.Write(h.L.LogName, "info",
				fmt.Sprintf("Consume message, correlation id : %s, Data: %#v ...", d.CorrelationId, string(d.Body)),
			)

			fmt.Printf("Consume message, correlation id : %s, Data: %#v ...\n", d.CorrelationId, string(d.Body))
			m.Unlock()

			// Manual consume queue
			d.Ack(false)

		}

	}()

	h.L.Write(h.L.LogName, "info", "[*] Waiting for data...")

	<-forever

	return c.Status(fiber.StatusOK).SendString("OK")
}
