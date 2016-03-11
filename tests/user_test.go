package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"wm/controllers"
	_ "wm/routers"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "tmp.db")
	orm.RunCommand()
	err := orm.RunSyncdb("default", true, true)

	if err != nil {
		fmt.Println(err)
	}
}

func TestSignin(t *testing.T) {
	user := controllers.UserSigninForm{Username: "admin", Password: "admin", Email: "admin@gmail.com"}
	data, err := json.Marshal(user)
	if err != nil {
		t.Error("转化失败!")
	} else {
		r, _ := http.NewRequest("POST", "/v1/signin", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		beego.Trace("testing", "TestSignin", "Code[%d]\n%s", w.Code, w.Body.String())
		fmt.Println(w.Body.String())
		Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	}
}

func TestLogin(t *testing.T) {
	user := controllers.UserSigninForm{Username: "admin", Password: "admin", Email: "admin@gmail.com"}
	data, err := json.Marshal(user)
	if err != nil {
		t.Error("转化失败!")
	} else {
		http.NewRequest("POST", "/v1/signin", bytes.NewBuffer(data))
		r, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		beego.Trace("testing", "TestLogin", "Code[%d]\n%s", w.Code, w.Body.String())
		fmt.Println(w.Body.String())
		Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	}
}
