package models

import (
	"fmt"
	"github.com/dineshappavoo/basex"
	"strconv"
	"time"
)

func Encode(Input string) string {
	now := time.Now()
	nanos := now.UnixNano()
	s := strconv.FormatInt(nanos, 10)
	encoded := basex.Encode(s)
	fmt.Println("Encoded : ", encoded)
	return encoded
}

func GenerateObjectId() string {
	now := time.Now()
	nanos := now.UnixNano()
	s := strconv.FormatInt(nanos, 10)
	encoded := basex.Encode(s)
	fmt.Println("Encoded : ", encoded)
	return encoded
}

func Decode(Input string) string {
	decoded := basex.Decode(Input)
	fmt.Println("Decoded : ", decoded)
	return decoded
}
