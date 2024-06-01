package hashing

//go:generate mockgen -package=mockHashing -destination=./mockHashing/mock.go -source=hashing.go
type Hashing interface {
	Generate(raw string) (string, error)
	Check(hashed, raw string) error
}
