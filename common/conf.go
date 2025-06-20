package common

import (
	"os"
	"strconv"
)

// Version is the git hash of the application. Do not edit. This is
// automatically set using -ldflags during build time.
var Version string

// Conf is the configuration file data for the ripple API.
type Conf struct {
	DatabaseType           string `description:"At the moment, 'mysql' is the only supported database type."`
	DSN                    string `description:"The Data Source Name for the database. More: https://github.com/go-sql-driver/mysql#dsn-data-source-name"`
	ListenTo               string `description:"The IP/Port combination from which to take connections, e.g. :8080"`
	Unix                   bool   `description:"Bool indicating whether ListenTo is a UNIX socket or an address."`
	SentryDSN              string `description:"thing for sentry whatever"`
	HanayoKey              string
	BeatmapRequestsPerUser int
	RankQueueSize          int
	OsuAPIKey              string
	RedisAddr              string
	RedisPassword          string
	RedisDB                int
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return fallback
}

func envConf() Conf {
	return Conf{
		DatabaseType:           getEnv("DB_SCHEME", "mysql"),
		DSN:                    getEnv("DSN", getEnv("DB_USER", "root")+":"+getEnv("DB_PASS", "osu")+"@tcp("+getEnv("DB_HOST", "db")+":3306)/"+getEnv("DB_NAME", "ripple")),
		ListenTo:               ":80",
		Unix:                   getEnvBool("UNIX", false),
		SentryDSN:              getEnv("SENTRY_DSN", ""),
		HanayoKey:              getEnv("APP_HANAYO_KEY", "APISECRETVALUE"),
		BeatmapRequestsPerUser: getEnvInt("BEATMAP_REQUESTS_PER_USER", 2),
		RankQueueSize:          getEnvInt("RANK_QUEUE_SIZE", 25),
		OsuAPIKey:             getEnv("API_KEYS_POOL", "OSUAPIKEY"),
		RedisAddr:              getEnv("REDIS_HOST", "redis:6379")+":"+getEnv("REDIS_PORT", "6379"),
		RedisPassword:          getEnv("REDIS_PASS", ""),
		RedisDB:                getEnvInt("REDIS_DB", 0),
	}
}

// Load returns the configuration from environment variables.
func Load() (c Conf, halt bool) {
	c = envConf()
	halt = false
	return
}

// GetConf returns the configuration from environment variables.
func GetConf() *Conf {
	conf := envConf()
	return &conf
}
