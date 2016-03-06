package test

import (
	"testing"
	"wm/controllers"
)

func TestCreate(t *testing.T) {
	token := controllers.Token{"auth", "wm", 60 * 60 * 24 * 30}
	str := token.Create("123", 60)
	data, b := token.Retrieve(str).(string)
	if !b || data != "123" {
		t.Errorf("error! %s", data)
	}

}
