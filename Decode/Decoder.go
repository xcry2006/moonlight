package Decoder

import (
	"fmt"
	"moonlight/Decode/Base32"
	"moonlight/Decode/Base64"
	"moonlight/Decode/Hex"
)

func Base64(data string) string {
	result := Base64Decoder.Base64Decode(data)
	fmt.Println(result)
	return result
}

func Base32(data string) string {
	result := Base32Decoder.Base32Decode(data)
	fmt.Println(result)
	return result
}

func Hex(data string) string {
	result := HexDecoder.HexDecode(data)
	fmt.Println(result)
	return result
}
