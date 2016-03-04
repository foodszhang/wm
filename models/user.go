package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Username string `orm:"unique;index"`
	Password string
	Profile  *Profile `orm:"rel(one)"`
}

type Profile struct {
	Id      int
	Gender  string
	Age     int
	Email   string
	Summary string
	Avatar  string
	User    *User `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(User), new(Profile))
}
