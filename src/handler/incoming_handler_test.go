package handler

import (
	"database/sql"
	"testing"

	"github.com/gofiber/fiber/v2"
	U "github.com/wiliehidayat87/mylib/v2"
	"github.com/wiliehidayat87/rmqp"
	"github.com/wiliehidayat87/testMQ/src/datasource/redis"
)

func TestIncomingHandler_PublishMessage(t *testing.T) {
	type fields struct {
		L   *U.Utils
		DB  *sql.DB
		R   *redis.Red
		Msg rmqp.AMQP
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &IncomingHandler{
				L:   tt.fields.L,
				DB:  tt.fields.DB,
				R:   tt.fields.R,
				Msg: tt.fields.Msg,
			}
			if err := h.PublishMessage(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("IncomingHandler.PublishMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIncomingHandler_ConsumeMessage(t *testing.T) {
	type fields struct {
		L   *U.Utils
		DB  *sql.DB
		R   *redis.Red
		Msg rmqp.AMQP
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &IncomingHandler{
				L:   tt.fields.L,
				DB:  tt.fields.DB,
				R:   tt.fields.R,
				Msg: tt.fields.Msg,
			}
			if err := h.ConsumeMessage(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("IncomingHandler.ConsumeMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
