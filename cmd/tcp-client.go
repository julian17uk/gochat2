package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"sync"
	"os"
	"../internal/encipher"
	"../internal/keyexchange"
	"../internal/utils"
)

var wg1 = sync.WaitGroup{}

func main() {
	fmt.Println("Please enter IPv6 address to connect:")

	ipv6 := GetIPv6Address()
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
			bytemessage := utils.MessageToByteArray(message)
			plaintext := encipher.AesDecrypt(symmetricKey, bytemessage)
			fmt.Print("\nMessage:", plaintext)
			fmt.Print("Text to send:")
		}
	}()

	wg1.Wait()
}

func GetIPv6Address() string {
	reader0 := bufio.NewReader(os.Stdin)
	ipv6, _ := reader0.ReadString('\n')
	ipv6 = strings.TrimSuffix(ipv6, "\n")
	ipv6 = strings.TrimSpace(ipv6)
	ipv6 = "[" + ipv6 + "]:8081"
	return ipv6
}
