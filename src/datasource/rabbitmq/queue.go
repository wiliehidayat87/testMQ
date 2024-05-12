package rabbitmq

import (
	"strconv"

	"github.com/wiliehidayat87/rmqp"
)

type CfgAMQP struct {
	Host  string
	Port  string
	User  string
	Pass  string
	Vhost string
}

func InitQueue(cfg CfgAMQP) rmqp.AMQP {
	var rb rmqp.AMQP

	port, _ := strconv.Atoi(cfg.Port)

	rb.SetAmqpURL(cfg.Host, port, cfg.User, cfg.Pass, cfg.Vhost)

	rb.SetUpConnectionAmqp()
	return rb
}
