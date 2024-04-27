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
		Log := U.InitLog(U.Utils{
			LogPath:             LOG_PATH,
			LogLevelInit:        LOG_LEVEL,
			AccessLogFormat:     "${ip} ${time} ${method} ${url} ${body} ${referer} ${ua} ${header} ${status} ${latency}\n",
			AccessLogTimeFormat: "[02/Jan/2006:15:04:05 Z0700]",
			TimeZone:            APP_TZ,
		})
		Log.SetUpLog(U.Utils{LogThread: Log.GetUniqId(), LogName: "consumer"})

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
		queue := rabbitmq.InitQueue(rabbitmq.CfgAMQP{Host: RMQ_HOST, User: RMQ_USER, Pass: RMQ_PASS, Port: RMQ_PORT})

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

				Log.Write("info",
					fmt.Sprintf("Consume message, correlation id : %s, Data: %s ...", d.CorrelationId, string(d.Body)),
				)

				m.Unlock()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		Log.Write("info", "[*] Waiting for data...")

		<-forever

	},
}
