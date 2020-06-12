# README #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary: gochat is a project writen by Julian Karnik at ECS Digital. The service is written in golang and provides a chat tool via tcp using IPv6 only. The current version requires the tcp-server.go to be run first. The IPv6 address can then be used as an input into the tcp-client.go for connection.

* Golang: The project makes use of the following golang packages: net, fmt, bufio, strings, sync, os, io, crypto/aes, crypto/cipher, crypto/rsa, crypto/rand, crypto/sha512, encoding/json.

* Encryption Handshake: After the tcp connection between the server and client has been made, there is an initial handshake to secretly exchange a symmetric AES key. This shared key is then used for encryption and decryption of communication.
* 1 First the server creates RSA keys and transmits the public key to the client. 
* 2 The client creates an AES symmetric key and uses the RSA public key to encrypt the AES key which it returns to the server.
* 3 The server then uses it's RSA private key to decrypt the AES key.

* Version 0.1

### How do I get set up? ###

* To set up first install golang on your device. See https://golang.org/doc/install
* Configuration: none required
* Dependencies: This project only makes use of standard golang packages. See https://golang.org/pkg/
* How to run tests: 
* IPv6 check: This app uses IPv6 only. To confirm your machine is connected via IPv6 visit https://test-ipv6.com/ in a browser
* Deployment instructions:
* 1 Run the server user$ go run tcp-server.go
* 2 Run the client user$ go run tcp-client.go

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

For comments contact: julian.karnik@ecs-digital.co.uk
