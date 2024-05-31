package base64Encoding

import (
	"encoding/base64"

	"github.com/naufalfmm/dayatani-farmer-api/utils/encoding"
)

type base64Enc struct {
	c   config
	enc *base64.Encoding
}

func NewBase64Encoding(cfs ...Base64EncodingConfig) (encoding.Encoding, error) {
	c := config{}
	for _, cf := range cfs {
		cf(&c)
	}

	enc := base64.StdEncoding
	if c.encType == "URL" {
		enc = base64.URLEncoding
	}

	return &base64Enc{
		c:   c,
		enc: enc,
	}, nil
}

func (b *base64Enc) Encode(raw string) (string, error) {
	return b.enc.EncodeToString([]byte(raw)), nil
}

func (b *base64Enc) Decode(encoded string) (string, error) {
	dcd, err := b.enc.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(dcd), nil
}

func (b *base64Enc) Compare(encoded, raw string) error {
	encd := b.enc.EncodeToString([]byte(raw))
	if encd != encoded {
		return ErrMismatchedEncodedRaw
	}

	return nil
}
