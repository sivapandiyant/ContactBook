package main

import (
	"ContactBook/controllers/contact/createContact"
	"ContactBook/controllers/contact/deleteContact"
	"ContactBook/controllers/contact/searchContact"
	"ContactBook/controllers/contact/updateContact"
	"ContactBook/controllers/contact/viewContact"
	"ContactBook/controllers/general/login"

	"github.com/astaxie/beego"
)

func main() {

	beego.Router(beego.AppConfig.String("MAIN_PATH"), &login.Login{})
	beego.Router(beego.AppConfig.String("LOGIN_PATH"), &login.Login{})
	beego.Router(beego.AppConfig.String("LOGOUT_PATH"), &login.Logout{})
	beego.Router(beego.AppConfig.String("HOME_PATH"), &login.Home{})
	beego.Router(beego.AppConfig.String("SEARCH_CONTACT_PATH"), &searchContact.SearchContact{})
	beego.Router(beego.AppConfig.String("CREATE_CONTACT_PATH"), &createContact.CreateContact{})
	beego.Router(beego.AppConfig.String("VIEW_CONTACT_PATH"), &viewContact.ViewContact{})
	beego.Router(beego.AppConfig.String("UPDATE_CONTACT_PATH"), &updateContact.UpdateContact{})
	beego.Router(beego.AppConfig.String("DELETE_CONTACT_PATH"), &deleteContact.DeleteContact{})
	beego.Run()
}
