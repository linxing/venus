package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SnakeCasedName(t *testing.T) {

	name := "OK"
	newStr := make([]byte, 0, len(name)+1)
	for i := 0; i < len(name); i++ {
		c := name[i]
		if isUpper := 'A' <= c && c <= 'Z'; isUpper {
			if i > 0 {
				newStr = append(newStr, '_')
			}
			c += 'a' - 'A'
		}
		newStr = append(newStr, c)
	}

	assert.Equal(t, b2s(newStr), "o_k")
}

func Test_Struct2Json(t *testing.T) {

	data := make(map[string]interface{})
	data["a"] = "b"

	rs := "{\"a\":\"b\"}"

	assert.Equal(t, rs, Struct2Json(data))
}

func Test_Base64UrlEncode(t *testing.T) {
	url := "http://127.0.0.1"
	fmt.Println(Base64UrlEncode(url))
	assert.Equal(t, "aHR0cDovLzEyNy4wLjAuMQ==", Base64UrlEncode(url))
}

func Test_GetUrlBuild(t *testing.T) {
	url := "http://127.0.0.1"
	data := map[string]string{
		"a": "1",
		"b": "2",
	}
	assert.Equal(t, "http://127.0.0.1?a=1&b=2", GetUrlBuild(url, data))
}
