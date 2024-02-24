package Encoder

import (
	"fmt"
	"moonlight/Encode/Base32"
	"moonlight/Encode/Base64"
)

func Base64(data interface{}) string {
	result := Base64Encoder.Base64Encode(data)
	fmt.Println(result)
	return result
}

func Base32(data interface{}) string {
	result := Base32Encoder.Base32Encode(data)
	fmt.Println(result)
	return result
}
