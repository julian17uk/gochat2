package main

import (
	"bytes"
	"testing"
	"../internal/ipv6"
)

func TestReadIPv6Address(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("2a00:23c4:8682:5e00:dc3c:35eb:ccef:4d1b"))
	result := ipv6.ReadIPv6Address(&stdin)

	expectedResult := "2a00:23c4:8682:5e00:dc3c:35eb:ccef:4d1b"

	if result != expectedResult {
		t.Errorf("GetIPv6Address test failure")
	}
}

func TestGetIPv6Address(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("2a00:23c4:8682:5e00:dc3c:35eb:ccef:4d1b"))
	result := ipv6.GetIPv6Address(&stdin, "8081")

	expectedResult := "[2a00:23c4:8682:5e00:dc3c:35eb:ccef:4d1b]:8081"

	if result != expectedResult {
		t.Errorf("GetIPv6Address test failure")
	}
}