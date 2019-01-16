package createContact

import (
	"tk.com/util/log"

	"ContactBook/model/db"
	"runtime/debug"
	"strings"

	"github.com/astaxie/beego"
)

type CreateContact struct {
	beego.Controller
}

func (c *CreateContact) Get() {

	c.TplName = "contact/createContact/createContact.html"
	return
}

func (c *CreateContact) Post() {

	responseMsg := ""
	var err error

	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))

		}

		c.Data["Message"] = responseMsg
		c.TplName = "contact/createContact/createContact.html"
		return
	}()

	inputMobile := c.Input().Get("inputMobile")
	if inputMobile == "" {
		responseMsg = ("inputMobile can't be blank.")
		return
	}

	inputName := c.Input().Get("inputName")
	if inputName == "" {
		responseMsg = ("inputName can't be blank.")
		return
	}

	inputEmail := c.Input().Get("inputEmail")
	if inputEmail == "" {
		responseMsg = ("inputEmail can't be blank.")
		return
	}
	err = db.InsertDB(inputMobile, inputEmail, inputName)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			responseMsg = "Entered email ID is duplicate"
		} else {
			responseMsg = "contact creation Fail"
		}

		return
	}
	responseMsg = "Contact creation Success"

	return
}
