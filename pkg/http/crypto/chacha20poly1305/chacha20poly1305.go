package chacha20poly1305

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/chacha20poly1305"
)

// Encrypt example usage
//
// plaintext := []byte("Hello, World!")
//
// ciphertext, err := Encrypt(plaintext, key)
func Encrypt(plainText []byte, key []byte) (string, error) {
	c, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, c.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	cipherText := c.Seal(nil, nonce, plainText, nil)

	result := make([]byte, 0, len(nonce)+len(cipherText))
	result = append(result, nonce...)
	result = append(result, cipherText...)
	return base64.StdEncoding.EncodeToString(result), nil
}

// Decrypt example usage
//
// decryptedText, err := decrypt(ciphertext, key)
func Decrypt(cipherText string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	c, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	nonce := data[:c.NonceSize()]
	cipherText = string(data[c.NonceSize():])

	plainText, err := c.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
