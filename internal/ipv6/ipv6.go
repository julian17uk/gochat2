package ipv6

import (
	"net"
	"strings"
	"bufio"
	"os"
)

func GetIPv6Address() string {
	reader0 := bufio.NewReader(os.Stdin)
	ipv6, _ := reader0.ReadString('\n')
	ipv6 = strings.TrimSuffix(ipv6, "\n")
	ipv6 = strings.TrimSpace(ipv6)
	ipv6 = "[" + ipv6 + "]:8081"
	return ipv6
}

func FindIpv6Address() net.IP {
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
	return ipaddress
}
