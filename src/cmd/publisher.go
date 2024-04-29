package cmd

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/testMQ/src/config"
	U "github.com/wiliehidayat87/testMQ/src/utils"
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

		e := echo.New()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.GET("/", func(c echo.Context) error {
			return rootHandler(Log, c)
		})

		httpPort := config.APP_PORT
		if httpPort == "" {
			httpPort = "8080"
		}

		e.Logger.Fatal(e.Start(":" + httpPort))
	},
}

func rootHandler(Log *U.Utils, c echo.Context) error {

	_echo := fmt.Sprintf("Hello, Docker %s!\n", "DEV")
	Log.Write("info", _echo)
	Log.Write("info", "logname = "+Log.LogName)

	return c.HTML(http.StatusOK, _echo)
}
