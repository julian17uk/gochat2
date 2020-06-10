package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"sync"
	"crypto/aes"
	"crypto/cipher"
//	"io/ioutil"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)

	fmt.Println("Please enter ipv6 address to connect to:")
	reader0 := bufio.NewReader(os.Stdin)
	ipv6, _ := reader0.ReadString('\n')
	fmt.Println("Echo ipv6:", ipv6)
	// connect to this socket
	conn, _ := net.Dial("tcp6", "[2a00:23c4:8682:5e00:dc3c:35eb:ccef:4d1b]:8081")
	//conn, _ := net.Dial("tcp4", "86.145.80.193:8081")
	
	fmt.Println("Connection established")
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
			fmt.Print("\nMessage received: "+message)
			if message == "" {
				wg.Done()
				break
			}
			// fmt.Print("Decrypting ciphertext...")
			// bytearray := []byte(message)
			// plaintext := receivedecrypt(bytearray)
			// fmt.Println("Decrypt: ", plaintext)
			fmt.Print("Text to send:")
		}
	}()

	wg.Wait()
}

func receivedecrypt(ciphertext []byte) string {
	fmt.Println("Decryption Program v0.01")

	key := []byte("passphrasewhichneedstobe32bytes!")
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
