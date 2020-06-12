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
)

var wg = sync.WaitGroup{}

func main() {
	FindIpv6Address()
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
		bytemessage := utils.MessageToByteArray(message)
		plaintext := encipher.AesDecrypt(symmetricKey, bytemessage)
		fmt.Print("\nMessage:", plaintext)
		fmt.Print("Text to send:")
		}
	}()

	wg.Wait()
}



func FindIpv6Address() {
	ifaces, _ := net.Interfaces()
	var ipaddress net.IP

	for _, i := range ifaces {
    	addrs, _ := i.Addrs()
    	for _, addr := range addrs {
        	var ip net.IP
        	switch v := addr.(type) {
	    	case *net.IPNet:
                ip = v.IP
	      	case *net.IPAddr:
              ip = v.IP
        	}
			var bytearray []byte
			bytearray = ip
			if bytearray[0] != 0 && bytearray[0] != 254 {
				ipaddress = ip
			}
		}
	}
	fmt.Println("This machines IPv6 address is ", ipaddress)
}