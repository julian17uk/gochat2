package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"sync"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha512"
	"encoding/json"
	"strings"
	"io"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)

	fmt.Println("Please enter IPv6 address to connect to:")
	reader0 := bufio.NewReader(os.Stdin)
	ipv6, _ := reader0.ReadString('\n')
	ipv6 = "[" + strings.TrimSuffix(ipv6, "\n") + "]:8081"
	conn, _ := net.Dial("tcp6", ipv6)
	
	symmetricKey := handleAesKeyExchangeClient(conn)
	fmt.Println("Key exchange successful. Connection established")

	go func() { 
		for {
		 reader := bufio.NewReader(os.Stdin)
		 fmt.Print("Text to send: ")
		 text, _ := reader.ReadString('\n')
		 conn.Write(aesEncrypt(symmetricKey, text))
		 conn.Write([]byte("\n"))
		}
	}()
	go func() { 
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			if message == "" {
				wg.Done()
				break
			}
			bytemessage := []byte(strings.TrimSuffix(message, "\n"))
			plaintext := aesDecrypt(symmetricKey, bytemessage)
			fmt.Print("\nMessage:", plaintext)
			fmt.Print("Text to send:")
		}
	}()

	wg.Wait()
}

func handleAesKeyExchangeClient(conn net.Conn) []byte {
	publicKeyMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	publicKeyMessage = strings.TrimSuffix(publicKeyMessage, "\n")
	publicKeyByte := []byte(publicKeyMessage)
	publicKey := &rsa.PublicKey{}
	err = json.Unmarshal(publicKeyByte, publicKey)
	if err != nil {
		panic(err)
	}
	
	// generate aes symmetric key, encrypt (with rsa public key) and send to server
	symmetrickey := make([]byte, 32)
	rand.Read(symmetrickey)

	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, symmetrickey, nil)
	if err != nil {
		panic(err)
	}
	pubInJson, err := json.Marshal(ciphertext)
	conn.Write(pubInJson)
	conn.Write([]byte("\n"))

	return symmetrickey
}

func aesEncrypt(symmetrickey []byte, text string) []byte {
	bytetext := []byte(text)
	key := symmetrickey

	c, err := aes.NewCipher(key)

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

func aesDecrypt(symmetrickey []byte, ciphertext []byte) string {
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
