package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"sync"
	"os"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/rsa"
	"io"
	"io/ioutil"
	"encoding/json"
)

var wg = sync.WaitGroup{}


func main() {
	wg.Add(1)

	findmyipaddress1()
	
	fmt.Println("Launching server")

	fmt.Println("Welcome to rsa, generating keys...")

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey
	
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte("super secret message"),
		nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypted bytes: ", encryptedBytes)

	fmt.Println("Now decrypting bytes...")

	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}

	fmt.Println("decrypted message: ", string(decryptedBytes))

	fmt.Println("By the way the public key is: ", publicKey)

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

	//publicKeyString := string([]byte publicKey)
	//fmt.Println("Press enter to send the public key")
	//fmt.Fprintf(conn, "%s", publicKey)
	// use JSON to serialised rsa.PublicKey to a []byte
	pubInJson, err := json.Marshal(publicKey)
	fmt.Println("public key in json: ", string(pubInJson))

	conn.Write(pubInJson)

	var carriagereturn string
	carriagereturn = "\n"
	bytereturn := []byte(carriagereturn)
	conn.Write(bytereturn)

	fmt.Println("Waiting for encrypted symmetric key")
	symmetricKeyMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	symmetricKeyMessage = strings.TrimSuffix(symmetricKeyMessage, "\n")
	symmetricKeyByte := []byte(symmetricKeyMessage)
	// Decrypt symmetricKey with RSA
	var v []byte
	err = json.Unmarshal(symmetricKeyByte, &v)
	hash := sha512.New()
	symmetricKey, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, v, nil)

	fmt.Println("This is the decrypted symmetric key!", symmetricKey)

	go func() { // wait for text input with \n
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send:")
			text, _ := reader.ReadString('\n')
			encodedtext := sendencrypt(symmetricKey, text)
			fmt.Println("Encoded text: ", encodedtext)
			conn.Write(encodedtext)
			conn.Write(bytereturn)
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

func sendencrypt(symmetrickey []byte,texttosend string) []byte {
	fmt.Println("Encryption Program v0.01")

	text := []byte(texttosend)
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
	encryptedtext := gcm.Seal(nonce, nonce, text, nil)

	fmt.Println("Encrypted text as byte array", encryptedtext)

	err = ioutil.WriteFile("myfile.data", encryptedtext, 0777)

	if err != nil {
		fmt.Println(err)
	}

	return encryptedtext
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