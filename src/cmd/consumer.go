package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	U "github.com/wiliehidayat87/mylib/v2"
	"github.com/wiliehidayat87/testMQ/src/config"
	"github.com/wiliehidayat87/testMQ/src/datasource/rabbitmq"
)

var consumerCmd = &cobra.Command{
	Use:   "consumerCmd",
	Short: "Publisher CMD CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// Setup / Init the log
		l := U.InitLog(U.Utils{
			LogPath:             config.LOG_PATH,
			LogLevelInit:        config.LOG_LEVEL,
			AccessLogFormat:     "${ip} ${time} ${method} ${url} ${body} ${referer} ${ua} ${header} ${status} ${latency}\n",
			AccessLogTimeFormat: "[02/Jan/2006:15:04:05 Z0700]",
			TimeZone:            config.APP_TZ,
		})
		l.SetUpLog(U.Utils{LogThread: l.GetUniqId(), LogName: "consumer"})

		// Postgre SQL
		/*
			pgSql, err := psql.InitPSQL(URI_POSTGRES)

			if err != nil {

				Log.Write("error",
					fmt.Sprintf("Error db init occured: %#v", err),
				)
			}
		*/

		// Redis
		//red := redis.InitRedis(redis.CfgRed{Host: REDIS_HOST, Port: REDIS_PORT, Password: REDIS_PASS})

		// Initial sync waiting group
		var m sync.Mutex

		// RabbitMQ
		queue := rabbitmq.InitQueue(rabbitmq.CfgAMQP{Host: config.RMQ_HOST, User: config.RMQ_USER, Pass: config.RMQ_PASS, Port: config.RMQ_PORT})

		// SETUP CHANNEL
		queue.SetUpChannel(config.RMQ_EXCHANGETYPE, true, config.RMQ_MOEXCHANGE, true, config.RMQ_MOQUEUE)

		messagesData := queue.Subscribe(1, false, config.RMQ_MOQUEUE, config.RMQ_MOEXCHANGE, config.RMQ_MOQUEUE)

		// Loop forever listening incoming data
		forever := make(chan bool)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				m.Lock()

				l.Write(l.LogName, "info",
					fmt.Sprintf("Consume message, correlation id : %s, Data: %s ...", d.CorrelationId, string(d.Body)),
				)

				m.Unlock()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		l.Write(l.LogName, "info", "[*] Waiting for data...")

		<-forever

	},
}
