package cmd

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var (
	PATH         string
	APP_HOST     string = "host.docker.internal"
	APP_PORT     string = "9022"
	APP_TZ       string = "Asia/Jakarta"
	APP_URL      string = "http://host.docker.internal"
	URI_POSTGRES string = "postgresql://admin:admin135@host.docker.internal:5433/mydb?sslmode=disable"
	REDIS_HOST   string = "host.docker.internal"
	REDIS_PORT   string = "6380"
	REDIS_PASS   string = "password"
	RMQ_HOST     string = "host.docker.internal"
	RMQ_USER     string = "admin"
	RMQ_PASS     string = "testMQ"
	RMQ_PORT     string = "5673"
	LOG_PATH     string = "/Users/wiliewahyuhidayat/Documents/GO/testMQ/logs"
	LOG_LEVEL    int    = 0

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

func init() {

	PATH, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	loc, _ := time.LoadLocation(APP_TZ)
	time.Local = loc

	/**
	 * WEBSERVER SERVICE
	 */
	rootCmd.AddCommand(serverCmd)

	/**
	 * CONSUME SERVICE
	 */
	rootCmd.AddCommand(consumerCmd)

	/**
	 * PUBLISH SERVICE
	 */
	rootCmd.AddCommand(publisherCmd)

}

func Execute() error {
	return rootCmd.Execute()
}

/*
func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return value
}

func getEnvInt(key string) int {
	value := os.Getenv(key)
	valInt, _ := strconv.Atoi(value)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return valInt
}


// Connect to postgresql
func connectPgsql() (*sql.DB, error) {
	db, err := sql.Open("postgres", URI_POSTGRES)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Connect to redis
func connectRedis() (*redis.Client, error) {
	opts, err := redis.ParseURL(URI_REDIS)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}

// Connect to rabbitmq
func connectRabbitMq() rmqp.AMQP {
	var rb rmqp.AMQP
	port, _ := strconv.Atoi(RMQ_PORT)
	rb.SetAmqpURL(RMQ_HOST, port, RMQ_USER, RMQ_PASS)
	rb.SetUpConnectionAmqp()
	return rb
}
*/
