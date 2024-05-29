package middlewares

type (
	Middlewares interface{}

	middlewares struct{}
)

func Init() (Middlewares, error) {
	return &middlewares{}, nil
}
