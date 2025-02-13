package chacha20poly1305

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plainText := []byte("Hello, World!")
	key := generateChaCha20Key("securekey")

	encrypted, err := Encrypt(plainText, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	fmt.Println(string(encrypted))

	if encrypted == "" {
		t.Errorf("Encrypted text is empty")
	}
}

func TestDecrypt(t *testing.T) {
	plainText := []byte("Hello, World!")
	key := generateChaCha20Key("securekey")

	encrypted, _ := Encrypt(plainText, key)
	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}
	fmt.Println(string(decrypted))

	if string(decrypted) != string(plainText) {
		t.Errorf("Decrypted text is not the same as the original, got: %s, want: %s", decrypted, plainText)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	plainText := []byte("Hello, World!")
	key := generateChaCha20Key("securekey")

	encrypted, errEnc := Encrypt(plainText, key)
	if errEnc != nil {
		t.Fatalf("Encrypt failed: %v", errEnc)
	}

	decrypted, errDec := Decrypt(encrypted, key)
	if errDec != nil {
		t.Fatalf("Decrypt failed: %v", errDec)
	}

	if string(decrypted) != string(plainText) {
		t.Errorf("Original text and decrypted text do not match, got: %s, want: %s", decrypted, plainText)
	}
}

func generateChaCha20Key(str string) []byte {
	hash := sha256.New()
	hash.Write([]byte(str))
	key := hash.Sum(nil)
	return key
}
