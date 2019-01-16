package viewContact

import (
	"runtime/debug"

	"tk.com/util/log"

	"ContactBook/model/db"
	sqlx "database/sql"

	"tk.com/database/sql"

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
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			return
		}

		c.Data["Message"] = responseMsg
		c.TplName = "contact/viewContact/viewContact.html"
		return
	}()

	contactId := c.Ctx.Input.Param(":ID")

	var row *sqlx.Rows
	row, err = db.Db.Query(`SELECT id, name,email,mobile,to_char(create_date,'YYYY-MM-DD HH24:MI:SS') FROM contact.details WHERE id=$1`, contactId)
	if err != nil {
		responseMsg = "contact View Fail"
	}

	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		responseMsg = "contact View Fail"
		return
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
