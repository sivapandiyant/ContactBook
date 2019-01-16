package searchContact

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

type SearchContact struct {
	beego.Controller
}

func (c *SearchContact) Get() {

	c.TplName = "contact/searchContact/searchContact.html"
	return
}

func (c *SearchContact) Post() {

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

	inputName := c.Input().Get("inputName")

	inputEmail := c.Input().Get("inputEmail")

	data, err := db.SelectDB(inputName, inputEmail)

	if err != nil {
		responseMsg = "contact Search Fail"
	}

	var result []Contact
	for i := range data {
		var c Contact
		c.ID = data[i][0]
		c.Name = data[i][1]
		c.Email = data[i][2]
		c.Mobile = data[i][3]
		c.CreateDate = data[i][4]

		result = append(result, c)
	}

	c.Data["Contact"] = result
	responseMsg = "SUCCESS"

	return
}
