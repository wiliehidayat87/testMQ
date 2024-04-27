package redis

import (
	"reflect"
	"testing"

	"github.com/go-redis/redis"
	U "github.com/wiliehidayat87/mylib/v2"
)

func TestInitRedis(t *testing.T) {
	type args struct {
		cfg CfgRed
	}
	tests := []struct {
		name string
		args args
		want *Red
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRedis(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed_Get(t *testing.T) {
	type fields struct {
		Redis *redis.Client
	}
	type args struct {
		Log *U.Utils
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Red{
				Redis: tt.fields.Redis,
			}
			if got := r.Get(tt.args.Log, tt.args.key); got != tt.want {
				t.Errorf("Red.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed_Scan(t *testing.T) {
	type fields struct {
		Redis *redis.Client
	}
	type args struct {
		key string
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Red{
				Redis: tt.fields.Redis,
			}
			if got := r.Scan(tt.args.key, tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Red.Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed_GetValue(t *testing.T) {
	type fields struct {
		Redis *redis.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Red{
				Redis: tt.fields.Redis,
			}
			if got := r.GetValue(tt.args.key); got != tt.want {
				t.Errorf("Red.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed_Put(t *testing.T) {
	type fields struct {
		Redis *redis.Client
	}
	type args struct {
		Log *U.Utils
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Red{
				Redis: tt.fields.Redis,
			}
			if got := r.Put(tt.args.Log, tt.args.key, tt.args.val); got != tt.want {
				t.Errorf("Red.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed_Set(t *testing.T) {
	type fields struct {
		Redis *redis.Client
	}
	type args struct {
		Log *U.Utils
		key string
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Red{
				Redis: tt.fields.Redis,
			}
			if got := r.Set(tt.args.Log, tt.args.key, tt.args.val); got != tt.want {
				t.Errorf("Red.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed_Rm(t *testing.T) {
	type fields struct {
		Redis *redis.Client
	}
	type args struct {
		Log *U.Utils
		key string
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
			r := &Red{
				Redis: tt.fields.Redis,
			}
			if err := r.Rm(tt.args.Log, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Red.Rm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
