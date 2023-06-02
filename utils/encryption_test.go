package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"golang.org/x/crypto/pbkdf2"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "Hello, World!"
	password := "password123"

	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		t.Fatal("Failed to generate salt:", err)
	}

	key := pbkdf2.Key([]byte(password), salt, 1000, 32, sha256.New)

	encodedData, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatal("Encrypt failed:", err)
	}

	decryptedText, err := Decrypt(encodedData, key)
	if err != nil {
		t.Fatal("Decrypt failed:", err)
	}

	if decryptedText != plaintext {
		t.Errorf("Decrypted text does not match: expected '%s', got '%s'", plaintext, decryptedText)
	}
}

func TestEncryptDecryptWithEmptyPlaintext(t *testing.T) {
	password := "password123"

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		t.Fatal("Failed to generate salt:", err)
	}

	key := pbkdf2.Key([]byte(password), salt, 1000, 32, sha256.New)

	encodedData, err := Encrypt("", key)
	if err != nil {
		t.Fatal("Encrypt failed:", err)
	}

	decryptedText, err := Decrypt(encodedData, key)
	if err != nil {
		t.Fatal("Decrypt failed:", err)
	}

	if decryptedText != "" {
		t.Error("Decrypted text should be empty")
	}
}

func TestEncryptDecryptWithInvalidEncodedData(t *testing.T) {
	password := "password123"

	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		t.Fatal("Failed to generate salt:", err)
	}

	key := pbkdf2.Key([]byte(password), salt, 1000, 32, sha256.New)

	// Decrypt with invalid encoded data
	_, err := Decrypt("invaliddata", key)
	if err == nil {
		t.Error("Decrypt should have failed with invalid encoded data")
	}
}

func TestEncryptDecryptRandomKey(t *testing.T) {
	plaintext := "Hello, World!"
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal("Failed to generate key:", err)
	}

	encodedData, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatal("Encrypt failed:", err)
	}

	decryptedText, err := Decrypt(encodedData, key)
	if err != nil {
		t.Fatal("Decrypt failed:", err)
	}

	if decryptedText != plaintext {
		t.Errorf("Decrypted text does not match: expected '%s', got '%s'", plaintext, decryptedText)
	}
}

func TestEncryptDecryptEmptyKey(t *testing.T) {
	plaintext := "Hello, World!"
	key := []byte{}

	_, err := Encrypt(plaintext, key)
	if err == nil {
		t.Error("Encrypt should have failed with an empty key")
	}

	_, err = Decrypt("encodeddata", key)
	if err == nil {
		t.Error("Decrypt should have failed with an empty key")
	}
}
