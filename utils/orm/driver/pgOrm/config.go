package pgOrm

import (
	"fmt"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger/zeroLogger"
	"github.com/uptrace/bun/driver/pgdriver"
)

type postgresConfig struct {
	host     string
	port     string
	username string
	password string
	dbname   string

	sslMode                   bool
	statementTimeout          time.Duration
	transactionSessionTimeout time.Duration
	sessionTimeout            time.Duration
	connParams                map[string]interface{}

	logger           logger.Logger
	logMode          bool
	logSlowThreshold time.Duration

	retry     int
	waitSleep time.Duration
}

func generateDefault() (postgresConfig, error) {
	log, err := zeroLogger.NewZeroLogger()
	if err != nil {
		return postgresConfig{}, err
	}

	return postgresConfig{
		host:     "localhost",
		port:     "5432",
		username: "root",

		statementTimeout: 0,

		logger:           log,
		logMode:          false,
		logSlowThreshold: 200 * time.Millisecond,
	}, nil
}

func (c postgresConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}

func (c postgresConfig) toDriverOptions() []pgdriver.Option {
	opts := []pgdriver.Option{
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithInsecure(true),
	}

	if c.host != "" && c.port != "" {
		opts = append(opts, pgdriver.WithAddr(c.Addr()))
	}

	if c.username != "" {
		opts = append(opts, pgdriver.WithUser(c.username))
	}

	if c.password != "" {
		opts = append(opts, pgdriver.WithPassword(c.password))
	}

	if c.dbname != "" {
		opts = append(opts, pgdriver.WithDatabase(c.dbname))
	}

	if len(c.connParams) > 0 {
		opts = append(opts, pgdriver.WithConnParams(c.connParams))
	}

	return opts
}

type PostgresConfig func(c *postgresConfig)

func WithHostPort(host, port string) PostgresConfig {
	return func(c *postgresConfig) {
		c.host = host
		c.port = port
	}
}

func WithUsernamePassword(username, password string) PostgresConfig {
	return func(c *postgresConfig) {
		c.username = username
		c.password = password
	}
}

func WithDatabaseName(dbname string) PostgresConfig {
	return func(c *postgresConfig) {
		c.dbname = dbname
	}
}

func WithSSLMode(sslMode bool) PostgresConfig {
	return func(c *postgresConfig) {
		c.sslMode = sslMode

		if c.connParams == nil {
			c.connParams = map[string]interface{}{}
		}

		c.connParams["sslmode"] = c.sslMode
	}
}

func WithTimeout(timeout time.Duration) PostgresConfig {
	return func(c *postgresConfig) {
		c.statementTimeout = timeout
		c.transactionSessionTimeout = timeout
		c.sessionTimeout = timeout

		if c.connParams == nil {
			c.connParams = map[string]interface{}{}
		}

		c.connParams["statement_timeout"] = timeout.String()
		c.connParams["idle_in_transaction_session_timeout"] = timeout.String()
		c.connParams["idle_session_timeout"] = timeout.String()
	}
}

func WithStatementTimeout(timeout time.Duration) PostgresConfig {
	return func(c *postgresConfig) {
		c.sessionTimeout = timeout

		if c.connParams == nil {
			c.connParams = map[string]interface{}{}
		}
		c.connParams["statement_timeout"] = timeout.String()
	}
}

func WithTransactionSessionTimeout(timeout time.Duration) PostgresConfig {
	return func(c *postgresConfig) {
		c.transactionSessionTimeout = timeout

		if c.connParams == nil {
			c.connParams = map[string]interface{}{}
		}
		c.connParams["idle_in_transaction_session_timeout"] = timeout.String()
	}
}

func WithSessionTimeout(timeout time.Duration) PostgresConfig {
	return func(c *postgresConfig) {
		c.sessionTimeout = timeout

		if c.connParams == nil {
			c.connParams = map[string]interface{}{}
		}
		c.connParams["idle_session_timeout"] = timeout.String()
	}
}

func WithLogger(logger logger.Logger) PostgresConfig {
	return func(c *postgresConfig) {
		c.logger = logger
		c.logMode = true
	}
}

func WithLog(logger logger.Logger, slowThreshold time.Duration) PostgresConfig {
	return func(c *postgresConfig) {
		c.logger = logger
		c.logSlowThreshold = slowThreshold
		c.logMode = true
	}
}

func WithRetry(retry int, waitSleep time.Duration) PostgresConfig {
	return func(c *postgresConfig) {
		c.retry = retry
		c.waitSleep = waitSleep
	}
}
