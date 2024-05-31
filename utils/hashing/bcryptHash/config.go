package bcryptHash

type (
	config struct {
		cost int
	}

	BcryptConfig func(c *config)
)

func WithCost(cost int) BcryptConfig {
	return func(c *config) {
		c.cost = cost
	}
}
