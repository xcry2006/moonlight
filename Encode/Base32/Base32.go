package Base32Encoder

import (
	"encoding/base32"
	"fmt"
)

func Base32Encode(data interface{}) string {
	byteData := []byte(fmt.Sprintf("%v", data) + "\n")
	encode := base32.StdEncoding.EncodeToString(byteData)
	result := encode
	return result
}
