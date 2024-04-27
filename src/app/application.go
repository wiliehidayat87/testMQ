package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	U "github.com/wiliehidayat87/mylib/v2"
	"github.com/wiliehidayat87/rmqp"
	r "github.com/wiliehidayat87/testMQ/src/datasource/redis"

	_ "github.com/lib/pq"
)

func StartApplication(Log *U.Utils, db *sql.DB, r *r.Red, msg rmqp.AMQP) *fiber.App {
	return mapUrls(Log, db, r, msg)
}
