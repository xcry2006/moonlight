package HexDecoder

import (
	"encoding/hex"
	"fmt"
)

func HexDecode(data string) string {
	decodeBytes, err := hex.DecodeString(data)
	if err != nil {
		fmt.Println(err)
	}
	decodeString := string(decodeBytes)
	result := decodeString
	return result
}
