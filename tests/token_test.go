package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"wm/controllers"
)

func TestCreate(t *testing.T) {
	token := controllers.Token{"auth", "wm", 60 * 60 * 24 * 30}
	map_data := map[string]string{
		"name": "wm",
	}
	str := token.Create(map_data, 60)
	data_bytes := token.Retrieve(str)
	var data map[string]string
	b := json.Unmarshal(data_bytes, &data)
	if b != nil {
		t.Errorf("转换失败\n")
		fmt.Println(b)
	} else if data["name"] != map_data["name"] {
		fmt.Println(data)
		fmt.Println(map_data)
	}
	str = token.Create(token, 60)
	data_bytes = token.Retrieve(str)
	var ndata controllers.Token
	b = json.Unmarshal(data_bytes, &ndata)
	if b != nil {
		t.Errorf("转换失败\n")
		fmt.Println(b)
	} else if ndata.RedisPrefix != token.RedisPrefix {
		fmt.Println(ndata)
		fmt.Println(token)
	}
}
