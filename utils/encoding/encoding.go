package encoding

type Encoding interface {
	Encode(raw string) (string, error)
	Decode(encoded string) (string, error)
	Compare(encoded, raw string) error
}
