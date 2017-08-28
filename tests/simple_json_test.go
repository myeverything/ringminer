package tests

import (
	"encoding/json"
	"testing"
)

// 构造的时候不能缺少字段，解析的时候可以
func TestJson(t *testing.T) {
	type JT struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Ted  string `json:ted`
	}

	//
	//t := &JT{"dylenfu", 30}
	//m,_ := json.Marshal(t)
	//log.Println("test\t-", "json marshal", string(m))

	m := `{"name":"dylenfu","age":30}`
	u := &JT{}
	json.Unmarshal([]byte(m), u)
	t.Log("test\t-", "json unmarshal", u.Name, u.Age)
}