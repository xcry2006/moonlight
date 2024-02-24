package HexEncoder

import (
	"encoding/hex"
	"fmt"
)

func HexEncode(data interface{}) string {
	byteData := []byte(fmt.Sprintf("%v", data))
	encode := hex.EncodeToString(byteData)
	result := encode
	return result
}
