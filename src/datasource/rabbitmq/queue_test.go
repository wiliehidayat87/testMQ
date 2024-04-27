package rabbitmq

import (
	"reflect"
	"testing"

	"github.com/wiliehidayat87/rmqp"
)

func TestInitQueue(t *testing.T) {
	type args struct {
		cfg CfgAMQP
	}
	tests := []struct {
		name string
		args args
		want rmqp.AMQP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitQueue(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
