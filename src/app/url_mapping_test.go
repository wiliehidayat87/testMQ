package app

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	U "github.com/wiliehidayat87/mylib/v2"
	"github.com/wiliehidayat87/rmqp"
	"github.com/wiliehidayat87/testMQ/src/datasource/redis"
)

func Test_mapUrls(t *testing.T) {
	type args struct {
		log *U.Utils
		db  *sql.DB
		r   *redis.Red
		msg rmqp.AMQP
	}
	tests := []struct {
		name string
		args args
		want *fiber.App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapUrls(tt.args.log, tt.args.db, tt.args.r, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapUrls() = %v, want %v", got, tt.want)
			}
		})
	}
}
