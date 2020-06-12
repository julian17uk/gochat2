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

var wg1 = sync.WaitGroup{}

func main() {
	fmt.Println("Please enter IPv6 address to connect:")
	ipv6 := ipv6.GetIPv6Address(os.Stdin , "8081")

	conn, _ := net.Dial("tcp6", ipv6)
	symmetricKey := keyexchange.HandleClient(conn)
	
	fmt.Println("Key exchange successful. Connection established")
	wg1.Add(1)
	go func() { 
		for {
		 reader := bufio.NewReader(os.Stdin)
		 fmt.Print("Text to send: ")
		 text, _ := reader.ReadString('\n')
		 conn.Write(encipher.AesEncrypt(symmetricKey, text))
		 conn.Write([]byte("\n"))
		}
	}()
	go func() { 
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			if message == "" {
				wg1.Done()
				break
			}
			plaintext := encipher.AesDecrypt(symmetricKey, utils.MessageToByteArray(message))
			fmt.Print("\nMessage:", plaintext)
			fmt.Print("Text to send:")
		}
	}()
	wg1.Wait()
}
