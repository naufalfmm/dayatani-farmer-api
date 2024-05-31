package base64Encoding

type (
	config struct {
		encType string
	}

	Base64EncodingConfig func(c *config)
)

func WithEncType(encType string) Base64EncodingConfig {
	return func(c *config) {
		c.encType = encType
	}
}
