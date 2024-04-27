package cmd

import (
	"github.com/spf13/cobra"
	U "github.com/wiliehidayat87/mylib/v2"

	"github.com/wiliehidayat87/testMQ/src/config"
)

var publisherCmd = &cobra.Command{
	Use:   "publisherCmd",
	Short: "Publisher CMD CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		Log := U.InitLog(U.Utils{
			LogPath:             config.LOG_PATH,
			LogLevelInit:        config.LOG_LEVEL,
			AccessLogFormat:     "${ip} ${time} ${method} ${url} ${body} ${referer} ${ua} ${header} ${status} ${latency}\n",
			AccessLogTimeFormat: "[02/Jan/2006:15:04:05 Z0700]",
			TimeZone:            config.APP_TZ,
		})
		Log.SetUpLog(U.Utils{LogThread: Log.GetUniqId(), LogName: "publisher"})

		Log.Write("info", "This is publisher arg")
	},
}
