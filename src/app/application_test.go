package app

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	U "github.com/wiliehidayat87/mylib/v2"
	"github.com/wiliehidayat87/rmqp"
	r "github.com/wiliehidayat87/testMQ/src/datasource/redis"
)

func TestStartApplication(t *testing.T) {
	type args struct {
		Log *U.Utils
		db  *sql.DB
		r   *r.Red
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
			if got := StartApplication(tt.args.Log, tt.args.db, tt.args.r, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
