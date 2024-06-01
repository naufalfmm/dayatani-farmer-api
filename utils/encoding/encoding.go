package encoding

//go:generate mockgen -package=mockEncoding -destination=./mockEncoding/mock.go -source=encoding.go
type Encoding interface {
	Encode(raw string) (string, error)
	Decode(encoded string) (string, error)
	Compare(encoded, raw string) error
}
