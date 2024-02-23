package Encoder

import (
	"fmt"
	"moonlight/Encode/Base64"
)

func Base64(data interface{}) string {
	result := Base64Encoder.Base64Encode(data)
	fmt.Println(result)
	return result
}
