package logout

import (
	"log"
	"model/db"
	"model/session"
	"model/utils"
	"runtime/debug"

	"github.com/astaxie/beego"
)

type Logout struct {
	beego.Controller
}

func (c *Logout) Get() {
	uname := ""
	sess_id := ""
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)
	defer func() {
		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
			c.Data["DisplayMessage"] = "Something went wrong.Please Contact CustomerCare."
			c.TplName = "error/error.html"

			//			_, _ = db.Db.Exec(`INSERT INTO event."user"(url, status,parameter, pip,session_id,username,activity, dump_time) VALUES ('` + c.Ctx.Input.URL() + `','EXCEPTION','` + string(stack) + `','` + c.Ctx.Input.IP() + `','` + sess_id + `','` + uname + `','LOGOUT_PAGE_GET',now())`)
			_, _ = db.Db.Exec("INSERT INTO event.\"user\"(url, status,parameter, pip,session_id,username,activity, dump_time) VALUES ($1,'EXCEPTION',$2,$3,$4,$5,'LOGOUT_PAGE_GET',now())", c.Ctx.Input.URL(), string(stack), c.Ctx.Input.IP(), sess_id, uname)
		}
		return
	}()
	utils.SetHTTPHeader(c.Ctx)

	sess, _ := session.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)

	log.Println(beego.AppConfig.String("loglevel"), "Info", "UserName Nil Found")
	un := sess.Get("uname")
	if un != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Info", "UserName Nil Found")
		session.SeTLogoutSession(un.(string))
	} else {
		uname = ""
	}
	sess.SessionRelease(c.Ctx.ResponseWriter)
	session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
	//	_, _ = db.Db.Exec(`INSERT INTO event."user"(url, status,parameter, pip, session_id,username,activity, dump_time) VALUES ('` + c.Ctx.Input.URL() + `','SUCCESS','` + "Redirected to dashboard" + `','` + c.Ctx.Input.IP() + `','` + sess_id + `','` + uname + `','LOGOUT_PAGE_GET',now())`)
	_, _ = db.Db.Exec("INSERT INTO event.\"user\"(url, status,parameter, pip, session_id,username,activity, dump_time) VALUES ($1,'SUCCESS','Redirected to dashboard',$2,$3,$4,'LOGOUT_PAGE_POST',now())", c.Ctx.Input.URL(), c.Ctx.Input.IP(), sess_id, uname)
	c.Redirect("/SSB/", 302)
	return
}
