package utils

import (
	"net"
	"strings"
	"encoding/json"
)

func MessageToByteArray(message string) []byte {
	message = strings.TrimSuffix(message, "\n")
	return []byte(message)
}

func JsonWrite(conn net.Conn, message []byte) {
	pubInJson, _ := json.Marshal(message)
	conn.Write(pubInJson)
	conn.Write([]byte("\n"))
}

