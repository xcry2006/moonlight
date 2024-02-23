package Base64Decoder

import (
	"encoding/base64"
	"fmt"
)

func Base64Decode(data string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
	}
	decodeString := string(decodeBytes)
	result := decodeString
	return result
}
