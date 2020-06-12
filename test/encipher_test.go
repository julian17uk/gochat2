package main

import (
	"testing"
	"../internal/encipher"
)

func TestAesEncrypt(t *testing.T) {
	text := "My Super Secret Code Stuff"
	key := []byte("passphrasewhichneedstobe32bytes!")

	encryptedByteArray := encipher.AesEncrypt(key, text)
	if encryptedByteArray == nil {
		t.Errorf("aesEncrypt function returned nil")
	}
	sampleEncryptedByteArray := [...]byte{16, 209, 225, 110, 168, 68, 48, 249, 62, 97, 49, 228, 220, 24, 254, 184, 217, 213, 42, 92, 66, 92, 63, 104, 18, 202, 110, 41, 44, 106, 110, 133, 170, 89, 117, 164, 227, 214, 48, 14, 18, 249, 52, 7, 221, 240, 186, 1, 42, 242, 191, 35, 236, 206}
	if len(encryptedByteArray) != len(sampleEncryptedByteArray) {
		t.Errorf("aesEncrypt result is of differnt lenght to expected result")
	}
	plaintext := encipher.AesDecrypt(key, encryptedByteArray)

	if plaintext != text {
		t.Errorf("AesEncrypt and AesDecrypt test failure")
	}
}

func TestAesDecrypt(t *testing.T) {
	ciphertext := []byte{16, 209, 225, 110, 168, 68, 48, 249, 62, 97, 49, 228, 220, 24, 254, 184, 217, 213, 42, 92, 66, 92, 63, 104, 18, 202, 110, 41, 44, 106, 110, 133, 170, 89, 117, 164, 227, 214, 48, 14, 18, 249, 52, 7, 221, 240, 186, 1, 42, 242, 191, 35, 236, 206}
	key := []byte("passphrasewhichneedstobe32bytes!")

	plaintext := encipher.AesDecrypt(key, ciphertext)
	if plaintext == "" {
		t.Errorf("aesDecrypt function returned nil")
	}

	expectedresult := "My Super Secret Code Stuff"
	if plaintext != expectedresult {
		t.Errorf("AeDecrypt unit test error. Result: %s, Expected: %s", plaintext, expectedresult)
	}
}

