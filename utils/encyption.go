package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var SecretKey = []byte("0123456789abcdef0123456789abcdef")

func Encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Generate a random IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)

	// Create a new AES cipher block mode with the IV
	mode := cipher.NewCBCEncrypter(block, iv)

	// Pad the plaintext to a multiple of the block size
	paddedPlaintext := PKCS7Padding([]byte(plaintext), aes.BlockSize)

	// Encrypt the padded plaintext
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	// Concatenate the IV and ciphertext, and encode as base64
	encryptedData := append(iv, ciphertext...)
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	return encodedData, nil
}

func Decrypt(encodedData string, key []byte) (string, error) {
	// Decode the base64-encoded data
	encryptedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Extract the IV from the encrypted data
	iv := encryptedData[:aes.BlockSize]

	// Create a new AES cipher block mode with the IV
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the ciphertext
	decryptedText := make([]byte, len(encryptedData)-aes.BlockSize)
	mode.CryptBlocks(decryptedText, encryptedData[aes.BlockSize:])

	// Remove padding from the decrypted plaintext
	unpaddedText := PKCS7Unpadding(decryptedText)

	return string(unpaddedText), nil
}

// PKCS7Padding adds PKCS7 padding to the given data
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	paddedData := append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
	return paddedData
}

// PKCS7Unpadding removes PKCS7 padding from the given data
func PKCS7Unpadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}
