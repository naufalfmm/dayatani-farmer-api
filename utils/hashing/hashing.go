package hashing

type Hashing interface {
	Generate(raw string) (string, error)
	Check(hashed, raw string) error
}
