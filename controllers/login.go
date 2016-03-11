package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"wm/models"
)

type LoginController struct {
	beego.Controller
}

// Operations about Users

// @Title Login
// @Description Login
// @Param	body	body 	models.User	true		"body for user content"
// @Success 200 {map[string]string} {"ok":"true"}
// @Failure 403 body is empty
// @router / [post]
func (l *LoginController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	query := o.QueryTable("user")
	var form UserSigninForm
	valid := validation.Validation{}
	json.Unmarshal(l.Ctx.Input.RequestBody, &form)
	b, err := valid.Valid(form)
	if err != nil {
		log.Println("valid error!")
		log.Println(err)
		l.Data["json"] = map[string]interface{}{"ok": "false", "message": err}

	}
	if !b {
		data := map[string]string{}
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			data[err.Key] = err.Message
		}
		l.Data["json"] = &data
	} else {
		user := models.User{}
		is_hav_err := query.Filter(
			"username", form.Username).Filter(
			"password", form.Password).One(&user)
		if is_hav_err == orm.ErrNoRows {
			l.Data["json"] = map[string]string{"ok": "false", "message": "username or password is incorrect."}
		} else {
			data := map[string]int{"id": user.Id}

			token := AuthToken.Create(data, 60*60*24*30)

			l.Data["json"] = map[string]string{"ok": "true", "access_token": token}
		}
	}
	l.ServeJSON()
}
