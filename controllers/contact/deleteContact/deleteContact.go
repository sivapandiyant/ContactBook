package deleteContact

import (
	"ContactBook/model/db"
	"runtime/debug"

	"fmt"

	"github.com/astaxie/beego"
)

type DeleteContact struct {
	beego.Controller
}

func (c *DeleteContact) Get() {

	responseMsg := ""
	var err error

	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			fmt.Println("Exception", string(stack))
			return
		}

		c.Data["Message"] = responseMsg
		c.TplName = "contact/searchContact/searchContact.html"
		return
	}()

	contactId := c.Ctx.Input.Param(":ID")

	err = db.UpdateStatus(contactId)

	if err != nil {
		responseMsg = "contact Delete Fail"
		return
	}

	responseMsg = "Contact Deleted Success"
	return
}
