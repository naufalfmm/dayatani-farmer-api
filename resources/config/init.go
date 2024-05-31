package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvConfig struct {
	Name string `envconfig:"NAME" default:"Dayatani Farmer API" required:"true"`
	Port int    `envconfig:"PORT" default:"8080" required:"true"`

	DbHost     string `envconfig:"DB_HOST" required:"true"`
	DbPort     string `envconfig:"DB_PORT" required:"true"`
	DbUsername string `envconfig:"DB_USERNAME" required:"true"`
	DbPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DbName     string `envconfig:"DB_NAME" required:"true"`

	DbSslMode                   string        `envconfig:"DB_SSL_MODE" default:"prefer"`
	DbStatementTimeout          time.Duration `envconfig:"DB_STATEMENT_TIMEOUT" default:"0ms"`
	DbTransactionSessionTimeout time.Duration `envconfig:"DB_TRANSACTION_SESSION_TIMEOUT" default:"0ms"`
	DbSessionTimeout            time.Duration `envconfig:"DB_SESSION_TIMEOUT" default:"0ms"`

	DbLogMode          bool          `envconfig:"DB_LOG_MODE" default:"false"`
	DbLogSlowThreshold time.Duration `envconfig:"DB_LOG_SLOW_THRESHOLD"`

	DbRetry     int           `envconfig:"DB_RETRY" default:"3"`
	DbWaitSleep time.Duration `envconfig:"DB_WAIT_SLEEP" default:"1s"`

	LogMode bool `envconfig:"LOG_MODE" default:"false"`

	HashedCost         int    `envconfig:"HASHED_COST" default:"5" required:"true"`
	HashedAuthUsername string `envconfig:"HASHED_AUTH_USERNAME" required:"true"`
	HashedAuthPassword string `envconfig:"HASHED_AUTH_PASSWORD" required:"true"`

	Base64EncodingType string `envconfig:"BASE64_ENCODING_TYPE" default:"std" required:"true"`
}

func NewConfig() (*EnvConfig, error) {
	var config EnvConfig

	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &config); err != nil {
			return nil, errors.Wrap(err, "failed to read from env variable")
		}

		return &config, nil
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil
}
