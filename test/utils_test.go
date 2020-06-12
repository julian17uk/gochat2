package main

import (
	"net"
	"bufio"
	"testing"
	"encoding/json"
	"../internal/utils"
)

func TestMessageToByteArray(t *testing.T) {
	message := "Hello" + "\n"
	expectedResult := [5]byte{72, 101, 108, 108, 111}

	result := utils.MessageToByteArray(message)

	if result == nil {
		t.Errorf("MessageToByteArray function returned nil")
	}

	var arr [5]byte

	copy(arr[:], result[:5])

	if len(result) != len(expectedResult) {
		t.Errorf("MessageToByteArray of incorrect length")
	} else {
		if arr != expectedResult {
			t.Errorf("MessageToByteArray test failure")
		}
	}
}

func TestJsonWrite(t *testing.T) {
	message := []byte{72, 101, 108, 108, 111}
	expectedResult := "Hello" + "\n"

	go func () {
		conn, err := net.Dial("tcp", ":8082")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		utils.JsonWrite(conn, message)
	}()

	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		t.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		jsonMessage, _ := bufio.NewReader(conn).ReadString('\n')
		jsonMessageByteArray := utils.MessageToByteArray(jsonMessage)
		var msg []byte
		err = json.Unmarshal(jsonMessageByteArray, &msg)
		var arr [5]byte
		copy(arr[:], msg[:5])
		var arrResult [5]byte
		copy(arrResult[:], []byte(expectedResult))

		if arr != arrResult {
			t.Fatal("JsonWrite fatal error")
		}
		return
	}
}