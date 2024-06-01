package validator

//go:generate mockgen -package=mockValidator -destination=./mockValidator/mock.go -source=validator.go
type Validator interface {
	ValidateStruct(i interface{}) error
	Engine() interface{}
}
