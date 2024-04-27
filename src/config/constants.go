package config

const (
	APP_HOST         string = "host.docker.internal"
	APP_PORT         string = "9022"
	APP_TZ           string = "Asia/Jakarta"
	APP_URL          string = "http://host.docker.internal"
	URI_POSTGRES     string = "postgresql://admin:corePu5hkomx@host.docker.internal:5433/vasDB?sslmode=disable"
	REDIS_HOST       string = "host.docker.internal"
	REDIS_PORT       string = "6380"
	REDIS_PASS       string = "password"
	RMQ_HOST         string = "host.docker.internal"
	RMQ_USER         string = "admin"
	RMQ_PASS         string = "corePu5hkomx"
	RMQ_PORT         string = "5673"
	RMQ_EXCHANGETYPE string = "direct"
	RMQ_DATATYPE     string = "application/json"
	RMQ_MOEXCHANGE   string = "E_MO"
	RMQ_MOQUEUE      string = "Q_MO"
	LOG_PATH         string = "/Users/wiliewahyuhidayat/Documents/GO/testMQ/logs"
	LOG_LEVEL        int    = 0
)
