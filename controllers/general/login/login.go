package login

import (
	"fmt"
	"runtime/debug"

	"ContactBook/model/db"
	sqlx "database/sql"
	"encoding/base64"
	"errors"

	"github.com/astaxie/beego"
	"tk.com/database/sql"
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

	c.TplName = "contact/home/home.html"
	return
}

func (c *Login) Get() {

	c.TplName = "contact/login/login.html"
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
			c.TplName = "contact/login/login.html"

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

	var row *sqlx.Rows
	row, err = db.Db.Query(`SELECT username,password FROM contact.user WHERE username=$1`, username)
	if err != nil {
		responseMsg = "contact View Fail"
	}

	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		responseMsg = "login Fail"
		return
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
