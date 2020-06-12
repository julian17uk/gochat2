# README #

This application provides encrypted chat communication between any two computers connected to the internet via IPv6. 

### What is this repository for? ###

* Quick summary: gochat is a project writen by Julian Karnik at ECS Digital. The service is written in golang and provides a chat tool via tcp using IPv6 only. The current version requires the tcp-server.go to be run first. The IPv6 address of the computer running tcp-server.go can then be used as an input into the tcp-client.go for connection.

* Golang: The project is written in go (golang) and makes use of the following golang packages: net, fmt, bufio, strings, sync, os, io, crypto/aes, crypto/cipher, crypto/rsa, crypto/rand, crypto/sha512, encoding/json.

* Encryption Handshake: After the tcp connection between the server and client has been made, there is an initial handshake to secretly exchange a symmetric AES key. This shared key is then used for encryption and decryption of communication.
* Step 1. First the server creates RSA public and private keys and transmits the public key to the client. 
* Step 2. The client creates an AES symmetric key and uses the RSA public key to encrypt the AES key which it returns to the server.
* Step 3. The server then uses it's RSA private key to decrypt the AES key.

* Version 1.0

### How do I get set up? ###

* To set up first install golang on your device. See https://golang.org/doc/install
* Configuration: none required
* Dependencies: This project only makes use of standard golang packages. See https://golang.org/pkg/
* Cmd folder holds the go code for tcp-server and tcp-client
* Internal folder holds the internal functions
* How to run tests: user$ go test -v (from within the test folder)
* IPv6 check: This app uses IPv6 only. To confirm your machine is connected via IPv6 visit https://test-ipv6.com/ in a browser
* Deployment instructions:
* Step 1. On Machine A run the server user$ go run tcp-server.go
* Step 2. On Machine B run the client user$ go run tcp-client.go

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

For comments contact: julian.karnik@ecs-digital.co.uk
