package Base32Decoder

import (
	"encoding/base32"
	"fmt"
)

func Base32Decode(data string) string {
	decodeBytes, err := base32.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
	}
	decodeString := string(decodeBytes)
	result := decodeString
	return result
}
