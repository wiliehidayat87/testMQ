package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	U "github.com/wiliehidayat87/mylib/v2"
)

type (
	CfgRed struct {
		Host     string
		Port     string
		Password string
	}

	Red struct {
		Redis *redis.Client
	}
)

func InitRedis(cfg CfgRed) *Red {

	// Setup redis
	cRedis := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       0,
	})

	return &Red{
		Redis: cRedis,
	}
}

func (r *Red) Get(l *U.Utils, key string) bool {

	theKeyValue, _ := r.Redis.Get(key).Result()

	var result bool

	if theKeyValue == "" {

		l.Write(l.LogName, "debug",
			fmt.Sprintf("key [%s], is not exist", key),
		)

		result = true

	} else if theKeyValue != "" {

		l.Write(l.LogName, "debug",
			fmt.Sprintf("key [%s] is existed", key),
		)

		result = false
	}

	return result
}

func (r *Red) Scan(key string, obj interface{}) interface{} {

	r.Redis.Get(key).Scan(&obj)

	return obj
}

func (r *Red) GetValue(key string) string {

	theKeyValue, _ := r.Redis.Get(key).Result()

	return theKeyValue
}

func (r *Red) Put(l *U.Utils, key string, val string) bool {

	err := r.Redis.Set(key, val, 0).Err()

	if err != nil {

		l.Write(l.LogName, "debug",
			fmt.Sprintf("Couldn't store this key : %s", key),
		)

		return false
	} else {

		l.Write(l.LogName, "debug",
			fmt.Sprintf("Key stored : %s", key),
		)

		return true
	}

}

func (r *Red) Set(l *U.Utils, key string, val interface{}) bool {

	err := r.Redis.Set(key, val, 0).Err()

	if err != nil {

		l.Write(l.LogName, "debug",
			fmt.Sprintf("Couldn't store this key : %s", key),
		)

		return false
	} else {

		l.Write(l.LogName, "debug",
			fmt.Sprintf("Key stored : %s", key),
		)

		return true
	}

}

func (r *Red) Rm(l *U.Utils, key string) error {

	l.Write(l.LogName, "debug",
		fmt.Sprintf("Key removed : %s", key),
	)

	r.Redis.Del(key)

	return nil
}
