package login

import (
	"fmt"
	"runtime/debug"

	"ContactBook/model/db"
	"encoding/base64"
	"errors"

	"github.com/astaxie/beego"
)

type Login struct {
	beego.Controller
}

type Home struct {
	beego.Controller
}

type Logout struct {
	beego.Controller
}

func (c *Home) Get() {

	c.TplName = "general/home/home.html"
	return
}

func (c *Login) Get() {

	c.TplName = "general/login/login.html"
	return
}

func (c *Login) Post() {

	responseMsg := ""
	var err error

	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			fmt.Println("Exception", string(stack))
			return
		}

		if err != nil {
			c.Data["Message"] = responseMsg
			c.TplName = "general/login/login.html"

		} else {
			c.Redirect("/Home", 302)

		}

		return
	}()

	username := c.Input().Get("username")
	if username == "" {
		responseMsg = ("username can't be blank.")
		return
	}

	password := c.Input().Get("password")
	if password == "" {
		responseMsg = ("password can't be blank.")
		return
	}

	data, err := db.SelectUser(username)

	if err != nil {
		responseMsg = "login Fail"
	}

	if len(data) <= 0 {
		responseMsg = "login Fail"
		return
	}

	auth := username + ":" + password
	pass := base64.StdEncoding.EncodeToString([]byte(auth))
	if pass != data[0][1] {
		responseMsg = "username/ Password entered is Invalid"
		err = errors.New(responseMsg)

		return
	} else {
		err = nil
	}
	return

}

func (c *Logout) Get() {
	c.Redirect("/Login", 302)
	return
}
