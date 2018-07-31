package common

import (
	"encoding/json"
	"fmt"
)

type RD map[string]interface{}


//@name maptojsstr字典类型转json字符串
//@param mp map
//@return string
func MapToJsStr(mp map[string]interface{})string{
	s := []map[string]interface{}{}
	s = append(s, mp)
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("JSON转str失败!:", err)
		return ""
	}
	return string(b)
}