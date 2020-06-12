package keyexchange

import (
	"net"
	"bufio"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha512"
	"encoding/json"
	"../utils"
)

func HandleClient(conn net.Conn) []byte {
	publicKeyMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	publicKeyByte := utils.MessageToByteArray(publicKeyMessage)

	publicKey := &rsa.PublicKey{}
	err = json.Unmarshal(publicKeyByte, publicKey)
	if err != nil {
		panic(err)
	}
	
	symmetricKey := make([]byte, 32)
	rand.Read(symmetricKey)

	hash := sha512.New()
	cipherText, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, symmetricKey, nil)
	if err != nil {
		panic(err)
	}

	utils.JsonWrite(conn, cipherText)

	return symmetricKey
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

	cipherSymmetricKeyMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	cipherSymmetricKeyByte := utils.MessageToByteArray(cipherSymmetricKeyMessage)

	var cipherSymmetricKey []byte
	err = json.Unmarshal(cipherSymmetricKeyByte, &cipherSymmetricKey)
	hash := sha512.New()
	symmetricKey, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, cipherSymmetricKey, nil)

	return symmetricKey
}

