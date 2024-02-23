package Decoder

import (
	"fmt"
	"moonlight/Decode/Base64"
)

func Base64(data string) {
	result := Base64Decoder.Base64Decode(data)
	fmt.Println(result)
}
