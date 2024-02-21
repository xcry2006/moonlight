// @Title
// @Description
// @Author
// @Update
package Base64

import (
	"encoding/base64"
)

func Base64Encode[T any](data []T) string {
	str := string(data)
	encode := []byte(str)
	result := base64.StdEncoding.EncodeToString(encode)
	return result
}
