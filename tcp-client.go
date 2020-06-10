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
//	"math/rand"
//	"io/ioutil"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)

	fmt.Println("Please enter ipv6 address to connect to:")
	reader0 := bufio.NewReader(os.Stdin)
	ipv6, _ := reader0.ReadString('\n')
	//ipv6test = "[2a00:23c4:8682:5e00:dc3c:35eb:ccef:4d1b]:8081"
	ipv6 = strings.TrimSuffix(ipv6, "\n")
	ipv6 = "[" + ipv6 + "]:8081"
	fmt.Println("Echo ipv6:", ipv6)
	// connect to this socket
	//conn, _ := net.Dial("tcp6", "[2a00:23c4:8682:5e00:dc3c:35eb:ccef:4d1b]:8081")
	conn, _ := net.Dial("tcp6", ipv6)
	//conn, _ := net.Dial("tcp4", "86.145.80.193:8081")
	
	fmt.Println("Connection established")
	//publicKey := conn.Read
	publicKeyMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	// need to remove '\n' from end of json message
	publicKeyMessage = strings.TrimSuffix(publicKeyMessage, "\n")
	// convert string back into byte array
	publicKeyByte := []byte(publicKeyMessage)
	publicKey := &rsa.PublicKey{}
	err = json.Unmarshal(publicKeyByte, publicKey)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Public key is: ", publicKey.N)

	// generate, encrypt and send the aes symmetric key to server!

	fmt.Println("Generating and sending encypted key")
	symmetrickey := make([]byte, 32)
	rand.Read(symmetrickey)
	fmt.Println("32 byte key is:", symmetrickey)

	// EncryptWithPublicKey
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, symmetrickey, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("cipher text", ciphertext)
	pubInJson, err := json.Marshal(ciphertext)
	conn.Write(pubInJson)
	var carriagereturn string
	carriagereturn = "\n"
	bytereturn := []byte(carriagereturn)
	conn.Write(bytereturn)

	fmt.Println("encrypted key sent to server!")

	go func() { // wait for text input with \n
		for {
		 reader := bufio.NewReader(os.Stdin)
		 fmt.Print("Text to send: ")
		 text, _ := reader.ReadString('\n')
		 fmt.Fprintf(conn, text)
		 fmt.Println("Unencoded message sent!")
		}
	}()
	go func() { // wait for message from tcp-server.go
		for {
			// fmt.Println("Awaiting messages")
			message, _ := bufio.NewReader(conn).ReadString('\n')
			if message == "" {
				wg.Done()
				break
			}
			message = strings.TrimSuffix(message, "\n")
			bytemessage := []byte(message)
			fmt.Println("Byte Array received: ", bytemessage)

			fmt.Print("Decrypting ciphertext...")
			plaintext := receivedecrypt(symmetrickey, bytemessage)
			fmt.Println("Decrypt: ", plaintext)
			fmt.Print("Text to send:")
		}
	}()

	wg.Wait()
}

func receivedecrypt(symmetrickey []byte, ciphertext []byte) string {
	fmt.Println("Decryption Program v0.01")

	key := symmetrickey
//	ciphertext, err := ioutil.ReadFile("myfile.data")
	fmt.Println(ciphertext)

	c, err := aes.NewCipher(key)
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

	fmt.Println(string(plaintext))

	return string(plaintext)
}
