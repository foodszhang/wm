package controllers

type UserSigninForm struct {
	Username string `json:"username" valid:"Required;MaxSize(20)"`
	Password string `json:"password" valid:"Required"`
	Email    string `json:"email" valid:Email"`
}
