package bcryptHash

import (
	"github.com/naufalfmm/dayatani-farmer-api/utils/hashing"
	"golang.org/x/crypto/bcrypt"
)

type bcrHash struct {
	config config
}

func NewBcrypt(confs ...BcryptConfig) (hashing.Hashing, error) {
	config := config{}

	for _, conf := range confs {
		conf(&config)
	}

	return &bcrHash{
		config: config,
	}, nil
}

func (b *bcrHash) Generate(raw string) (string, error) {
	bytesGenPass, err := bcrypt.GenerateFromPassword([]byte(raw), b.config.cost)
	if err != nil {
		return "", err
	}

	return string(bytesGenPass), nil
}

func (b *bcrHash) Check(hashed, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
}
