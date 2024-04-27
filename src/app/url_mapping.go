package app

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	U "github.com/wiliehidayat87/mylib/v2"
	"github.com/wiliehidayat87/rmqp"

	"github.com/wiliehidayat87/testMQ/src/datasource/redis"
	"github.com/wiliehidayat87/testMQ/src/handler"
)

func mapUrls(log *U.Utils, db *sql.DB, r *redis.Red, msg rmqp.AMQP) *fiber.App {

	f := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	f.Use(logger.New(logger.Config{
		Format:     log.AccessLogFormat,
		TimeFormat: log.AccessLogTimeFormat,
		TimeZone:   log.TimeZone,
		Output:     log.LogOS,
	}))

	h := handler.NewIncomingHandler(handler.IncomingHandler{
		Log: log, DB: db, R: r, Msg: msg,
	})

	// API request/register token init
	f.Get("/api/publish/:message", h.PublishMessage).Name("Token create generator")

	// API request/register token init
	f.Get("/api/consume", h.ConsumeMessage).Name("Token update generator")

	return f
}
