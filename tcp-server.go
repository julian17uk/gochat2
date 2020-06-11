package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"sync"
	"os"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"crypto/rsa"
	"io"
	"encoding/json"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	findmyipaddress1()
	fmt.Println("Launching server, waiting for incoming connection...")

	ln, _ := net.Listen("tcp6", "[::]:8081")
	conn, _ := ln.Accept()

	symmetricKey := handleAesKeyExchange(conn)
	fmt.Println("Key exchange successful. Connection established")

	go func() { 
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send:")
			text, _ := reader.ReadString('\n')
			encodedtext := aesEncrypt(symmetricKey, text)
			conn.Write(encodedtext)
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

func handleAesKeyExchange(conn net.Conn) []byte {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey
	pubInJson, err := json.Marshal(publicKey)
	conn.Write(pubInJson)
	conn.Write([]byte("\n"))

	symmetricKeyMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	symmetricKeyMessage = strings.TrimSuffix(symmetricKeyMessage, "\n")
	symmetricKeyByte := []byte(symmetricKeyMessage)

	// decrypt symmetricKey with rsa private key
	var v []byte
	err = json.Unmarshal(symmetricKeyByte, &v)
	hash := sha512.New()
	symmetricKey, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, v, nil)

	return symmetricKey
}

func aesEncrypt(symmetrickey []byte, text string) []byte {
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

func findmyipaddress1() {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
    	addrs, _ := i.Addrs()
    	// handle err
    	for _, addr := range addrs {
        	var ip net.IP
        	switch v := addr.(type) {
        	case *net.IPNet:
                ip = v.IP
        	case *net.IPAddr:
                ip = v.IP
        	}
		// process IP address
			fmt.Println("This machines IP address is ", ip)
    	}
	}
}

// func TestRSA(privateKey *PrivateKey) {
// 	publicKey := privateKey.PublicKey
	
// 	encryptedBytes, err := rsa.EncryptOAEP(
// 		sha256.New(),
// 		rand.Reader,
// 		&publicKey,
// 		[]byte("super secret message"),
// 		nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("encrypted bytes: ", encryptedBytes)

// 	fmt.Println("Now decrypting bytes...")

// 	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("decrypted message: ", string(decryptedBytes))

// 	fmt.Println("By the way the public key is: ", publicKey)
// }