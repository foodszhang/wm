package controllers

type UserSiginForm struct {
	Username string `json:"username" valid:"Required;Maxsize(20)"`
	Password string `json:"password" valid:"Required"`
	Email    string `json:"email" valid:Email"`
}
