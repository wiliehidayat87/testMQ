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
		Log := U.InitLog(U.Utils{
			LogPath:             LOG_PATH,
			LogLevelInit:        LOG_LEVEL,
			AccessLogFormat:     "${ip} ${time} ${method} ${url} ${body} ${referer} ${ua} ${header} ${status} ${latency}\n",
			AccessLogTimeFormat: "[02/Jan/2006:15:04:05 Z0700]",
			TimeZone:            APP_TZ,
		})
		Log.SetUpLog(U.Utils{LogThread: Log.GetUniqId(), LogName: "access_log"})

		Log.Write("info",
			fmt.Sprintf("RMQ_HOST: %#v", RMQ_USER),
		)

		// Postgre SQL
		pgSql, err := psql.InitPSQL(URI_POSTGRES)

		if err != nil {

			Log.Write("error",
				fmt.Sprintf("Error db init occured: %#v", err),
			)
		}

		// Redis
		red := redis.InitRedis(redis.CfgRed{Host: REDIS_HOST, Port: REDIS_PORT, Password: REDIS_PASS})

		// RabbitMQ
		queue := rabbitmq.InitQueue(rabbitmq.CfgAMQP{Host: RMQ_HOST, User: RMQ_USER, Pass: RMQ_PASS, Port: RMQ_PORT})

		// SETUP CHANNEL
		queue.SetUpChannel(config.RMQ_EXCHANGETYPE, true, config.RMQ_MOEXCHANGE, true, config.RMQ_MOQUEUE)

		router := app.StartApplication(Log, pgSql, red, queue)
		log.Fatal(router.Listen(":" + APP_PORT))

	},
}
