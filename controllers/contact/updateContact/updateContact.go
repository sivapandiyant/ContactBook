package updateContact

import (
	"runtime/debug"

	"tk.com/util/log"

	"ContactBook/model/db"
	sqlx "database/sql"
	"strings"

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

type UpdateContact struct {
	beego.Controller
}

func (c *UpdateContact) Get() {

	contactId := c.Ctx.Input.Param(":ID")

	responseMsg := ""

	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			return
		}

		c.Data["Message"] = responseMsg
		c.TplName = "contact/updateContact/updateContact.html"
		return
	}()

	var row *sqlx.Rows
	row, err := db.Db.Query(`SELECT id, name,email,mobile,to_char(create_date,'YYYY-MM-DD HH24:MI:SS') FROM contact.details WHERE id=$1`, contactId)
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

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Data - ", data)

	c.Data["ID"] = data[0][0]
	c.Data["Name"] = data[0][1]
	c.Data["Email"] = data[0][2]
	c.Data["Mobile"] = data[0][3]
	return
}

func (c *UpdateContact) Post() {

	responseMsg := ""

	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			return
		}

		c.Data["Message"] = responseMsg
		c.TplName = "contact/updateContact/updateContact.html"
		return
	}()

	contactId := c.Ctx.Input.Param(":ID")

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

	err = db.InsertDB(inputMobile, inputEmail, inputName, contactId)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			responseMsg = "Entered email ID is duplicate"
		} else {
			responseMsg = "contact update Fail"
		}

		return
	}
	responseMsg = "Contact Update Success"

	return
}
