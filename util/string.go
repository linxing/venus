package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
)

func b2s(b []byte) string {

	return (string)(b)
}

func SnakeCasedName(name string) string {

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

	return b2s(newStr)
}

func Struct2Json(a interface{}) string {

	b, err := json.Marshal(a)
	if err != nil {
		log.Fatalf("Marshal failed, err: %v map:%v", err, a)
		return ""
	}
	c := string(b)
	return c
}

// string to int64
func String2Int64(s string) int64 {
	var re int64
	re, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return re
	}
	return re
}

func Base64UrlEncode(b string) string {
	bBytes := []byte(b)
	return base64.URLEncoding.EncodeToString(bBytes)
}

func GetUrlBuild(link string, data map[string]string) string {
	u, _ := url.Parse(link)
	q := u.Query()
	for k, v := range data {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func Struct2Map(obj interface{}) map[string]interface{} {
	objV := reflect.ValueOf(obj)
	v := objV.Elem()
	typeOfType := v.Type()
	var data = make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		data[typeOfType.Field(i).Name] = field.Interface()
	}
	return data
}

func Interface2Map(m map[string]interface{}) map[string]string {
	ret := make(map[string]string, len(m))
	for k, v := range m {
		ret[k] = fmt.Sprint(v)
	}
	return ret
}
