package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	U "github.com/wiliehidayat87/mylib/v2"
	"github.com/wiliehidayat87/testMQ/src/app"
	"github.com/wiliehidayat87/testMQ/src/config"
	psql "github.com/wiliehidayat87/testMQ/src/datasource/psql"
	"github.com/wiliehidayat87/testMQ/src/datasource/rabbitmq"
	"github.com/wiliehidayat87/testMQ/src/datasource/redis"
)

var serverCmd = &cobra.Command{
	Use:   "serverCmd",
	Short: "Server CMD CLI",
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
		l.SetUpLog(U.Utils{LogThread: l.GetUniqId(), LogName: "access_log"})

		l.Write(l.LogName, "info",
			fmt.Sprintf("RMQ_HOST: %#v", config.RMQ_USER),
		)

		// Postgre SQL
		pgSql, err := psql.InitPSQL(config.URI_POSTGRES)

		if err != nil {

			l.Write(l.LogName, "error",
				fmt.Sprintf("Error db init occured: %#v", err),
			)
		}

		// Redis
		red := redis.InitRedis(redis.CfgRed{Host: config.REDIS_HOST, Port: config.REDIS_PORT, Password: config.REDIS_PASS})

		// RabbitMQ
		queue := rabbitmq.InitQueue(rabbitmq.CfgAMQP{Host: config.RMQ_HOST, User: config.RMQ_USER, Pass: config.RMQ_PASS, Port: config.RMQ_PORT, Vhost: config.RMQ_VHOST})

		// SETUP CHANNEL
		queue.SetUpChannel(config.RMQ_EXCHANGETYPE, true, config.RMQ_MOEXCHANGE, true, config.RMQ_MOQUEUE)

		router := app.StartApplication(l, pgSql, red, queue)
		log.Fatal(router.Listen(":" + config.APP_PORT))

	},
}
