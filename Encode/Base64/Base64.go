package Base64Encoder

import (
	"encoding/base64"
	"fmt"
)

func Base64Encode(data interface{}) string {
	byteData := []byte(fmt.Sprintf("%v", data) + "\n")
	encode := base64.StdEncoding.EncodeToString(byteData)
	result := (encode + "\n")
	return result
}
