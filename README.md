# README #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary: gochat is a project writen by Julian Karnik at ECS Digital. The service is written in golang and provides a chat tool via tcp. The current version requires the tcp-server.go to be run first. The IPv6 address can then be used as an input into the tcp-client.go for connection.

* Encryption: The tcp connection between the server and client involves an initial handshake. First the server creates RSA keys and publishes the public key. The client creates an AES symmetric key and uses the RSA public key to encrypt the AES key which it returns to the server. The server then uses it's RSA private key to decrypt the AES key. This shared symmetric AES key is then used for encryption and decryption of communication

* Version 0.1

### How do I get set up? ###

To set up first install golang on your device. See https://golang.org/doc/install
* Configuration: none required
* Dependencies: This project makes use of the following standard golang packages
	"net"
	"fmt"
	"bufio"
	"strings"
	"sync"
	"os"
	"io"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha512"
	"encoding/json"
* Database configuration: no db used
* How to run tests
* Deployment instructions:
	To run the server user$ go run tcp-server.go
	To run the client user$ go run tcp-client.go

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

For comments contact: julian.karnik@ecs-digital.co.uk
