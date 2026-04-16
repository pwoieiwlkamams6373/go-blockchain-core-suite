package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func AESEncrypt(plainText, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func AESDecrypt(cipherText string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decoded, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decoded, decoded)
	return decoded, nil
}

func main() {
	key := []byte("12345678901234567890123456789012")
	plain := []byte("blockchain secret data")
	enc, _ := AESEncrypt(plain, key)
	dec, _ := AESDecrypt(enc, key)
	fmt.Printf("解密结果: %s\n", string(dec))
}
