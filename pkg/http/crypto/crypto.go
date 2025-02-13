package crypto

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/andys920605/meme-coin/pkg/http/crypto/chacha20poly1305"
)

type Crypto struct{}

func NewCrypto() *Crypto {
	return &Crypto{}
}

func (c *Crypto) Encrypt(data any, key []byte) (string, error) {
	plainText, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return chacha20poly1305.Encrypt(plainText, key)
}

func (c *Crypto) Decrypt(cipherText string, key []byte) ([]byte, error) {
	return chacha20poly1305.Decrypt(cipherText, key)
}

func (c *Crypto) GenerateKey(str string) ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(str))
	key := hash.Sum(nil)
	return key, nil
}
