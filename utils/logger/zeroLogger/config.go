package zeroLogger

type (
	config struct {
		enabled bool
	}

	LoggerConfig func(c *config)
)

func WithEnabled(enabled bool) LoggerConfig {
	return func(c *config) {
		c.enabled = enabled
	}
}
