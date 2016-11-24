package kit

import (
	"bytes"
	"strings"
)

// FirstLower : HelloWorld -> helloWorld
func FirstLower(data string) string {
	if len(data) == 0 {
		return data
	}
	charArray := []rune(data)
	firstChar := string(charArray[0])
	lowerFirstChar := strings.ToLower(firstChar)
	result := lowerFirstChar + string(charArray[1:])
	return result
}

// WriteKV : 写查询参数的 键值 对.
func WriteKV(query *bytes.Buffer, k string, v string) {
	if query == nil {
		return
	}
	query.WriteString(FirstLower(k))
	query.WriteString("=")
	query.WriteString(v)
	query.WriteString("&")
}
