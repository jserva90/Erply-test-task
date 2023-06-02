package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// var SecretKey []byte

func GetSecretKey() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return []byte(os.Getenv("SECRET_KEY"))
}

func Encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)

	mode := cipher.NewCBCEncrypter(block, iv)

	paddedPlaintext := PKCS7Padding([]byte(plaintext), aes.BlockSize)

	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	encryptedData := append(iv, ciphertext...)
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	return encodedData, nil
}

func Decrypt(encodedData string, key []byte) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := encryptedData[:aes.BlockSize]

	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedText := make([]byte, len(encryptedData)-aes.BlockSize)
	mode.CryptBlocks(decryptedText, encryptedData[aes.BlockSize:])

	unpaddedText := PKCS7Unpadding(decryptedText)

	return string(unpaddedText), nil
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	paddedData := append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
	return paddedData
}

func PKCS7Unpadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}
