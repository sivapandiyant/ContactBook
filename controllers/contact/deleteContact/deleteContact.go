package deleteContact

import (
	"ContactBook/model/db"
	"runtime/debug"

	"github.com/astaxie/beego"
	"tk.com/util/log"
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
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			return
		}

		c.Data["Message"] = responseMsg
		c.TplName = "contact/searchContact/searchContact.html"
		return
	}()

	contactId := c.Ctx.Input.Param(":ID")

	_, err = db.Db.Exec("UPDATE contact.details SET status = $1 WHERE id=$2 ", "INACTIVE", contactId)

	if err != nil {
		responseMsg = "contact Delete Fail"
		return
	}

	responseMsg = "Contact Deleted Success"
	return
}
