package keyexchange

import (
	"net"
	"bufio"
	"strings"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha512"
	"encoding/json"
)

func HandleClient(conn net.Conn) []byte {
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

func HandleServer(conn net.Conn) []byte {
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

	// decrypt aes symmetric key (with rsa private key)
	var v []byte
	err = json.Unmarshal(symmetricKeyByte, &v)
	hash := sha512.New()
	symmetricKey, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, v, nil)

	return symmetricKey
}