package main

import (
	"net"
	"fmt"
	"bufio"
//	"strings"
	"sync"
	"os"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

var wg = sync.WaitGroup{}


func main() {
	wg.Add(1)

	findmyipaddress1()
	
	fmt.Println("Launching server")

	// this encryption once working should be done after a connection is established!!!
	//	sendencrypt()
	// newstring := "My Super Secret Code Stuff"
	// bytearray := sendencrypt(newstring)
	// fmt.Println("Byte Array:", bytearray)

	// listen on all interfaces
	// ipv6
	ln, _ := net.Listen("tcp6", "[::]:8081")
	//ln, _ := net.Listen("tcp4", ":8081")
	// accept connection on port
	conn, _ := ln.Accept()
	fmt.Println("Connection established")

	go func() { // wait for text input with \n
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send:")
			text, _ := reader.ReadString('\n')
			//encodedtext := sendencrypt(text)
			//fmt.Println("Encoded text: ", encodedtext)
			fmt.Fprintf(conn, text)
			fmt.Println("Unencoded message sent!")
		}
	}()
	go func() { // wait for message from tcp-client.go
		for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\nMessage received:"+message)
		if message == "" {
			wg.Done()
			break
		}
		fmt.Print("Text to send:")
		}
	}()
	wg.Wait()
}

func sendencrypt(texttosend string) string {
	fmt.Println("Encryption Program v0.01")

	text := []byte(texttosend)
	key :=[]byte("passphrasewhichneedstobe32bytes!")

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
	encryptedtext := gcm.Seal(nonce, nonce, text, nil)

	fmt.Println("Encrypted text as byte array", encryptedtext)

	err = ioutil.WriteFile("myfile.data", encryptedtext, 0777)

	if err != nil {
		fmt.Println(err)
	}

	return string(encryptedtext)
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