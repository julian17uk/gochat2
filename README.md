# README #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary: gochat is a project writen by Julian Karnik at ECS Digital. The service is written in golang and provides a chat tool via tcp. The current version requires the tcp-server.go to be run first. The IPv6 address can then be used as an input into the tcp-client.go for connection.

* The tcp connection between the server and client involves a handshake. First the server publishes the RSA public key to the client. The client then uses this public key to encrypt an AES symmetric key and returns to the server. The server then uses it's RSA private key to decrypt the symmetric key. This shared symmetric AES key is then used for all communication encryption and decryption

* Version 0.1
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

* Summary of set up
To set up first install go on your device. See https://golang.org/doc/install
* Configuration
* Dependencies
* Database configuration
* How to run tests
* Deployment instructions

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact