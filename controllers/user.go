package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"wm/models"
)

type SigninController struct {
	beego.Controller
}

// Operations about Users

// @Title signin
// @Description signin
// @Param	body	body 	models.User	true		"body for user content"
// @Success 200 {map[string]string} {"ok":"true"}
// @Failure 403 body is empty
// @router / [post]
func (s *SigninController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	query := o.QueryTable("user")
	var form UserSigninForm
	valid := validation.Validation{}
	json.Unmarshal(s.Ctx.Input.RequestBody, &form)
	b, err := valid.Valid(form)
	if err != nil {
		log.Println("valid error!")
		log.Println(err)
		s.Data["json"] = map[string]interface{}{"ok": "false", "message": err}

	}
	if !b {
		data := map[string]string{}
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			data[err.Key] = err.Message
		}
		s.Data["json"] = &data
	} else {
		user := models.User{}
		is_hav_err := query.Filter("username", form.Username).One(&user)
		if is_hav_err != orm.ErrNoRows {
			s.Data["json"] = map[string]string{"ok": "false", "message": "username is existed"}
		} else {
			user.Username = form.Username
			user.Password = form.Password
			profile := models.Profile{Email: form.Email}
			user.Profile = &profile
			id, err := o.Insert(&profile)
			if err != nil {
				log.Println(err)
				s.Data["json"] = map[string]string{"ok": "false", "err": "create profile error"}
				s.ServeJSON()
				return

			}
			id, err = o.Insert(&user)
			if err != nil {
				log.Println(err)
				s.Data["json"] = map[string]string{"ok": "false", "err": "create user error"}
				s.ServeJSON()
				return

			}
			data := map[string]int64{"id": id}

			token := AuthToken.Create(data, 60*60*24*30)

			s.Data["json"] = map[string]string{"ok": "true", "access_token": token}
		}
	}
	s.ServeJSON()
}

//// @Title Get
//// @Description get all Users
//// @Success 200 {object} models.User
//// @router / [get]
//func (u *UserController) GetAll() {
//	users := models.GetAllUsers()
//	u.Data["json"] = users
//	u.ServeJSON()
//}
//
//// @Title Get
//// @Description get user by uid
//// @Param	uid		path 	string	true		"The key for staticblock"
//// @Success 200 {object} models.User
//// @Failure 403 :uid is empty
//// @router /:uid [get]
//func (u *UserController) Get() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		user, err := models.GetUser(uid)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = user
//		}
//	}
//	u.ServeJSON()
//}
//
//// @Title update
//// @Description update the user
//// @Param	uid		path 	string	true		"The uid you want to update"
//// @Param	body		body 	models.User	true		"body for user content"
//// @Success 200 {object} models.User
//// @Failure 403 :uid is not int
//// @router /:uid [put]
//func (u *UserController) Put() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		var user models.User
//		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//		uu, err := models.UpdateUser(uid, &user)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = uu
//		}
//	}
//	u.ServeJSON()
//}
//
//// @Title delete
//// @Description delete the user
//// @Param	uid		path 	string	true		"The uid you want to delete"
//// @Success 200 {string} delete success!
//// @Failure 403 uid is empty
//// @router /:uid [delete]
//func (u *UserController) Delete() {
//	uid := u.GetString(":uid")
//	models.DeleteUser(uid)
//	u.Data["json"] = "delete success!"
//	u.ServeJSON()
//}
//
//// @Title login
//// @Description Logs user into the system
//// @Param	username		query 	string	true		"The username for login"
//// @Param	password		query 	string	true		"The password for login"
//// @Success 200 {string} login success
//// @Failure 403 user not exist
//// @router /login [get]
//func (u *UserController) Login() {
//	username := u.GetString("username")
//	password := u.GetString("password")
//	if models.Login(username, password) {
//		u.Data["json"] = "login success"
//	} else {
//		u.Data["json"] = "user not exist"
//	}
//	u.ServeJSON()
//}
//
//// @Title logout
//// @Description Logs out current logged in user session
//// @Success 200 {string} logout success
//// @router /logout [get]
//func (u *UserController) Logout() {
//	u.Data["json"] = "logout success"
//	u.ServeJSON()
//}
