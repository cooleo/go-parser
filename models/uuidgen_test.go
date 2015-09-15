package models

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	input := "123456789012345678901234567890"
	encoded := Encode(input)
	fmt.Println("Encoded : ", encoded)
	if encoded == "" {
		t.Error("Expected Encode ", encoded)
	}
}

func TestDecode(t *testing.T) {
	input := "1iUcxiBoEac"
	decoded := Decode(input)
	fmt.Println("Decoded : ", decoded)
	if decoded == "" {
		t.Error("Expected decoded ", decoded)
	}	
}
