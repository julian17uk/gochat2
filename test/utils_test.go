package main

import (
	"testing"
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
