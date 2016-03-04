package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id          int
	DateCreated time.Time       `orm:"auto_now_add";type(datetime)`
	DateUpdated time.Time       `orm:"auto_now";type(datetime)`
	IsPublished bool            `orm:"default(false)"`
	Author      *User           `orm:"rel(one)"`
	Title       string          `orm:"index"`
	Content     *ArticleContent `orm:"rel(one)"`
	Tags        []*Tag          `orm:"rel(m2m)"`
}

type ArticleContent struct {
	Id      int
	Content string   `orm:"type(text)"`
	meta    *Article `orm:"reverse(one)"`
}

type Tag struct {
	Id       int
	name     string     `orm:"size(30)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Article), new(ArticleContent), new(Tag))
}
