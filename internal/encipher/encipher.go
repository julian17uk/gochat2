package encipher

import (
	"io"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func AesEncrypt(symmetrickey []byte, text string) []byte {
	bytetext := []byte(text)

	c, err := aes.NewCipher(symmetrickey)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	encryptedtext := gcm.Seal(nonce, nonce, bytetext, nil)

	if err != nil {
		fmt.Println(err)
	}

	return encryptedtext
}

func AesDecrypt(symmetrickey []byte, ciphertext []byte) string {
	c, err := aes.NewCipher(symmetrickey)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plaintext)
}