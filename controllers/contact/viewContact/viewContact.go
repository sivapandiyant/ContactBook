package viewContact

import (
	"runtime/debug"

	"ContactBook/model/db"
	"fmt"

	"github.com/astaxie/beego"
)

type Contact struct {
	ID         string
	Name       string
	Email      string
	Mobile     string
	CreateDate string
}

type ViewContact struct {
	beego.Controller
}

func (c *ViewContact) Get() {

	responseMsg := ""
	var err error

	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			fmt.Println("Exception", string(stack))
			return
		}

		c.Data["Message"] = responseMsg
		c.TplName = "contact/viewContact/viewContact.html"
		return
	}()

	contactId := c.Ctx.Input.Param(":ID")
	data, err := db.SelectContact(contactId)
	if err != nil {
		responseMsg = "contact View Fail"
	}

	if len(data) <= 0 {
		responseMsg = "contact View Fail"
		return
	}

	c.Data["ID"] = data[0][0]
	c.Data["Name"] = data[0][1]
	c.Data["Email"] = data[0][2]
	c.Data["Mobile"] = data[0][3]

	return
}
