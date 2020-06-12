package main

import (
	"net"
	"fmt"
	"bufio"
	"sync"
	"os"
	"../internal/encipher"
	"../internal/keyexchange"
	"../internal/utils"
	"../internal/ipv6"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("This machines IPv6 address is ", ipv6.FindIpv6Address())
	fmt.Println("Launching server, waiting for incoming connection...")

	ln, _ := net.Listen("tcp6", "[::]:8081")
	conn, _ := ln.Accept()
	symmetricKey := keyexchange.HandleServer(conn)
	
	fmt.Println("Key exchange successful. Connection established")
	wg.Add(1)
	go func() { 
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send:")
			text, _ := reader.ReadString('\n')
			encodedtext := encipher.AesEncrypt(symmetricKey, text)
			conn.Write(encodedtext)
			conn.Write([]byte("\n"))
		}
	}()
	go func() { 
		for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "" {
			wg.Done()
			break
		}
		plaintext := encipher.AesDecrypt(symmetricKey, utils.MessageToByteArray(message))
		fmt.Print("\nMessage:", plaintext)
		fmt.Print("Text to send:")
		}
	}()
	wg.Wait()
}